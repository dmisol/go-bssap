package abisrsl

import (
	"errors"
)

type RSL struct {
	MT
	IEs []IE
}

var (
	ErrInvalidLen = errors.New("invalid lenght")
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
