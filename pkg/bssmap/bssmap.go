package bssmap

import "fmt"

type IE interface {
	Tag() BSSMAP_IE
	Encode() []byte
	String() string
}

type BssmapMsg struct {
	Msg BSSMAP_MsgType
	IEs []IE
}

func NewBssmapMsg(mt BSSMAP_MsgType, ie ...IE) *BssmapMsg {
	m := &BssmapMsg{
		Msg: mt,
		IEs: ie,
	}
	return m
}

func (m *BssmapMsg) Encode() []byte {
	r := make([]byte, 1)
	r[0] = byte(m.Msg)

	for i := range len(m.IEs) {
		r = append(r, m.IEs[i].Encode()...)
	}

	return r
}

func (m *BssmapMsg) String() string {
	s := m.Msg.String() + "{"
	for i := range len(m.IEs) {
		if i == 0 {
			s += m.IEs[i].String()
		} else {
			s += fmt.Sprintf(",%s", m.IEs[i])
		}
	}
	return s + "}"
}

func BssmapDecode(b []byte) (*BssmapMsg, error) {
	if len(b) < 1 {
		return nil, fmt.Errorf("invalid bssmap message")
	}
	m := &BssmapMsg{
		Msg: BSSMAP_MsgType(b[0]),
		IEs: make([]IE, 0),
	}
	offset := 1
	for offset < len(b) {
		ie := BSSMAP_IE(b[offset])
		isVar, fixed := ie.Format()
		if isVar {
			// next byte is length
			if (offset + 1) >= len(b) {
				return nil, fmt.Errorf("variable IE %s is invalid: [len]", ie)
			}
			if (offset + 1 + int(b[offset+1])) >= len(b) {
				return nil, fmt.Errorf("variable IE %s is invalid: [val]", ie)
			}
			l := int(b[offset+1])
			d := &DummyIE{
				Data: b[offset : offset+l+2],
			}
			m.IEs = append(m.IEs, d)
		} else {
			// fixed
			if (offset + fixed) >= len(b) {
				return nil, fmt.Errorf("fixed IE %s can't fit", ie)
			}
			d := &DummyIE{
				Data: b[offset : offset+fixed],
			}
			m.IEs = append(m.IEs, d)
		}
	}
	return m, nil
}
