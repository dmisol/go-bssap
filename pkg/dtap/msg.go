package dtap

type PD_Type byte

// ETSI TS 124 007 V19.5.0 Table 11.2
const (
	PD_GROUP_CC PD_Type = 0x00
	PD_BCAST_CC PD_Type = 0x01
	PD_EPS_SMM  PD_Type = 0x02
	PD_CC       PD_Type = 0x03
	PD_GTTP     PD_Type = 0x04
	PD_MM       PD_Type = 0x05
	PD_RR       PD_Type = 0x06
	PD_EPS_MMM  PD_Type = 0x07
	PD_GPRS_MMM PD_Type = 0x08
	PD_SMS      PD_Type = 0x09
	PD_GPRS_SMM PD_Type = 0x0a
	PD_SS_NCL   PD_Type = 0x0b
	PD_LOC      PD_Type = 0x0c
	PD_EXTEND   PD_Type = 0x0e
	PD_TEST     PD_Type = 0x0f
)

func PD(b byte) PD_Type {
	return PD_Type(b & 0x0F)
}

type Msg_Type byte

// both types to use with stringer
type MM_Msg_Type Msg_Type
type RR_Msg_Type Msg_Type
type CC_Msg_Type Msg_Type

func Mt(b byte) (seq int, mt Msg_Type) {
	seq = int(b >> 6)
	mt = Msg_Type(b & 0x3F)
	return
}

type CAUSE byte

// TS 144.018 10.5.2.31
const (
	CauseNORMAL CAUSE = iota
	CauseAbnormalUnspec
	CauseAbnormalChanUnaccept
	CauseAbnormalTimerExpired
	CauseAbnormalNoRadio
)
