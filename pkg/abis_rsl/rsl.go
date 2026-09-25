package abisrsl

import (
	"errors"
	"fmt"

	"github.com/dmisol/go-bssap/pkg/dtap"
)

type RSL struct {
	MT
	IEs []IE
}

var (
	ErrInvalidLen = errors.New("invalid length")
)

//ip.access 0x40 0x48 0x49 0x4A 0x4B 0x4C 0x4D 0x50 0x51 0x52 0x53 0x54 0x55 0x56 0x57 0x58 0x60 0x61 0x62 0x70 0x71 0x72 0x73 0x74 0x75 0x76 0x77 0x78 0x79 0x7F

func Parse(rsl []byte) (*RSL, error) {
	if len(rsl) < 2 {
		//fmt.Println("err0, rsl is", len(rsl))
		return nil, ErrInvalidLen
	}
	r := &RSL{
		MT:  MT(rsl[1]),
		IEs: make([]IE, 0),
	}
	if rsl[0] == 0x7e { //ABIS_RSL_MDISC_IPACCESS
		return r, nil
	}
	offset := 2
	for offset < len(rsl) {
		tag := TAG(rsl[offset])
		f := tag.format()
		switch f {
		case 0:
			//fmt.Println("err1", offset, f, rsl, ErrUnknownIE)
			return r, fmt.Errorf("unknown IE tag: %s", tag.String())

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
			return nil, false, ErrInvalidLength
		}
		length = int(ie[1])
		if len(ie) < 2+length {
			return nil, false, ErrInvalidLength
		}
		data = ie[2 : 2+length]

	case IE_L3_INFO:
		if len(ie) < 3 {
			return nil, false, ErrInvalidLength
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
