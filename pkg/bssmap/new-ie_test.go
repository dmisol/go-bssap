package bssmap

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewIe(t *testing.T) {
	ie := NewCellId(0x1234, 0xABCD)
	fmt.Println(hex.EncodeToString([]byte(ie)))

	lacs := []int{0xA111, 0xA112}
	cis := []int{0xC111, 0xC112}
	ie = NewCellIdList(612, 4, lacs, cis)
	fmt.Println(hex.EncodeToString([]byte(ie)))
	ie = NewCellIdList(0, 0, lacs, cis)
	fmt.Println(hex.EncodeToString([]byte(ie)))
}
