package bssap

import (
	"fmt"

	"github.com/dmisol/go-bssap/pkg/bssmap"
	"github.com/dmisol/go-bssap/pkg/dtap"
)

type BssapType = uint8

const (
	BssapTypeBssMap BssapType = 0x0 // the BSS Management Application Part
	BssapTypeDtap   BssapType = 0x1 // the Direct Transfer Application Part
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

func ParseBssap(b []byte) (*bssmap.BssmapMsg, *dtap.DtapMsg, error) {
	x := BssapType(b[0])

	if x == BssapTypeBssMap {
		pl := b[2:]
		if int(b[1]) != len(pl) {
			return nil, nil, fmt.Errorf("error wrong bssmap data len(%d), %v", int(b[1]), pl)
		}

		m, err := bssmap.BssmapDecode(pl)
		return m, nil, err
	}
	if x == BssapTypeDtap {
		// dlci := b[1]
		l := int(b[2])
		pl := b[3:]
		if l != len(pl) {
			return nil, nil, fmt.Errorf("error wrong dtap data len(%d), %v", int(b[1]), pl)
		}

		m, err := dtap.DtapDecode(pl)
		return nil, m, err
	}
	return nil, nil, fmt.Errorf("error unknown bssap payload type")
}
