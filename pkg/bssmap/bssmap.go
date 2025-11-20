package bssmap

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type IE []byte

func (i IE) Tag() BssmapIE {
	return BssmapIE(i[0])
}

// todo: refactor
func (i IE) String() string {
	switch i.Tag() {
	case AOIP_TRASP_ADDR:
		s := ""
		for j := 2; j < 5; j++ {
			s += fmt.Sprintf("%d.", i[j])
		}
		x := uint16(i[6])<<8 + uint16(i[7])
		s += fmt.Sprintf("%d:%d", i[5], x)
		return s
	case CALL_ID:
		return fmt.Sprintf("%d", binary.LittleEndian.Uint32(i[1:5]))
	case IMSI:
		return i.xMSI()
	case TMSI:
		return i.xMSI()

	case CELL_ID:
		mcc, mnc, ci, lac, err := i.ParseCellId()
		if err != nil {
			return err.Error()
		}
		if mcc == 0 || mnc == 0 {
			return fmt.Sprintf("(lac=%d,ci=%d)", lac, ci)
		}
		return fmt.Sprintf("(mcc=%d,mnc=%d,lac=%d,ci=%d)", mcc, mnc, lac, ci)

	default:
		return ""
	}
}

func (i IE) xMSI() (s string) {
	for j := 2; j < len(i); j++ {
		if j == 2 {
			s += fmt.Sprintf("%d", (i[j]>>4)&0x0F)
		} else {
			s += fmt.Sprintf("%d%d", i[j]&0x0F, (i[j]>>4)&0x0F)
		}
	}
	if i[1]&0x08 == 0 {
		s = s[:len(s)-1]
	}
	return
}

func (i IE) Int() int {
	switch i.Tag() {
	case CALL_ID:
		// todo
		return 0
	case CELL_ID:
		_, _, ci, _, err := i.ParseCellId()
		if err != nil {
			return 0
		}
		return int(ci)
	}
	return 0
}

func (i IE) ListLACs() []int {
	switch i.Tag() {
	case CELL_ID_LIST:
		// ATTN: looking for CellIDs, not LACs
		res := make([]int, 0)
		if len(i) < 3 {
			// error here!
			return res
		}

		pd := i[2]
		var lac_pos, step int

		switch pd {
		case 0: // CellGloalID
			lac_pos = 3
			step = 7
		case 1: // LAC+CI
			step = 4
		case 4:
			step = 5
			lac_pos = 3
		case 5:
			step = 2
		default:
			return res
		}

		offs := 3
		for (offs + step) <= len(i) {
			v := binary.BigEndian.Uint16(i[offs+lac_pos:])
			res = append(res, int(v))
			offs += step
		}
		return res
	}
	return nil
}

func (i IE) ListCIs() []int {
	switch i.Tag() {
	case CELL_ID_LIST:
		// ATTN: looking for CellIDs, not LACs
		res := make([]int, 0)
		if len(i) < 3 {
			// error here!
			return res
		}

		pd := i[2]
		var ci_pos, step int

		switch pd {
		case 0: // CellGloalID
			ci_pos = 5
			step = 7
		case 1: // LAC+CI
			ci_pos = 2
			step = 4
		case 2: // Cell Identity, CI, is used to identify the cells
			ci_pos = 2
			step = 2

		default:
			return res
		}

		offs := 3
		for (offs + step) <= len(i) {
			v := binary.BigEndian.Uint16(i[offs+ci_pos:])
			res = append(res, int(v))
			offs += step
		}
		return res
	}
	return nil
}

func (i IE) Cause() (BSSMAP_Cause, error) {
	if i.Tag() != CAUSE {
		return BSSMAP_Cause(0xFF), fmt.Errorf("wrong TAG(%s)", i.Tag())
	}
	if len(i) != 3 {
		return BSSMAP_Cause(0xFF), fmt.Errorf("invalind len(%s)", hex.EncodeToString(i))
	}
	return BSSMAP_Cause(i[2]), nil
}

type Bssmap struct {
	Msg Msg_Type
	IEs []IE
}

func NewBssmap(mt Msg_Type, ies ...IE) *Bssmap {
	m := &Bssmap{
		Msg: mt,
		IEs: ies,
	}
	return m
}

// Encode() issues BSSAP (not BSSMAP) bytes!
func (m *Bssmap) Encode() []byte {
	r := make([]byte, 3)
	r[0] = 0
	r[2] = byte(m.Msg)

	for i := range len(m.IEs) {
		r = append(r, m.IEs[i]...)
	}

	r[1] = byte(len(r)) - 2
	return r
}

func (m *Bssmap) String() string {
	s := m.Msg.String() + "{"
	for i := range len(m.IEs) {
		if i != 0 {
			s += ","
		}
		s += m.IEs[i].Tag().String()

		v := m.IEs[i].String()
		if len(v) > 0 {
			s += "=" + v
		}
	}
	return s + "}"
}

