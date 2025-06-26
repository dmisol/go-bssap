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
)

// format returns a length of the Information Element (IE)
// 1. If length > 0, it means that IE has FIXED length
// 2. If length = 0, it means that IE has VARIABLE length
// 3. If length = -1, it means that IE is unsupported
func (ie TAG) format() int {
	switch ie {
	case IE_CHAN_NR:
		return
	case IE_LINK_IDENT:
		return
	case IE_ACT_TYPE:
		return
	case IE_BS_POWER:
		return
	case IE_CHAN_IDENT:
		return
	case IE_CHAN_MODE:
		return
	case IE_ENCR_INFO:
		return
	case IE_FRAME_NUMBER:
		return
	case IE_HANDO_REF:
		return
	case IE_L1_INFO:
		return
	case IE_L3_INFO:
		return
	case IE_MS_IDENTITY:
		return
	case IE_MS_POWER:
		return
	case IE_PAGING_GROUP:
		return
	case IE_PAGING_LOAD:
		return
	case IE_PYHS_CONTEXT:
		return
	case IE_ACCESS_DELAY:
		return
	case IE_RACH_LOAD:
		return
	case IE_REQ_REFERENCE:
		return
	case IE_RELEASE_MODE:
		return
	case IE_RESOURCE_INFO:
		return
	case IE_RLM_CAUSE:
		return
	case IE_STARTNG_TIME:
		return
	case IE_TIMING_ADVANCE:
		return
	case IE_UPLINK_MEAS:
		return
	case IE_CAUSE:
		return
	case IE_MEAS_RES_NR:
		return
	case IE_MSG_ID:
		return
	case IE_SYSINFO_TYPE:
		return
	case IE_MS_POWER_PARAM:
		return
	case IE_BS_POWER_PARAM:
		return
	case IE_PREPROC_PARAM:
		return
	case IE_PREPROC_MEAS:
		return
	case IE_IMM_ASS_INFO:
		return
	case IE_SMSCB_INFO:
		return
	case IE_MS_TIMING_OFFSET:
		return
	case IE_ERR_MSG:
		return
	case IE_FULL_BCCH_INFO:
		return
	case IE_CHAN_NEEDED:
		return
	case IE_CB_CMD_TYPE:
		return
	case IE_SMSCB_MSG:
		return
	case IE_FULL_IMM_ASS_INFO:
		return
	case IE_SACCH_INFO:
		return
	case IE_CBCH_LOAD_INFO:
		return
	case IE_SMSCB_CHAN_INDICATOR:
		return
	case IE_GROUP_CALL_REF:
		return
	case IE_CHAN_DESC:
		return
	case IE_NCH_DRX_INFO:
		return
	case IE_CMD_INDICATOR:
		return
	case IE_EMLPP_PRIO:
		return
	case IE_UIC:
		return
	case IE_MAIN_CHAN_REF:
		return
	case IE_MR_CONFIG:
		return
	case IE_MR_CONTROL:
		return
	case IE_SUP_CODEC_TYPES:
		return
	case IE_CODEC_CONFIG:
		return
	case IE_RTD:
		return
	case IE_TFO_STATUS:
		return
	case IE_LLP_APDU:
	default:
		return -1
	}
}
