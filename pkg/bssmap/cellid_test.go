package bssmap

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var (
	tst_cis = []IE{
		[]byte{0x5, 0x8, 0x0, 0x62, 0xf2, 0x8, 0x0, 0x2, 0x17, 0x72},
		[]byte{0x5, 0x5, 0x1, 0x0, 0x1, 0x17, 0x71},
	}
)

func TestCI(t *testing.T) {
	v, _ := hex.DecodeString("05080016221400011b39")
	tst_cis = append(tst_cis, v)

	for _, v := range tst_cis {
		mcc, mnc, ci, lac, err := v.ParseCellId()
		if err != nil {
			t.Error(err)
		}
		x := NewBssmap(MSG_CMPL_LAYER_3, v)
		fmt.Println(x)
		fmt.Printf("mcc: %d, mnc: %d, ci: %x, lac: %x\n", mcc, mnc, ci, lac)
	}
}
