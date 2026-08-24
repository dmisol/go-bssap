package abisrsl

import (
	"errors"
	"github.com/dmisol/go-bssap/pkg/dtap"
	"fmt"
)

type RSL struct {
	MT
	IEs []IE
}

var (
	ErrInvalidLen = errors.New("invalid length")
	ErrUnknownIE  = errors.New("unknown IE")
)

func Parse(rsl []byte) (*RSL, error) {
	if len(rsl) < 2 {
		//fmt.Println("err0, rsl is", len(rsl))
		return nil, ErrInvalidLen
	}
	r := &RSL{
		MT:  MT(rsl[1]),
		IEs: make([]IE, 0),
	}
	offset := 2
	for offset < len(rsl) {
		tag := TAG(rsl[offset])
		f := tag.format()
		switch f {
		case 0:
			//fmt.Println("err1", offset, f, rsl, ErrUnknownIE)
			return r, ErrUnknownIE

		case -1:
			if offset+2 > len(rsl) {
				//fmt.Println("err2", offset, f, rsl, ErrUnknownIE)
				return r, ErrInvalidLen
			}
			l := int(rsl[offset+1])
			if offset+l+2 > len(rsl) {
				//fmt.Println("err3", offset, f, rsl, ErrUnknownIE)
				return r, ErrInvalidLen
			}
			ie := rsl[offset : offset+2+l]
			r.IEs = append(r.IEs, ie)
			offset += 2 + l

		case -2:
			if offset+3 > len(rsl) {
				//fmt.Println("err4", offset, f, rsl, ErrUnknownIE)
				return r, ErrInvalidLen
			}
			l := (int(rsl[offset+1]) << 8) + int(rsl[offset+2])
			if offset+l+3 > len(rsl) {
				//fmt.Println("err5", offset, f, rsl, ErrUnknownIE)
				return r, ErrInvalidLen
			}
			ie := rsl[offset : offset+3+l]
			r.IEs = append(r.IEs, ie)
			offset += 3 + l

		default:
			if offset+f > len(rsl) {
				//fmt.Println("err6", offset, f, rsl, ErrUnknownIE)
				return r, ErrInvalidLen
			}
			ie := rsl[offset : offset+f]
			r.IEs = append(r.IEs, ie)
			offset += f
		}
		//fmt.Println(r.MT, offset, len(r.IEs))
	}
	return r, nil
}

func Get(ies []IE, tag TAG) (IE, bool) {
	for _, v := range ies {
		if v.Tag() == tag {
			return v, true
		}
	}
	return nil, false
}

func getDtap(ies []IE, tag TAG) ([]byte, bool, error) {
    ie, found := Get(ies, tag)
    if !found {
        return nil, false, ErrWrongIE
    }

    if len(ie) < 2 {
        return nil, false, ErrInvalidLength
    }
	
    var data []byte
    var length int
    var hasL2PseudoLength bool = false

    switch tag {
    case IE_FULL_IMM_ASS_INFO:
        if len(ie) < 2 {
            return nil,  false, ErrInvalidLength
        }
        length = int(ie[1])
        if len(ie) < 2+length {
            return nil, false, ErrInvalidLength
        }
        data = ie[2 : 2+length]

    case IE_L3_INFO:
        if len(ie) < 3 {
            return nil,  false, ErrInvalidLength
        }
        length = int(ie[1])<<8 + int(ie[2])
        if len(ie) < 3+length {
            return nil, false, ErrInvalidLength
        }
        data = ie[3 : 3+length]
    default:
        return nil, false, ErrWrongIE
    }

	channelIE, found := Get(ies, IE_CHAN_NR)
    if found && len(channelIE) >= 2 {
        cbits := GetCbits(channelIE[1])
        hasL2PseudoLength = cbits.HasL2PseudoLength()
    }

    return data, hasL2PseudoLength, nil
}

func ExtractAndDecodeDTAPFromRSL(rsl *RSL) (*dtap.Dtap, error) {
    if rsl == nil {
        return nil, fmt.Errorf("RSL is nil")
    }

    if _, found := Get(rsl.IEs, IE_L3_INFO); found {
        ie, isl2, err := getDtap(rsl.IEs, IE_L3_INFO)
        if err != nil {
            return nil, fmt.Errorf("getDtap IE_L3_INFO: %w", err)
        }
        return dtap.DtapDecode(ie, dtap.SetL2PseudoLength(isl2))
    }
    
    if _, found := Get(rsl.IEs, IE_FULL_IMM_ASS_INFO); found {
        ie, isl2, err := getDtap(rsl.IEs, IE_FULL_IMM_ASS_INFO)
        if err != nil {
            return nil, fmt.Errorf("getDtap IE_FULL_IMM_ASS_INFO: %w", err)
        }
        return dtap.DtapDecode(ie, dtap.SetL2PseudoLength(isl2))
    }
    
    return nil, fmt.Errorf("%w, (tried: 0x%02X, 0x%02X)", ErrNoDtapIEFound, IE_L3_INFO, IE_FULL_IMM_ASS_INFO)
}