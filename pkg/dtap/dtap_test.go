package dtap

import (
    "testing"
)

//go test -v -run TestRealLocationUpdatingRequest 
//12 пакет
func TestRealLocationUpdatingRequest(t *testing.T) {

    rawData := []byte{
		0x05, 0x08, 0x70, 0x16, 0xf2,
        0x40, 0xff, 0xfe, 0x53, 0x08, 0x69, 0x21, 0x40, 0x00, 0x00,
        0x00, 0x34, 0x12,
    }

    t.Logf("Raw data length: %d bytes", len(rawData))
    dtap, err := DtapDecode(rawData, false)
    if err != nil {
        t.Fatalf("DtapDecode failed: %v", err)
    }

    t.Logf("=== HEADER ===")
    t.Logf("ProtocolDisc: 0x%02X (%v)", dtap.Header.ProtocolDisc, dtap.Header.ProtocolDisc)
    t.Logf("SkipInd: 0x%02X", dtap.Header.SkipInd)
    t.Logf("MsgType: %d", dtap.Header.MsgType)

    for i, ie := range dtap.IEs {
        t.Logf("IE[%d]: Tag=0x%02X, Length=%d, Value=%X", i, ie.Tag, len(ie.Value), ie.Value)
    }    
  
}

func TestIdentityRequest(t *testing.T) {

    rawData := []byte{
		0x05, 0x18, 0x01,
    }

    t.Logf("Raw data length: %d bytes", len(rawData))
    dtap, err := DtapDecode(rawData, false)
    if err != nil {
        t.Fatalf("DtapDecode failed: %v", err)
    }

    t.Logf("=== HEADER ===")
    t.Logf("ProtocolDisc: 0x%02X (%v)", dtap.Header.ProtocolDisc, dtap.Header.ProtocolDisc)
    t.Logf("SkipInd: 0x%02X", dtap.Header.SkipInd)
    t.Logf("MsgType: %d", dtap.Header.MsgType)

    for i, ie := range dtap.IEs {
        t.Logf("IE[%d]: Tag=0x%02X, Length=%d, Value=%X", i, ie.Tag, len(ie.Value), ie.Value)
    }
}

func TestIdentityResponse(t *testing.T) {

        rawData := []byte{
        0x05, 0x59, 0x08, 0x29, 0x05, 0x10, 0x39, 0x30, 0x55, 0x25,
        0x11,
    }

    t.Logf("Raw data length: %d bytes", len(rawData))
    dtap, err := DtapDecode(rawData, false)
    if err != nil {
        t.Fatalf("DtapDecode failed: %v", err)
    }

    t.Logf("=== HEADER ===")
    t.Logf("ProtocolDisc: 0x%02X (%v)", dtap.Header.ProtocolDisc, dtap.Header.ProtocolDisc)
    t.Logf("SkipInd: 0x%02X", dtap.Header.SkipInd)
    t.Logf("MsgType: %d", dtap.Header.MsgType)

    for i, ie := range dtap.IEs {
        t.Logf("IE[%d]: Tag=0x%02X, Length=%d, Value=%X", i, ie.Tag, len(ie.Value), ie.Value)
    }
}
//18
func TestServiceRequest(t *testing.T) {	

    rawData := []byte{
		0x5, 0x24, 0x1, 0x3, 0x53, 0x59, 0x86, 0x5, 0xf4, 0xbe, 0x8f, 0x54, 0xc1,
    }

    t.Logf("Raw data length: %d bytes", len(rawData))
    dtap, err := DtapDecode(rawData, false)
    if err != nil {
        t.Fatalf("DtapDecode failed: %v", err)
    }

    t.Logf("=== HEADER ===")
    t.Logf("ProtocolDisc: 0x%02X (%v)", dtap.Header.ProtocolDisc, dtap.Header.ProtocolDisc)
    t.Logf("SkipInd: 0x%02X", dtap.Header.SkipInd)
    t.Logf("MsgType: %d", dtap.Header.MsgType)

    for i, ie := range dtap.IEs {
        t.Logf("IE[%d]: Tag=0x%02X, Length=%d, Value=%X", i, ie.Tag, len(ie.Value), ie.Value)
    }
}

func TestRRChannelRelease(t *testing.T) {

    rawData := []byte{0x6, 0xd, 0x0, 0x77, 0x4, 0x70, 0xd1, 0x6b, 0x0}
    
    t.Logf("Raw data length: %d bytes", len(rawData))
    dtap, err := DtapDecode(rawData, false)
    if err != nil {
        t.Fatalf("DtapDecode failed: %v", err)
    }

    t.Logf("=== HEADER ===")
    t.Logf("ProtocolDisc: 0x%02X (%v)", dtap.Header.ProtocolDisc, dtap.Header.ProtocolDisc)
    t.Logf("SkipInd: 0x%02X", dtap.Header.SkipInd)
    t.Logf("MsgType: %d", dtap.Header.MsgType)

    for i, ie := range dtap.IEs {
        t.Logf("IE[%d]: Tag=0x%02X, Length=%d, Value=%X", i, ie.Tag, len(ie.Value), ie.Value)
    }    
}

func TestRRImmAss(t *testing.T) {

    rawData := []byte{0x2d, 0x6, 0x3f, 0x3, 0x41, 0xa3, 0x66, 0x7, 0xbc, 0x27, 0x0, 0x0, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b}
    
    t.Logf("Raw data length: %d bytes", len(rawData))
    dtap, err := DtapDecode(rawData, true)
    if err != nil {
        t.Fatalf("DtapDecode failed: %v", err)
    }

    t.Logf("=== HEADER ===")
    t.Logf("ProtocolDisc: 0x%02X (%v)", dtap.Header.ProtocolDisc, dtap.Header.ProtocolDisc)
    t.Logf("SkipInd: 0x%02X", dtap.Header.SkipInd)
    t.Logf("MsgType: %d", dtap.Header.MsgType)

    for i, ie := range dtap.IEs {
        t.Logf("IE[%d]: Tag=0x%02X, Length=%d, Value=%X", i, ie.Tag, len(ie.Value), ie.Value) //для MObile Alloc вернет пустой срез так как там LV с L=0 и без V
    }    
}
