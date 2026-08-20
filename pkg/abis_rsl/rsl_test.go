package abisrsl

import (
	"testing"
	"github.com/dmisol/go-bssap/pkg/dtap"
	"fmt"
)


func parseRSL(rawData []byte) (*RSL, error) {
    rsl, err := Parse(rawData)
    if err != nil {
        return nil, fmt.Errorf("parsing RSL: %w", err)
    }
    return rsl, nil
}

func printRSL(t *testing.T, rsl *RSL) {
    if rsl == nil {
        t.Log("RSL is nil")
        return
    }
    
    t.Logf("MT: 0x%02X", rsl.MT)
    for i, ie := range rsl.IEs {
        format := ie.Tag().format()
        switch {
        case format > 0:
            t.Logf("IE[%d]: Tag=0x%02X, Len=%d, Val=%X", i, ie.Tag(), len(ie), ie[1:])
        case format == -2:
            t.Logf("IE[%d]: Tag=0x%02X, Len=%d, Val=%X", i, ie.Tag(), len(ie), ie[3:])
        case format == -1:
            t.Logf("IE[%d]: Tag=0x%02X, Len=%d, Val=%X", i, ie.Tag(), len(ie), ie[2:])
        default:
            t.Logf("Error unsupported format %v\n", format)
        }
    }
}

func printDTAP(t *testing.T, dt1 *dtap.Dtap) {
    if dt1 == nil {
        t.Log("DTAP is nil")
        return
    }
    
    t.Logf("=== HEADER ===")
    t.Logf("ProtocolDisc: 0x%02X (%v)", dt1.Header.ProtocolDisc, dt1.Header.ProtocolDisc)
    t.Logf("SkipInd: 0x%02X", dt1.Header.SkipInd)
    t.Logf("MsgType: 0x%02x", dt1.Header.MsgType)
    
    if len(dt1.IEs) == 0 {
        t.Log("No IEs found")
        return
    }
    
    for i, ie := range dt1.IEs {
        t.Logf("IE[%d]: Tag=0x%02X, Length=%d, Value=%X", i, ie.Tag, len(ie.Value), ie.Value)
    }
}


func TestRSL_TFO(t *testing.T) {
    data := []byte{0x0c, 0x16, 0x01, 0x90, 0x2b, 0x17, 0x2d, 0x06, 0x3f, 0x03, 0x41, 
        0xe0, 0x3c, 0x4e, 0x48, 0x0a, 0x00, 0x00, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 
        0x2b, 0x2b, 0x2b, 0x2b, 0x2b}
    
    t.Log("=== Test RSL_TFO ===")
    t.Logf("Input data length: %d bytes", len(data))

    rsl, err := parseRSL(data)
    if err != nil {
        t.Fatalf("Parse RSL failed: %v", err)
    }
    printRSL(t, rsl)

    dt1, err := ExtractDTAPFromRSL(rsl)
    if err != nil {
        t.Logf("Extract DTAP failed: %v", err)
        return
    }

    printDTAP(t, dt1)
}



func TestRSL_RR_PagingResponse(t *testing.T) {
    data := []byte{0x2, 0x6, 0x1, 0x49, 0x2, 0x0, 0xb, 0x0, 0xd, 0x6, 0x27, 0x7, 0x3, 0x53, 0x59, 
		0x86, 0x5, 0xf4, 0x87, 0x3c, 0x74, 0xd2}
    
    t.Log("=== Test RSL_TFO ===")
    t.Logf("Input data length: %d bytes", len(data))

    rsl, err := parseRSL(data)
    if err != nil {
        t.Fatalf("Parse RSL failed: %v", err)
    }
    printRSL(t, rsl)

    dt1, err := ExtractDTAPFromRSL(rsl)
    if err != nil {
        t.Logf("Extract DTAP failed: %v", err)
        return
    }

    printDTAP(t, dt1)
}


func TestRSL_RR_AssComplete(t *testing.T) {
    data := []byte{0x3, 0x2, 0x1, 0xa, 0x2, 0x0, 0xb, 0x0, 0x3, 0x6, 0x29, 0x0}
    
    t.Log("=== Test RSL_TFO ===")
    t.Logf("Input data length: %d bytes", len(data))

    rsl, err := parseRSL(data)
    if err != nil {
        t.Fatalf("Parse RSL failed: %v", err)
    }
    printRSL(t, rsl)

    dt1, err := ExtractDTAPFromRSL(rsl)
    if err != nil {
        t.Logf("Extract DTAP failed: %v", err)
        return
    }

    printDTAP(t, dt1)
}

func TestRSL_MM_SerAcc(t *testing.T) {
    data := []byte{0x3, 0x1, 0x1, 0x41, 0x2, 0x0, 0xb, 0x0, 0x2, 0x5, 0x21}
    
    t.Log("=== Test RSL_TFO ===")
    t.Logf("Input data length: %d bytes", len(data))

    rsl, err := parseRSL(data)
    if err != nil {
        t.Fatalf("Parse RSL failed: %v", err)
    }
    printRSL(t, rsl)

    dt1, err := ExtractDTAPFromRSL(rsl)
    if err != nil {
        t.Logf("Extract DTAP failed: %v", err)
        return
    }

    printDTAP(t, dt1)
}

func TestRSL_RR_AssCmd(t *testing.T) {
    data := []byte{0x3, 0x1, 0x1, 0x41, 0x2, 0x0, 0xb, 0x0, 0x8, 0x6, 0x2e, 0xa, 0xe0, 0x32, 0x6, 0x63, 0x1}
    
    t.Log("=== Test RSL_TFO ===")
    t.Logf("Input data length: %d bytes", len(data))

    rsl, err := parseRSL(data)
    if err != nil {
        t.Fatalf("Parse RSL failed: %v", err)
    }
    printRSL(t, rsl)

    dt1, err := ExtractDTAPFromRSL(rsl)
    if err != nil {
        t.Logf("Extract DTAP failed: %v", err)
        return
    }

    printDTAP(t, dt1)
}