func (m *Bssmap) GetIE(tag BssmapIE) (IE, bool) {
	for _, v := range m.IEs {
		if v.Tag() == tag {
			return v, true
		}
	}
	return nil, false
}

func (m *Bssmap) GetIEs(tag BssmapIE) ([]IE, bool) {
	var x []IE
	for _, v := range m.IEs {
		if v.Tag() == tag {
			x = append(x, v)
		}
	}
	if len(x) > 0 {
		return x, true
	}
	return nil, false
}

func (i IE) ParseCellId() (mcc, mnc, ci, lac uint16, err error) {
	if i.Tag() != CELL_ID {
		err = fmt.Errorf("error: wrong IE %s", i.Tag().String())
		return
	}
	if len(i) < 3 {
		err = fmt.Errorf("error: CELL_ID is too short %s", hex.EncodeToString(i))
		return
	}
	if int(i[1]) != len(i)-2 {
		err = fmt.Errorf("error: CELL_ID has invalid len %s", hex.EncodeToString(i))
		return
	}

	switch i[2] {
	case 0:
		if len(i) != 10 {
			err = fmt.Errorf("error: CELL_ID (whole CGI) unexpected len %d %s", len(i), hex.EncodeToString(i))
			return
		}
		mcc = uint16(i[3]&0x0F)*100 + uint16(i[3]>>4)*10 + uint16(i[4]&0x0F)

		mnc = uint16((i[5]&0x0F)*10 + i[5]>>4)
		f := i[4] >> 4
		if f != 0x0F {
			mnc = mnc*10 + uint16(f)
		}

		lac = binary.BigEndian.Uint16(i[6:])
		ci = binary.BigEndian.Uint16(i[8:])
	case 1:
		if len(i) != 7 {
			err = fmt.Errorf("error: CELL_ID (whole CGI) unexpected len %d %s", len(i), hex.EncodeToString(i))
			return
		}
		lac = binary.BigEndian.Uint16(i[3:])
		ci = binary.BigEndian.Uint16(i[5:])
	default:
		err = fmt.Errorf("error: CELL_ID option %d not supported %s", int(i[2]), hex.EncodeToString(i))
	}
	return
}

func (m *Bssmap) Remove(tag BssmapIE) {
	for i := 0; i < len(m.IEs); i++ {
		if len(m.IEs[i]) == 0 {
			continue
		}
		if m.IEs[i].Tag() == tag {
			// fmt.Println(tag, "found at", i)
			b := m.IEs[:i]
			b = append(b, m.IEs[i+1:]...)
			m.IEs = b
			return
		}
	}
}

func (m *Bssmap) Replace(upd IE) {
	for i := 0; i < len(m.IEs); i++ {
		if len(m.IEs[i]) == 0 {
			continue
		}
		if m.IEs[i].Tag() == upd.Tag() {
			m.IEs[i] = upd
			return
		}
	}
}

// BssmapDecode() decodes byte stream from BSSAP(!), aauming leading [0,{len}, {Msg} etc..]
func BssmapDecode(b []byte) (*Bssmap, error) {
	if len(b) < 3 {
		return nil, fmt.Errorf("invalid bssmap message")
	}
	m := &Bssmap{
		Msg: Msg_Type(b[2]),
		IEs: make([]IE, 0),
	}
	//fmt.Printf(hex.Dump(b))

	offset := 3
	for offset < len(b) {
		ie := BssmapIE(b[offset])
		ieLen := ie.format()
		// Variable part
		if ieLen == 0 { // (0 means a length is unknown and should be calculated)
			// Next byte is length
			if (offset + 1) >= len(b) {
				return nil, fmt.Errorf("variable IE %q is invalid: [len]", ie)
			}
			l := int(b[offset+1])

			if (offset + 1 + l) >= len(b) {

				//fmt.Println("given IE:", offset, hex.EncodeToString(b[offset:]))
				return nil, fmt.Errorf("variable IE %q (%d bytes) is invalid: [val]", ie, l)
			}
			d := b[offset : offset+l+2]

			m.IEs = append(m.IEs, d)
			offset += l + 2
			// Fixed part
		} else if ieLen > 0 {
			if (offset + ieLen) > len(b) {
				return nil, fmt.Errorf("fixed IE %q can't fit", ie)
			}
			d := b[offset : offset+ieLen]
			m.IEs = append(m.IEs, d)
			offset += ieLen
			// Unsupported IE
		} else {
			for i := 0; i < len(m.IEs); i++ {
				fmt.Printf("\t%02X %s\n", m.IEs[i].Tag(), m.IEs[i].Tag())
			}
			return nil, fmt.Errorf("found unsupported IE %02x at %d", b[offset], offset)
		}
	}
	return m, nil
}
