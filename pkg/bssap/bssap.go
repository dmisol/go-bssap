package bssap

import "go-bssap/bssmap"

type BssapType = uint8

const (
	BssapTypeBssMap 	BssapType = 0x0 // the BSS Management Application Part
	BssapTypeDtap 		BssapType = 0x1 // the Direct Transfer Application Part
)

// NewBssMapMsg builds a new BSS_MAP_ message
func NewBssMapMsg(msg *bssmap.BssmapMsg) []byte {
	b := make([]byte, 1)
	b[0] = BssapTypeBssMap
	b[1] = 0
	var iesTtlLen uint8
	for i := range msg.IEs {
		iesTtlLen += msg.IEs[i].Len()
		b = append(b, msg.IEs[i].Encode()...)
	}
	b[1] = iesTtlLen
	return nil
}

// NewBssDtapMsg builds a new BSS_DTAP_ message
func NewBssDtapMsg() []byte {
	b := make([]byte, 1)
	b[0] = BssapTypeDtap
	return nil
}