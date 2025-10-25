package abisrsl

type IE []byte
type TAG byte

func (ie IE) Tag() TAG {
	return TAG(ie[0])
}

const (
	_ TAG = iota
	IE_CHAN_NR
	IE_LINK_IDENT
	IE_ACT_TYPE
	IE_BS_POWER
	IE_CHAN_IDENT
	IE_CHAN_MODE
	IE_ENCR_INFO
	IE_FRAME_NUMBER
	IE_HANDO_REF
	IE_L1_INFO
	IE_L3_INFO
	IE_MS_IDENTITY
	IE_MS_POWER
	IE_PAGING_GROUP
	IE_PAGING_LOAD
	IE_PYHS_CONTEXT
	IE_ACCESS_DELAY
	IE_RACH_LOAD
	IE_REQ_REFERENCE
	IE_RELEASE_MODE
	IE_RESOURCE_INFO
	IE_RLM_CAUSE
	IE_STARTNG_TIME
	IE_TIMING_ADVANCE
	IE_UPLINK_MEAS
	IE_CAUSE
	IE_MEAS_RES_NR
	IE_MSG_ID
	_
	IE_SYSINFO_TYPE
	IE_MS_POWER_PARAM
	IE_BS_POWER_PARAM
	IE_PREPROC_PARAM
	IE_PREPROC_MEAS
	IE_IMM_ASS_INFO
	IE_SMSCB_INFO
	IE_MS_TIMING_OFFSET
	IE_ERR_MSG
	IE_FULL_BCCH_INFO
	IE_CHAN_NEEDED
	IE_CB_CMD_TYPE
	IE_SMSCB_MSG
	IE_FULL_IMM_ASS_INFO
	IE_SACCH_INFO
	IE_CBCH_LOAD_INFO
	IE_SMSCB_CHAN_INDICATOR
	IE_GROUP_CALL_REF
	IE_CHAN_DESC
	IE_NCH_DRX_INFO
	IE_CMD_INDICATOR
	IE_EMLPP_PRIO
	IE_UIC
	IE_MAIN_CHAN_REF
	IE_MR_CONFIG
	IE_MR_CONTROL
	IE_SUP_CODEC_TYPES
	IE_CODEC_CONFIG
	IE_RTD
	IE_TFO_STATUS
	IE_LLP_APDU

	// Osmocom specific
	RSL_IE_OSMO_REP_ACCH_CAP	= 0x60
	RSL_IE_OSMO_TRAINING_SEQUENCE	= 0x61
	RSL_IE_OSMO_TEMP_OVP_ACCH_CAP	= 0x62
	RSL_IE_OSMO_OSMUX_CID		= 0x63
	RSL_IE_OSMO_RTP_EXTENSIONS	= 0x64
)

