package bssmap

import "fmt"

type DummyIE struct {
	Data []byte
}

func (ie *DummyIE) Tag() BSSMAP_IE {
	return BSSMAP_IE(ie.Data[0])
}

func (ie *DummyIE) Encode() []byte {
	return ie.Data
}
func (ie *DummyIE) String() string {
	return fmt.Sprintf("IE 0x%2X", ie.Data[0])
}
