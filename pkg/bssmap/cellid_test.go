package bssmap

import (
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
	t.Skip()
	for _, v := range tst_cis {
		ci, lac, err := v.ParseCellId()
		if err != nil {
			t.Fatalf(err.Error())
		}
		fmt.Printf("ci: %x, lac: %x\n", ci, lac)
	}
}