// format() returns a format of the Information Element (IE)
// 1. If format >  0, it means that IE has FIXED length (equal to format)
// 2. If format =  0, it means that IE is unsupported/not exist anymore
// 3.    format = -1, IE has TLV structure, length is encoded with a single byte
// 4.    format = -2, IE has TLV structure, length is encoded with a pair of bytes
func (ie TAG) format() int {
	switch ie {
	case IE_CHAN_NR:
		return 2 // GSM 08.58 -> 9.3.1
	case IE_LINK_IDENT:
		return 2 // GSM 08.58 -> 9.3.2
	case IE_ACT_TYPE:
		return 2 // GSM 08.58 -> 9.3.3
	case IE_BS_POWER:
		return 2 // GSM 08.58 -> 9.3.4
	case IE_CHAN_IDENT:
		return -1 // GSM 08.58 -> 9.3.5
	case IE_CHAN_MODE:
		return -1 // GSM 08.58 -> 9.3.6
	case IE_ENCR_INFO:
		return -1 // GSM 08.58 -> 9.3.7
	case IE_FRAME_NUMBER:
		return 2 // GSM 08.58 -> 9.3.8
	case IE_HANDO_REF:
		return 2 // GSM 08.58 -> 9.3.9
	case IE_L1_INFO:
		return 3 // GSM 08.58 -> 9.3.10
	case IE_L3_INFO:
		return -2 // GSM 08.58 -> 9.3.11
	case IE_MS_IDENTITY:
		return -1 // GSM 08.58 -> 9.3.12
	case IE_MS_POWER:
		return 2 // GSM 08.58 -> 9.3.13
	case IE_PAGING_GROUP:
		return 2 // GSM 08.58 -> 9.3.14
	case IE_PAGING_LOAD:
		return 2 // GSM 08.58 -> 9.3.15
	case IE_PYHS_CONTEXT:
		return -1 // GSM 08.58 -> 9.3.16
	case IE_ACCESS_DELAY:
		return 2 // GSM 08.58 -> 9.3.17
	case IE_RACH_LOAD:
		return -1 // GSM 08.58 -> 9.3.18
	case IE_REQ_REFERENCE:
		return 4 // GSM 08.58 -> 9.3.19
	case IE_RELEASE_MODE:
		return 2 // GSM 08.58 -> 9.3.20
	case IE_RESOURCE_INFO:
		return -1 // GSM 08.58 -> 9.3.21
	case IE_RLM_CAUSE:
		return -1 // GSM 08.58 -> 9.3.22
	case IE_STARTNG_TIME:
		return 2 // GSM 08.58 -> 9.3.23
	case IE_TIMING_ADVANCE:
		return 2 // GSM 08.58 -> 9.3.24
	case IE_UPLINK_MEAS:
		return -1 // GSM 08.58 -> 9.3.25
	case IE_CAUSE:
		return -1 // GSM 08.58 -> 9.3.26
	case IE_MEAS_RES_NR:
		return 2 // GSM 08.58 -> 9.3.27
	case IE_MSG_ID:
		return 2 // GSM 08.58 -> 9.3.28
	case IE_SYSINFO_TYPE:
		return 2 // GSM 08.58 -> 9.3.30
	case IE_MS_POWER_PARAM:
		return -1 // GSM 08.58 -> 9.3.31
	case IE_BS_POWER_PARAM:
		return -1 // GSM 08.58 -> 9.3.32
	case IE_PREPROC_PARAM:
		return -1 // GSM 08.58 -> 9.3.33
	case IE_PREPROC_MEAS:
		return -1 // GSM 08.58 -> 9.3.34
	case IE_IMM_ASS_INFO:
		// Exist in Phase 1 (3.6.0), later Full below
		// Not exist in phase 2
		return 0 // GSM 08.58 -> (not exist)
	case IE_SMSCB_INFO:
		return 2 // GSM 08.58 -> 9.3.36
	case IE_MS_TIMING_OFFSET:
		return 2 // GSM 08.58 -> 9.3.37
	case IE_ERR_MSG:
		return -1 // GSM 08.58 -> 9.3.38
	case IE_FULL_BCCH_INFO:
		return -1 // GSM 08.58 -> 9.3.39
	case IE_CHAN_NEEDED:
		return 2 // GSM 08.58 -> 9.3.40
	case IE_CB_CMD_TYPE:
		return 2 // GSM 08.58 -> 9.3.41
	case IE_SMSCB_MSG:
		return -1 // GSM 08.58 -> 9.3.42
	case IE_FULL_IMM_ASS_INFO:
		return -1 // GSM 08.58 -> 9.3.35
	case IE_SACCH_INFO:
		return -1 // GSM 08.58 -> 9.3.29
	case IE_CBCH_LOAD_INFO:
		return 2 // GSM 08.58 -> 9.3.43
	case IE_SMSCB_CHAN_INDICATOR:
		return 2 // GSM 08.58 -> 9.3.44
	case IE_GROUP_CALL_REF:
		return -1 // GSM 08.58 -> 9.3.45
	case IE_CHAN_DESC:
		return -1 // GSM 08.58 -> 9.3.46
	case IE_NCH_DRX_INFO:
		return -1 // GSM 08.58 -> 9.3.47
	case IE_CMD_INDICATOR:
		return -1 // GSM 08.58 -> 9.3.48
	case IE_EMLPP_PRIO:
		return 2 // GSM 08.58 -> 9.3.49
	case IE_UIC:
		return -1 // GSM 08.58 -> 9.3.50
	case IE_MAIN_CHAN_REF:
		return 2 // GSM 08.58 -> 9.3.51
	case IE_MR_CONFIG:
		return -1 // GSM 08.58 -> 9.3.52
	case IE_MR_CONTROL:
		return 2 // GSM 08.58 -> 9.3.53
	case IE_SUP_CODEC_TYPES:
		return -1 // GSM 08.58 -> 9.3.54
	case IE_CODEC_CONFIG:
		return -1 // GSM 08.58 -> 9.3.55
	case IE_RTD:
		return 2 // GSM 08.58 -> 9.3.56
	case IE_TFO_STATUS:
		return 2 // GSM 08.58 -> 9.3.57
	case IE_LLP_APDU:
		return -1 // GSM 08.58 -> 9.3.58

	//osmocom specific
	case RSL_IE_OSMO_REP_ACCH_CAP:
		return -1
	case RSL_IE_OSMO_TRAINING_SEQUENCE:
		return -1
	case RSL_IE_OSMO_TEMP_OVP_ACCH_CAP:
		return -1
	case RSL_IE_OSMO_OSMUX_CID:
		return -1
	case RSL_IE_OSMO_RTP_EXTENSIONS:
		return -1
	default:
		return 0
	}
}
