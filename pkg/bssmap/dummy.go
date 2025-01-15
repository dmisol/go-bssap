package bssmap

import "fmt"

type DummyIE struct {
	Data []byte // Tag + Length + Value
}

func (ie *DummyIE) Tag() BssmapIE {
	return BssmapIE(ie.Data[0])
}

func (ie *DummyIE) Encode() []byte {
	return ie.Data
}

func (ie *DummyIE) Len() uint8 {
	return uint8(len(ie.Data))
}

func (ie *DummyIE) String() string {
	return fmt.Sprintf("IE 0x%2X", ie.Data[0])
}