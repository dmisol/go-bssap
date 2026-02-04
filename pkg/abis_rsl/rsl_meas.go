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
	LevF  uint32
	LevS  uint32
	QualF uint32
	QualS uint32
	N     uint32 // todo: 7 -> 0
}
type NCell struct {
	RxLev   uint32
	FreqIdx uint32
	BSIC    uint32
}

const (
	levMask  = 0x3F
	freqMask = 0x1F
	bsicMask = 0x3F
)

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

	u := binary.BigEndian.Uint32(ie[3+2:])
	n := (u >> 6) & 0x07
	if n == 7 {
		return
	}
	sc = &SCell{}
	ncs = make([]*NCell, 0)
	sc.Dtx = u&0x40000000 == 0x40000000
	sc.Valid = u&0x400000 == 0x400000
	sc.LevF = (u >> 24) & 0x3F
	sc.LevS = (u >> 16) & 0x3F
	sc.QualF = (u >> 12) & 0x07
	sc.QualS = (u >> 9) & 0x07
	sc.N = n
	if sc.N == 0 {
		return
	}

	nc := &NCell{ // ncell 1
		RxLev: u & levMask,
	}

	u = binary.BigEndian.Uint32(ie[3+2+4:]) // 10.5.2.20 Measurement Results, Octets 6+
	nc.FreqIdx = u >> (24 + 3) & freqMask
	nc.BSIC = u >> (16 + 5) & bsicMask
	ncs = append(ncs, nc)
	if len(ncs) >= int(sc.N) {
		return
	}

	nc = &NCell{ // ncell 2
		RxLev:   (u >> 15) & levMask,
		FreqIdx: (u >> 10) & freqMask,
		BSIC:    (u >> 4) & bsicMask,
	}
	ncs = append(ncs, nc)
	if len(ncs) >= int(sc.N) {
		return
	}

	nc = &NCell{ // ncell 3
		RxLev: (u << 2) & 0x3C,
	}
	u = binary.BigEndian.Uint32(ie[3+2+4:]) // 10.5.2.20 Measurement Results, Octets 10+
	nc.RxLev |= (u >> 30) & 0x3
	nc.FreqIdx = (u >> (24 + 1)) & freqMask
	nc.BSIC = (u >> (16 + 3)) & bsicMask
	ncs = append(ncs, nc)
	if len(ncs) >= int(sc.N) {
		return
	}

	nc = &NCell{ // ncell 4
		RxLev:   (u >> (8 + 5)) & levMask,
		FreqIdx: (u >> 8) & freqMask,
		BSIC:    (u >> 2) & bsicMask,
	}
	ncs = append(ncs, nc)
	if len(ncs) >= int(sc.N) {
		return
	}

	nc = &NCell{ // ncell 5
		RxLev: (u << 4) & 0xF0,
	}
	u = binary.BigEndian.Uint32(ie[3+2+8:]) // 10.5.2.20 Measurement Results, Octets 14+
	nc.RxLev |= (u >> (24 + 4)) & 0xF
	nc.FreqIdx = (u >> (16 + 7)) & freqMask
	nc.BSIC = (u >> (16 + 1)) & bsicMask
	ncs = append(ncs, nc)
	if len(ncs) >= int(sc.N) {
		return
	}

	nc = &NCell{ // ncell 6
		RxLev:   (u >> (8 + 3)) & levMask,
		FreqIdx: (u >> 6) & freqMask,
		BSIC:    u & bsicMask,
	}
	ncs = append(ncs, nc)

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
