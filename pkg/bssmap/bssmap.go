package bssmap

import "fmt"

type IE []byte

func (i IE) Tag() BssmapIE {
	return BssmapIE(i[0])
}

type Bssmap struct {
	Msg BSSMAP_MsgType
	IEs []IE
}

func NewBssmap(mt BSSMAP_MsgType, ies ...IE) *Bssmap {
	m := &Bssmap{
		Msg: mt,
		IEs: ies,
	}
	return m
}

func (m *Bssmap) Encode() []byte {
	r := make([]byte, 1)
	r[0] = byte(m.Msg)

	for i := range len(m.IEs) {
		r = append(r, m.IEs[i]...)
	}

	return r
}

func (m *Bssmap) String() string {
	s := m.Msg.String() + "{"
	for i := range len(m.IEs) {
		if i == 0 {
			s += m.IEs[i].Tag().String()
		} else {
			s += fmt.Sprintf(",%s", m.IEs[i])
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

func BssmapDecode(b []byte) (*Bssmap, error) {
	if len(b) < 1 {
		return nil, fmt.Errorf("invalid bssmap message")
	}
	m := &Bssmap{
		Msg: BSSMAP_MsgType(b[0]),
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
