package bssmap

import (
	"encoding/binary"
	"fmt"
)

const SEP = "/"

type IE []byte

func (i IE) Tag() BssmapIE {
	return BssmapIE(i[0])
}
func (i IE) String() string {
	switch i.Tag() {
	case AOIP_TRASP_ADDR:
		s := i.Tag().String() + SEP
		for j := 2; j < 5; j++ {
			s += fmt.Sprintf("%d.", i[j])
		}
		x := uint16(i[6])<<8 + uint16(i[7])
		s += fmt.Sprintf("%d:%d", i[5], x)
		return s
	case CALL_ID:
		return fmt.Sprintf("%s%s%d", i.Tag().String(), SEP, binary.LittleEndian.Uint32(i[1:5]))
	case IMSI:
		s := i.Tag().String() + SEP
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
		return s
	case TMSI:
		return fmt.Sprintf("%s%sTODO", i.Tag().String(), SEP)
	case CELL_ID:
		return fmt.Sprintf("%s%sTODO", i.Tag().String(), SEP)
	default:
		return i.Tag().String()
	}
}

func (i IE) Int() int {
	switch i.Tag() {
	case CALL_ID:
		// todo
		return -1
	case CELL_ID:
		// todo
		return -1
	}
	return 0
}

func (i IE) List() []int {
	switch i.Tag() {
	case CALL_ID_LIST:
		// todo
		a := make([]int, 0)
		return a
	case CELL_ID_LIST:
		// todo
		a := make([]int, 0)
		return a
	}
	return nil
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
		if i == 0 {
			s += m.IEs[i].String()
		} else {
			s += fmt.Sprintf(",%s", m.IEs[i].String())
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

func BssmapDecode(b []byte) (*Bssmap, error) {
	if len(b) < 1 {
		return nil, fmt.Errorf("invalid bssmap message")
	}
	m := &Bssmap{
		Msg: Msg_Type(b[0]),
		IEs: make([]IE, 0),
	}
	offset := 1
	for offset < len(b) {
		ie := BssmapIE(b[offset])
		ieLen := ie.Format()
		// Variable part
		if ieLen == 0 { // (0 means a length is unknown and should be calculated)
			// Next byte is length
			if (offset + 1) >= len(b) {
				return nil, fmt.Errorf("variable IE %q is invalid: [len]", ie)
			}
			if (offset + 1 + int(b[offset+1])) >= len(b) {
				return nil, fmt.Errorf("variable IE %q is invalid: [val]", ie)
			}
			l := int(b[offset+1])
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
			return nil, fmt.Errorf("found unsupported IE %02x", b[offset])
		}
	}
	return m, nil
}
