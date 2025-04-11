package bssap

import (
	"encoding/hex"
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
func NewBssap(bssap *bssmap.Bssmap, dtap []byte) []byte {
	if bssap != nil {
		b := make([]byte, 2)
		b[0] = BssapTypeBssMap
		data := bssap.Encode()
		b = append(b, data...)
		b[1] = 2 + byte(len(data))
		return b
	}
	b := []byte{BssapTypeDtap}
	b = append(b, dtap...)
	return b

}

// NewBssDtapMsg builds a new BSS_DTAP_ message
func NewBssDtapMsg() []byte {
	b := make([]byte, 1)
	b[0] = BssapTypeDtap
	return nil
}

func ParseBssap(b []byte) (*bssmap.Bssmap, *dtap.Dtap, error) {
	x := BssapType(b[0])

	if x == BssapTypeBssMap {
		if int(b[1]) != len(b)-2 {
			return nil, nil, fmt.Errorf("error wrong bssmap data len(%d), %v", int(b[1]), hex.Dump(b))
		}

		m, err := bssmap.BssmapDecode(b)
		return m, nil, err
	}
	if x == BssapTypeDtap {
		// dlci := b[1]
		l := int(b[2])
		pl := b[3:]
		if l != len(pl) {
			return nil, nil, fmt.Errorf("error wrong dtap data len(%d), %v", int(b[1]), pl)
		}

		dt, err := dtap.DtapDecode(pl)
		return nil, dt, err
	}
	return nil, nil, fmt.Errorf("error unknown bssap payload type")
}
