package bssmap

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

var (
	//003b100b05010aa1a5010a010112035058a605080062f2080002177205080062f208000117710401072c01080809101000000087407c060aac10045bcc
	mockHoRQST = []byte{0x0, 0x3b, 0x10, 0xb, 0x5, 0x1, 0xa, 0xa1, 0xa5, 0x1, 0xa, 0x1, 0x1, 0x12, 0x3, 0x50, 0x58, 0xa6, 0x5, 0x8, 0x0, 0x62, 0xf2, 0x8, 0x0, 0x2, 0x17, 0x72, 0x5, 0x8, 0x0, 0x62, 0xf2, 0x8, 0x0, 0x1, 0x17, 0x71, 0x4, 0x1, 0x7, 0x2c, 0x1, 0x8, 0x8, 0x9, 0x10, 0x10, 0x0, 0x0, 0x0, 0x87, 0x40, 0x7c, 0x6, 0xa, 0xac, 0x10, 0x4, 0x5b, 0xcc}

	mockLuRQST = []byte{0x0, 0x23, 0x57, 0x5, 0x8, 0x0, 0x0, 0xf1, 0x10, 0x2, 0x59, 0x4, 0x4c, 0x17, 0x10, 0x5, 0x8, 0x30, 0x0, 0xf1, 0x10, 0x2, 0x58, 0x50, 0x5, 0xf4, 0x3a, 0x36, 0x2, 0x81, 0xe1, 0x7d, 0x4, 0x80, 0x83, 0x85, 0x57}

	mockAssRQST = []byte{0x0, 0x1e, byte(MSG_ASSIGNMENT_RQST), 0xb, 0x5, 0x1, 0xa, 0xa1, 0xa5, 0x1, 0x7c, 0x6, 0xc0, 0xa8, 0x1e, 0x57, 0x3e, 0xa2, 0x7d, 0x7, 0x83, 0xff, 0x57, 0x84, 0x3f, 0x7, 0x80, 0x7f, 0x23, 0x0, 0x0, 0x0}

	mockAoIpTranspAddr = []byte{byte(AOIP_TRASP_ADDR), 6, 127, 0, 0, 1, 0x1234 >> 8, 0x1234 & 0xFF}

	// Bssmap: Handover Required
	mockHoRQRD []byte = []byte{
		0, 25,
		0x11, 0x04, 0x01, 0x0c, 0x1b, 0x1a, 0x08, 0x00, 0x52, 0xf0, 0x70, 0xc7, 0x38, 0x00, 0x79, 0x31,
		0x18, 0x40, 0x01, 0x3a, 0x04, 0x02, 0x02, 0x01, 0x08}
)

type expectedIE struct {
	IEType BssmapIE
	B      IE
}

func Test_NewBssMapMinLen(t *testing.T) {
	t.Skip()
	if _, err := BssmapDecode([]byte{}); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_NewBssmap(t *testing.T) {
	t.Skip()
	msg, err := BssmapDecode(mockHoRQRD)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
}

func Test_NewBssmapIE(t *testing.T) {
	t.Skip()
	msg, err := BssmapDecode(mockHoRQRD)
	if err != nil {
		t.Fatal(err)
	}

	expIEs := []expectedIE{
		{CAUSE, IE{byte(CAUSE), 0x1, 0xc}},
		{RESPONSE_RQST, IE{byte(RESPONSE_RQST)}},
		{CELL_ID_LIST, IE{byte(CELL_ID_LIST), 0x08, 0x00, 0x52, 0xf0, 0x70, 0xc7, 0x38, 0x00, 0x79}},
		{CURRENT_CHANNEL_TYPE_1, IE{byte(CURRENT_CHANNEL_TYPE_1), 0x18}},
		{SPEECH_VERSION, IE{byte(SPEECH_VERSION), 0x1}},
		{OLD_BSS_TO_NEW_BSS_INF, IE{byte(OLD_BSS_TO_NEW_BSS_INF), 0x4, 0x2, 0x2, 0x1, 0x8}},
	}

	for i := range expIEs {
		expIEType := expIEs[i].IEType
		expB := expIEs[i].B
		ie, ok := msg.GetIE(expIEType)
		if !ok {
			t.Fatalf("IE %d was not found", expIEType)
		}
		actB := ie
		if !bytes.Equal(expB, actB) {
			t.Fatalf("found error of IE %q: expected bytes '% x' (len=%d) <> actual bytes '% x' (len=%d)", expIEType, expB, len(expB), actB, len(actB))
		}
	}
}

func Test_Replace(t *testing.T) {
	t.Skip()
	msg, err := BssmapDecode(mockAssRQST)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("\n\nReplace\n", msg)
	fmt.Println(hex.Dump(msg.Encode()))

	msg.Replace(mockAoIpTranspAddr)

	fmt.Println(msg)
	fmt.Println(hex.Dump(msg.Encode()))
}

func Test_Remove0(t *testing.T) {
	t.Skip()
	msg, err := BssmapDecode(mockAssRQST)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("\n\nRemove(leading)\n", msg)
	//fmt.Println(hex.Dump(msg.Encode()))

	msg.Remove(CHANNEL_TYPE)

	fmt.Println(msg)
	fmt.Println(hex.Dump(msg.Encode()))
}
func Test_Remove1(t *testing.T) {
	t.Skip()
	msg, err := BssmapDecode(mockAssRQST)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("\n\nRemove(middle)\n", msg)
	//fmt.Println(hex.Dump(msg.Encode()))

	msg.Remove(AOIP_TRASP_ADDR)

	fmt.Println(msg)
	fmt.Println(hex.Dump(msg.Encode()))
}

func Test_Remove2(t *testing.T) {
	t.Skip()
	msg, err := BssmapDecode(mockAssRQST)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("\n\nRemove(last)\n", msg)
	//fmt.Println(hex.Dump(msg.Encode()))

	msg.Remove(CALL_ID)

	fmt.Println(msg)
	fmt.Println(hex.Dump(msg.Encode()))
}

func Test_String(t *testing.T) {
	t.Skip()
	msg, err := BssmapDecode(mockLuRQST)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("LU RQST\t", msg)
	msg, err = BssmapDecode(mockHoRQST)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("HO RQST\t", msg)
}

var (
	mockCellIdList0 = IE([]byte{0x1a, 0x8, 0x0, 0x0, 0xf1, 0x10, 0x2, 0x58, 0x3, 0xf2})
	mockCellIdList1 = IE([]byte{0x1a, 0x5, 0x1, 0x0, 0x2, 0x17, 0x72})
	mockCellIdList5 = IE([]byte{0x1a, 0x3, 0x5, 0x2, 0x59}) // lac only, should return empty
)

func TestLists(t *testing.T) {
	t.Skip()
	i := mockCellIdList0
	fmt.Println("PD 0, lacs:", i.ListLACs(), "cells:", i.ListCIs())
	i = mockCellIdList1
	fmt.Println("PD 1, lacs:", i.ListLACs(), "cells:", i.ListCIs())
	i = mockCellIdList5
	fmt.Println("PD 5, lacs:", i.ListLACs(), "cells:", i.ListCIs())
}
