package dtap

import (
	"bytes"
	"testing"
)

var expected = []byte{0x86, 0xc8, 0x32, 0x5e, 0x38, 0x1d, 0x2}

func TestPack(t *testing.T) {
	x := PackCB("HexBSC")
	if !bytes.Equal(x, expected) {
		t.Fatal(expected, x)
	}
}
