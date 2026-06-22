package dtap

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type IdentityType byte
type TMSI uint32
type IMSI string

const (
	ID_NONE IdentityType = iota
	ID_IMSI
	ID_IMEI
	ID_IMEISV
	ID_TMSI
	// ID_TMGI
)

func (i IdentityType) String() string {
	switch i {
	case ID_NONE:
		return "NONE"
	case ID_IMSI:
		return "IMSI"
	case ID_IMEI:
		return "IMEI"
	case ID_IMEISV:
		return "IMSISV"
	case ID_TMSI:
		return "TMSI"
	default:
		return "UNKNOWN"
	}
}

func identiny(b []byte) (it IdentityType, tmsi TMSI, imsi IMSI) {
	// fmt.Println(hex.EncodeToString(b))
	if len(b) < 2 {
		return
	}

	it = IdentityType(b[0] & 0x07)
	switch it {
	case ID_TMSI:
		if len(b) != 5 {
			return
		}
		tmsi = TMSI(binary.BigEndian.Uint32(b[1:]))

	case ID_IMSI:
		var s strings.Builder
		for j := 0; j < len(b); j++ {
			if j == 0 {
				s.WriteString(fmt.Sprintf("%d", (b[j]>>4)&0x0F))
			} else {
				s.WriteString(fmt.Sprintf("%d%d", b[j]&0x0F, (b[j]>>4)&0x0F))
			}
		}
		imsi = IMSI(s.String())

		if b[0]&0x08 == 0 {
			imsi = imsi[:len(imsi)-1]
		}
	}
	return
}

// in fact, we just worry abount IMSI and TMSI only, to route between MOCN cores or MSC pools
func InitialL3Info(b []byte) (mt Msg_Type, it IdentityType, tmsi TMSI, imsi IMSI) {
	if len(b) < 2 {
		return
	}

	_, mt = Mt(b[1])

	switch PD(b[0]) {
	case PD_RR:
		switch mt {
		case MSG_RR_PAG_RESP:
			// TS 24.008 Table 9.25/GSM 04.08: PAGING RESPONSE message content
			if len(b) < 8+2 {
				mt = 0
				return
			}
			l := b[7]
			if len(b) < 8+int(l) {
				mt = 0
				return
			}
			it, tmsi, imsi = identiny(b[8 : 8+l])
			return
		}
	case PD_MM:
		switch mt {
		case MSG_MM_LOC_UPD_REQUEST:
			// TS 24.008 Table 9.2.17
			if len(b) < 11 {
				mt = 0
				return
			}
			l := b[9]
			if len(b) < 10+int(l) {
				mt = 0
				return
			}
			it, tmsi, imsi = identiny(b[10 : 10+l])
			return
		case MSG_MM_CM_REEST_REQ:
			// Table 9.42/GSM 04.08: CM RE-ESTABLISHMENT REQUEST message content
			if len(b) < 8+6 {
				mt = 0
				return
			}
			l := b[7]
			if len(b) < 8+int(l) {
				mt = 0
				return
			}
			it, tmsi, imsi = identiny(b[8 : 8+l])
			return
		case MSG_MM_CM_SERV_REQ:
			// Table 9.45/GSM 04.08: CM SERVICE REQUEST message content
			if len(b) < 8+6 {
				mt = 0
				return
			}
			l := b[7]
			if len(b) < 8+int(l) {
				mt = 0
				return
			}
			it, tmsi, imsi = identiny(b[8 : 8+l])
			return
		case MSG_MM_IMSI_DETACH_IND:
			// Table 9.48/GSM 04.08: IMSI DETACH INDICATION message content
			if len(b) < 5 {
				mt = 0
				return
			}
			l := b[3]
			if len(b) < 4+int(l) {
				mt = 0
				return
			}
			it, tmsi, imsi = identiny(b[4 : 4+l])
			return
		}
	}
	return
}
