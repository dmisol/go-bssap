package dtap

import (
    "testing"
)

//go test -v -run TestRealLocationUpdatingRequest 
//12 пакет
func decodeAndLogDTAP(t *testing.T, rawData []byte, opts ...Option) {
    
    t.Logf("Raw data length: %d bytes", len(rawData))
    dtap, err := DtapDecode(rawData, opts...)
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
func TestMMRealLocationUpdatingRequest(t *testing.T) {
    rawData := []byte{
        0x05, 0x08, 0x70, 0x16, 0xf2,
        0x40, 0xff, 0xfe, 0x53, 0x08, 0x69, 0x21, 0x40, 0x00, 0x00,
        0x00, 0x34, 0x12,
    }
    decodeAndLogDTAP(t, rawData)
}

func TestMMIdentityRequest(t *testing.T) {
    rawData := []byte{
        0x05, 0x18, 0x01,
    }
    decodeAndLogDTAP(t, rawData)
}

func TestMMIdentityResponse(t *testing.T) {
    rawData := []byte{
        0x05, 0x59, 0x08, 0x29, 0x05, 0x10, 0x39, 0x30, 0x55, 0x25,
        0x11,
    }
    decodeAndLogDTAP(t, rawData)
}

func TestMMServiceRequest(t *testing.T) {
    rawData := []byte{
        0x5, 0x24, 0x1, 0x3, 0x53, 0x59, 0x86, 0x5, 0xf4, 0xbe, 0x8f, 0x54, 0xc1,
    }
    decodeAndLogDTAP(t, rawData)
}

func TestRRChannelRelease(t *testing.T) {
    rawData := []byte{0x6, 0xd, 0x0, 0x77, 0x4, 0x70, 0xd1, 0x6b, 0x0}
    decodeAndLogDTAP(t, rawData)
}

func TestRRImmAss(t *testing.T) {
    rawData := []byte{
        0x2d, 0x6, 0x3f, 0x3, 0x41, 0xa3, 0x66, 0x7, 0xbc,
        0x27, 0x0, 0x0, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b,
        0x2b, 0x2b, 0x2b, 0x2b, 0x2b,
    }
    decodeAndLogDTAP(t, rawData, WithL2PseudoLength())//L2 Pseudo Length
}

func TestRRAssCmd(t *testing.T) {
    rawData := []byte{0x6, 0x2e, 0xa, 0xe2, 0x15, 0x7, 0x63, 0x1}
    decodeAndLogDTAP(t, rawData)
}

func TestRRHandoCmd(t *testing.T) {
    rawData := []byte{0x6, 0x2b, 0x3f, 0x3c, 0xb, 0xe0, 0x3c, 0x2, 0xe, 0xd0, 0x63, 0x1, 0x90}
    decodeAndLogDTAP(t, rawData)
}