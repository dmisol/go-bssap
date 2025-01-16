package bssmap

import (
	"bytes"
	"fmt"
	"testing"
)

// Bssmap: Handover Required
var BssmapMockBytes1 []byte = []byte{
	0x11, 0x04, 0x01, 0x0c, 0x1b, 0x1a, 0x08, 0x00, 0x52, 0xf0, 0x70, 0xc7, 0x38, 0x00, 0x79, 0x31,
	0x18, 0x40, 0x01, 0x3a, 0x04, 0x02, 0x02, 0x01, 0x08,
}

type expectedIE struct {
	IEType BssmapIE
	B      []byte
}

func Test_NewBssMapMinLen(t *testing.T) {
	if _, err := BssmapDecode([]byte{}); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_NewBssmap(t *testing.T) {
	msg, err := BssmapDecode(BssmapMockBytes1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
}

func Test_NewBssmapIE(t *testing.T) {
	msg, err := BssmapDecode(BssmapMockBytes1)
	if err != nil {
		t.Fatal(err)
	}

	expIEs := []expectedIE{
		{CAUSE, []byte{0x4, 0x1, 0xc}},
		{RESPONSE_RQST, []byte{0x1b}},
		{CELL_ID_LIST, []byte{0x1a, 0x08, 0x00, 0x52, 0xf0, 0x70, 0xc7, 0x38, 0x00, 0x79}},
		{CURRENT_CHANNEL_TYPE_1, []byte{0x31, 0x18}},
		{SPEECH_VERSION, []byte{0x40, 0x1}},
		{OLD_BSS_TO_NEW_BSS_INF, []byte{0x3a, 0x4, 0x2, 0x2, 0x1, 0x8}},
	}

	for i := range expIEs {
		expIEType := expIEs[i].IEType
		expB := expIEs[i].B
		ie, ok := msg.GetIE(expIEType)
		if !ok {
			t.Fatalf("IE %d was not found", expIEType)
		}
		actB := ie.Encode()
		if !bytes.Equal(expB, actB) {
			t.Fatalf("found error of IE %q: expected bytes '% x' (len=%d) <> actual bytes '% x' (len=%d)", expIEType, expB, len(expB), actB, len(actB))
		}
	}
}

// @todo: 
// func Test_NewBssmapBuild(t *testing.T) {
// msg := new(BssmapMsg)
// }