package abisrsl

import (
	"encoding/binary"
	"errors"
)

var (
	ErrWrongIE    = errors.New("invalid IE for request")
	ErrWrongPload = errors.New("unexpected l3 payload")
)

// fixme: there should be a way to do it better

// from IE_L1_INFO -> MS power and TA
func (r *RSL) DecodeL1Info() (power int, ta int, err error) {
	ie, ok := Get(r.IEs, IE_L1_INFO)
	if !ok {
		err = ErrWrongIE
		return
	}
	if len(ie) != 3 {
		err = ErrInvalidLen
		return
	}
	power = int(ie[1]) >> 3
	ta = int(ie[2]) >> 2
	return
}

// todo:
type SCell struct {
	Dtx   bool
	Valid bool
	LevF  uint64
	LevS  uint64
	QualF uint64
	QualS uint64
	N     uint64 // todo: 7 -> 0
}
type NCell struct {
	RxLev uint64
	Freq  uint64
	BSIC  uint64
}

// from IE_L3_INFO -> DTAP -> ...
func (r *RSL) DecodeDownlinkMeas() (sc *SCell, ncs []*NCell, err error) {
	ie, ok := Get(r.IEs, IE_L3_INFO)
	if !ok {
		err = ErrWrongIE
		return
	}
	if len(ie) != 21 {
		err = ErrInvalidLen
		return
	}
	// fixme!
	if ie[3+1] != 0x15 { // Meas Rep in L3 info
		err = ErrWrongPload
		return
	}

	u := binary.BigEndian.Uint64(ie[3+2:])
	sc = &SCell{}
	ncs = make([]*NCell, 0)

	sc.Dtx = u&0x40000000 == 0x40000000
	sc.Valid = u&0x400000 == 0x400000
	sc.LevF = (u >> 24) & 0x3F
	sc.LevS = (u >> 16) & 0x3F
	sc.QualF = (u >> 12) & 0x07
	sc.QualS = (u >> 9) & 0x07
	sc.N = (u >> 6) & 0x07
	if sc.N == 7 {
		sc.N = 0
	}

	nc := &NCell{
		RxLev: u & 0x3F,
	}

	// todo: etc

	ncs = append(ncs, nc)
	// todo: etc

	return
}

// from IE_UPLINK_MEAS
func (r *RSL) DecodeUplinkMeas() (dtx bool, lev_full int, lev_sub int, qual_full int, qual_sub int, err error) {
	ie, ok := Get(r.IEs, IE_UPLINK_MEAS)
	if !ok {
		err = ErrWrongIE
		return
	}
	if len(ie) < 5 {
		err = ErrInvalidLen
		return
	}

	dtx = ie[2]&0x40 == 0x40

	lev_full = int(ie[2]) & 0x3F
	lev_sub = int(ie[3]) & 0x3F

	qual_full = (int(ie[4]) >> 3) & 0x07
	qual_sub = int(ie[4]) & 0x07

	return
}
