package abisrsl

import "errors"

type MT byte

const (
	_ MT = iota
	/* Radio Link Layer Management */
	MT_DATA_REQ // 0x01
	MT_DATA_IND
	MT_ERROR_IND
	MT_EST_REQ
	MT_EST_CONF
	MT_EST_IND
	MT_REL_REQ
	MT_REL_CONF
	MT_REL_IND
	MT_UNIT_DATA_REQ
	MT_UNIT_DATA_IND // 0x0b
	MT_SUSP_REQ      // non-standard elements
	MT_SUSP_CONF
	MT_RES_REQ
	MT_RECON_REQ // 0x0f
	_
	// Common Channel Management / TRX Management
	MT_BCCH_INFO // 0x11
	MT_CCCH_LOAD_IND
	MT_CHAN_RQD
	MT_DELETE_IND
	MT_PAGING_CMD
	MT_IMMEDIATE_ASSIGN_CMD
	MT_SMS_BC_REQ
	MT_CHAN_CONF  // non-standard element
	MT_RF_RES_IND // 0x19
	MT_SACCH_FILL
	MT_OVERLOAD
	MT_ERROR_REPORT
	MT_SMS_BC_CMD
	MT_CBCH_LOAD_IND
	MT_NOT_CMD // 0x1f
	_
	/* Dedicate Channel Management */
	MT_CHAN_ACTIV // 0x21
	MT_CHAN_ACTIV_ACK
	MT_CHAN_ACTIV_NACK
	MT_CONN_FAIL
	MT_DEACTIVATE_SACCH
	MT_ENCR_CMD
	MT_HANDO_DET
	MT_MEAS_RES
	MT_MODE_MODIFY_REQ
	MT_MODE_MODIFY_ACK
	MT_MODE_MODIFY_NACK
	MT_PHY_CONTEXT_REQ
	MT_PHY_CONTEXT_CONF
	MT_RF_CHAN_REL
	MT_MS_POWER_CONTROL
	MT_BS_POWER_CONTROL // 0x30
	MT_PREPROC_CONFIG
	MT_PREPROC_MEAS_RES
	MT_RF_CHAN_REL_ACK
	MT_SACCH_INFO_MODIFY
	MT_TALKER_DET
	MT_LISTENER_DET
	MT_REMOTE_CODEC_CONF_REP
	MT_RTD_REP
	MT_PRE_HANDO_NOTIF
	MT_MR_CODEC_MOD_REQ
	MT_MR_CODEC_MOD_ACK
	MT_MR_CODEC_MOD_NACK
	MT_MR_CODEC_MOD_PER
	MT_TFO_REP
	MT_TFO_MOD_REQ // 0x3f
	_
	MT_LOCATION_INFO // 0x41
)

var (
	ErrInvalidLen = errors.New("invalid lenght")
	ErrUnknownIE  = errors.New("unknown IE")
)

func Parse(rsl []byte) (MT, []IE, error) {
	if len(rsl) == 0 {
		return 0, nil, ErrInvalidLen
	}
	mt := MT(rsl[0])
	var ies []IE
	offset := 1
	for offset < len(rsl) {
		tag := TAG(rsl[offset])
		f := tag.format()
		switch f {
		case 0:
			return mt, ies, ErrUnknownIE

		case -1:
			if offset+2 >= len(rsl) {
				return mt, ies, ErrInvalidLen
			}
			l := int(rsl[offset+1])
			if offset+l+2 >= len(rsl) {
				return mt, ies, ErrInvalidLen
			}
			ie := rsl[offset : offset+2+l]
			ies = append(ies, ie)
			offset += 2 + l

		case -2:
			if offset+3 >= len(rsl) {
				return mt, ies, ErrInvalidLen
			}
			l := (int(rsl[offset+1]) << 8) + int(rsl[offset+2])
			if offset+l+3 >= len(rsl) {
				return mt, ies, ErrInvalidLen
			}
			ie := rsl[offset : offset+3+l]
			ies = append(ies, ie)
			offset += 3 + l

		default:
			if offset+f >= len(rsl) {
				return mt, ies, ErrInvalidLen
			}
			ie := rsl[offset : offset+f]
			ies = append(ies, ie)
			offset += f
		}
	}
	return mt, ies, nil
}

func Get(ies []IE, tag TAG) (IE, bool) {
	for _, v := range ies {
		if v.Tag() == tag {
			return v, true
		}
	}
	return nil, false
}
