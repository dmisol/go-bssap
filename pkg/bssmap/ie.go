package bssmap

import "fmt"

type BssmapIE byte

const (
	CIC                                 BssmapIE = 1 // Circuit Identity Code
	RSVD_0                              BssmapIE = 2
	RESRC_AVAILABLE                     BssmapIE = 3
	CAUSE                               BssmapIE = 4
	CELL_ID                             BssmapIE = 5
	PRIORITY                            BssmapIE = 6
	LAYER_3_HEADER_INF                  BssmapIE = 7
	IMSI                                BssmapIE = 8
	TMSI                                BssmapIE = 9
	ENCRYPTION_INF                      BssmapIE = 10
	CHANNEL_TYPE                        BssmapIE = 11
	PERIODICITY                         BssmapIE = 12
	EXTENDED_RESRC_INDICATOR            BssmapIE = 13
	NUMBER_OF_MSS                       BssmapIE = 14
	RSVD_1                              BssmapIE = 15
	RSVD_2                              BssmapIE = 16
	RSVD_3                              BssmapIE = 17
	CLASSMARK_INF_T2                    BssmapIE = 18
	CLASSMARK_INF_T3                    BssmapIE = 19
	INTERFERENCE_BAND_TO_USE            BssmapIE = 20
	RR_CAUSE                            BssmapIE = 21
	RSVD_4                              BssmapIE = 22
	LAYER_3_INF                         BssmapIE = 23
	DLCI                                BssmapIE = 24
	DOWNLINK_DTX_FLAG                   BssmapIE = 25
	CELL_ID_LIST                        BssmapIE = 26 // Cell Identifier List
	RESPONSE_RQST                       BssmapIE = 27 // Response Request
	RESRC_IND_METHOD                    BssmapIE = 28
	CLASSMARK_INF_TYPE_1                BssmapIE = 29
	CIC_LIST                            BssmapIE = 30
	DIAGNOSTIC                          BssmapIE = 31
	LAYER_3_MESSAGE_CONTENTS            BssmapIE = 32
	CHOSEN_CHANNEL                      BssmapIE = 33
	TOTAL_RESRC_ACCESSIBLE              BssmapIE = 34
	CIPHER_RESPONSE_MODE                BssmapIE = 35
	CHANNEL_NEEDED                      BssmapIE = 36
	TRACE_TYPE                          BssmapIE = 37
	TRIGGERID                           BssmapIE = 38
	TRACE_REFERENCE                     BssmapIE = 39
	TRANSACTIONID                       BssmapIE = 40
	MOBILE_IDENTITY                     BssmapIE = 41
	OMCID                               BssmapIE = 42
	FORWARD_INDICATOR                   BssmapIE = 43
	CHOSEN_ENCR_ALG                     BssmapIE = 44
	CIRCUIT_POOL                        BssmapIE = 45
	CIRCUIT_POOL_LIST                   BssmapIE = 46
	TIME_IND                            BssmapIE = 47
	RESRC_SITUATION                     BssmapIE = 48
	CURRENT_CHANNEL_TYPE_1              BssmapIE = 49 // Current Channel type 1
	QUEUEING_INDICATOR                  BssmapIE = 50
	SPEECH_VERSION                      BssmapIE = 64 // Speech Version
	ASSIGNMENT_REQUIREMENT              BssmapIE = 51
	TALKER_FLAG                         BssmapIE = 53
	CONNECTION_RELEASE_RQSTED           BssmapIE = 54
	GROUP_CALL_REFERENCE                BssmapIE = 55
	EMLPP_PRIORITY                      BssmapIE = 56
	CONFIG_EVO_INDI                     BssmapIE = 57
	OLD_BSS_TO_NEW_BSS_INF              BssmapIE = 58 // Old BSS to New BSS Information
	LSA_ID                              BssmapIE = 59
	LSA_ID_LIST                         BssmapIE = 60
	LSA_INF                             BssmapIE = 61
	LCS_QOS                             BssmapIE = 62
	LSA_ACCESS_CTRL_SUPPR               BssmapIE = 63
	LCS_PRIORITY                        BssmapIE = 67
	LOCATION_TYPE                       BssmapIE = 68
	LOCATION_ESTIMATE                   BssmapIE = 69
	POSITIONING_DATA                    BssmapIE = 70
	LCS_CAUSE                           BssmapIE = 71
	LCS_CLIENT_TYPE                     BssmapIE = 72
	APDU                                BssmapIE = 73
	NETW_ELEMENT_IDENTITY               BssmapIE = 74
	GPS_ASSISTANCE_DATA                 BssmapIE = 75
	DECIPHERING_KEYS                    BssmapIE = 76
	RET_ERROR_RQST                      BssmapIE = 77
	RET_ERROR_CAUSE                     BssmapIE = 78
	SEGMENTATION                        BssmapIE = 79
	SERVICE_HANDOVER                    BssmapIE = 80
	SRC_RNC_TO_TARG_RNC_TRANSP_UMTS     BssmapIE = 81
	SRC_RNC_TO_TARG_RNC_TRANSP_CDMA2000 BssmapIE = 82
	RSVD_5                              BssmapIE = 65
	RSVD_6                              BssmapIE = 66
	GERAN_CLASSMARK                     BssmapIE = 0x53
	GERAN_BSC_CONTAINER                 BssmapIE = 0x54
	NEW_BSS_TO_OLD_BSS_INFO             BssmapIE = 0x61
	GSM0800_IE_INTER_SYSTEM_INFO        BssmapIE = 0x63
	SNA_ACCESS_INFO                     BssmapIE = 0x64
	VSTK_RAND_INFO                      BssmapIE = 0x65
	VSTK_INFO                           BssmapIE = 0x66
	PAGING_INFO                         BssmapIE = 0x67
	IMEI                                BssmapIE = 0x68
	VELOCITY_ESTIMATE                   BssmapIE = 0x55
	VGCS_FEATURE_FLAGS                  BssmapIE = 0x69
	TALKER_PRIORITY                     BssmapIE = 0x6a
	EMERGENCY_SET_IND                   BssmapIE = 0x6b
	TALKER_IDENTITY                     BssmapIE = 0x6c
	CELL_ID_LIST_SEGMENT                BssmapIE = 0x6d
	SMS_TO_VGCS                         BssmapIE = 0x6e
	VGCS_TALKER_MODE                    BssmapIE = 0x6f
	VGCS_VBS_CELL_STATUS                BssmapIE = 0x70
	CELL_ID_LIST_SEG_EST_CELLS          BssmapIE = 0x71
	CELL_ID_LIST_SEG_CELLS_TBE          BssmapIE = 0x72
	CELL_ID_LIST_SEG_REL_CELLS          BssmapIE = 0x73
	CELL_ID_LIST_SEG_NE_CELLS           BssmapIE = 0x74
	GANSS_ASSISTANCE_DATA               BssmapIE = 0x75
	GANSS_POSITIONING_DATA              BssmapIE = 0x76
	GANSS_LOCATION_TYPE                 BssmapIE = 0x77
	APP_DATA                            BssmapIE = 0x78
	DATA_IDENTITY                       BssmapIE = 0x79
	APP_DATA_INFO                       BssmapIE = 0x7a
	MSISDN                              BssmapIE = 0x7b
	AOIP_TRASP_ADDR                     BssmapIE = 0x7c
	SPEECH_CODEC_LIST                   BssmapIE = 0x7d
	SPEECH_CODEC                        BssmapIE = 0x7e
	CALL_ID                             BssmapIE = 0x7f
	CALL_ID_LIST                        BssmapIE = 0x80
	A_IF_SEL_FOR_RESET                  BssmapIE = 0x81
	KC_128                              BssmapIE = 0x83
	CSG_ID                              BssmapIE = 0x84
	REDIR_ATTEMPT_FLAG                  BssmapIE = 0x85
	REROUTE_REJ_CAUSE                   BssmapIE = 0x86
	SEND_SEQ_NUM                        BssmapIE = 0x87
	REROUTE_COMPL_OUTCOME               BssmapIE = 0x88
	GLOBAL_CALL_REF                     BssmapIE = 0x89
	LCLS_CONFIG                         BssmapIE = 0x8a
	LCLS_CONN_STATUS_CTRL               BssmapIE = 0x8b
	LCLS_CORR_NOT_NEEDED                BssmapIE = 0x8c
	LCLS_BSS_STATUS                     BssmapIE = 0x8d
	LCLS_BREAK_REQ                      BssmapIE = 0x8e
	CSFB_IND                            BssmapIE = 0x8f
	CS_TO_PS_SRVCC                      BssmapIE = 0x90
	SRC_ENB_TO_TGT_ENB_TRANSP           BssmapIE = 0x91
	CS_TO_PS_SRVCC_IND                  BssmapIE = 0x92
	CN_TO_MS_TRANSP_INFO                BssmapIE = 0x93
	SELECTED_PLMN_ID                    BssmapIE = 0x94
	LAST_USED_EUTRAN_PLMN_ID            BssmapIE = 0x95
	OLD_LAI                             BssmapIE = 0x96
	ATTACH_INDICATOR                    BssmapIE = 0x97
	SELECTED_OPERATOR                   BssmapIE = 0x98
	PS_REGISTERED_OPERATOR              BssmapIE = 0x99
	CS_REGISTERED_OPERATOR              BssmapIE = 0x9a
)

func (ie BssmapIE) Format() (isVariable bool, fixed int) {
	switch ie {
	case RESPONSE_RQST:
		return false, 1
	case CURRENT_CHANNEL_TYPE_1, SPEECH_VERSION:
		return false, 2
	default:
		return true, 0
	}
}

func (ie BssmapIE) String() string {
	switch ie {
	case CIC:
		return "Circuit Identity Code"
	// RSVD_0
	// RESRC_AVAILABLE
	case CAUSE:
		return "Cause"
	// CELL_ID
	// PRIORITY
	// LAYER_3_HEADER_INF
	// IMSI
	// TMSI
	// ENCRYPTION_INF
	// CHANNEL_TYPE
	// PERIODICITY
	// EXTENDED_RESRC_INDICATOR
	// NUMBER_OF_MSS
	// RSVD_1
	// RSVD_2
	// RSVD_3
	// CLASSMARK_INF_T2
	// CLASSMARK_INF_T3
	// INTERFERENCE_BAND_TO_USE
	// RR_CAUSE
	// RSVD_4
	// LAYER_3_INF
	// DLCI
	// DOWNLINK_DTX_FLAG
	case CELL_ID_LIST:
		return "Cell Identifier List"
	case RESPONSE_RQST:
		return "Response Request"
	// RESRC_IND_METHOD
	// CLASSMARK_INF_TYPE_1
	// CIC_LIST
	// DIAGNOSTIC
	// LAYER_3_MESSAGE_CONTENTS
	// CHOSEN_CHANNEL
	// TOTAL_RESRC_ACCESSIBLE
	// CIPHER_RESPONSE_MODE
	// CHANNEL_NEEDED
	// TRACE_TYPE
	// TRIGGERID
	// TRACE_REFERENCE
	// TRANSACTIONID
	// MOBILE_IDENTITY
	// OMCID
	// FORWARD_INDICATOR
	// CHOSEN_ENCR_ALG
	// CIRCUIT_POOL
	// CIRCUIT_POOL_LIST
	// TIME_IND
	// RESRC_SITUATION
	case CURRENT_CHANNEL_TYPE_1:
		return "Current Channel type 1"
	// QUEUEING_INDICATOR
	case SPEECH_VERSION:
		return "Speech Version"
	// ASSIGNMENT_REQUIREMENT
	// TALKER_FLAG
	// CONNECTION_RELEASE_RQSTED
	// GROUP_CALL_REFERENCE
	// EMLPP_PRIORITY
	// CONFIG_EVO_INDI
	case OLD_BSS_TO_NEW_BSS_INF:
		return "Old BSS to New BSS Information"
	// LSA_ID
	// LSA_ID_LIST
	// LSA_INF
	// LCS_QOS
	// LSA_ACCESS_CTRL_SUPPR
	// LCS_PRIORITY
	// LOCATION_TYPE
	// LOCATION_ESTIMATE
	// POSITIONING_DATA
	// LCS_CAUSE
	// LCS_CLIENT_TYPE
	// APDU
	// NETW_ELEMENT_IDENTITY
	// GPS_ASSISTANCE_DATA
	// DECIPHERING_KEYS
	// RET_ERROR_RQST
	// RET_ERROR_CAUSE
	// SEGMENTATION
	// SERVICE_HANDOVER
	// SRC_RNC_TO_TARG_RNC_TRANSP_UMTS
	// SRC_RNC_TO_TARG_RNC_TRANSP_CDMA2000
	// RSVD_5
	// RSVD_6
	// GERAN_CLASSMARK
	// GERAN_BSC_CONTAINER
	// NEW_BSS_TO_OLD_BSS_INFO
	// GSM0800_IE_INTER_SYSTEM_INFO
	// SNA_ACCESS_INFO
	// VSTK_RAND_INFO
	// VSTK_INFO
	// PAGING_INFO
	// IMEI
	// VELOCITY_ESTIMATE
	// VGCS_FEATURE_FLAGS
	// TALKER_PRIORITY
	// EMERGENCY_SET_IND
	// TALKER_IDENTITY
	// CELL_ID_LIST_SEGMENT
	// SMS_TO_VGCS
	// VGCS_TALKER_MODE
	// VGCS_VBS_CELL_STATUS
	// CELL_ID_LIST_SEG_EST_CELLS
	// CELL_ID_LIST_SEG_CELLS_TBE
	// CELL_ID_LIST_SEG_REL_CELLS
	// CELL_ID_LIST_SEG_NE_CELLS
	// GANSS_ASSISTANCE_DATA
	// GANSS_POSITIONING_DATA
	// GANSS_LOCATION_TYPE
	// APP_DATA
	// DATA_IDENTITY
	// APP_DATA_INFO
	// MSISDN
	// AOIP_TRASP_ADDR
	// SPEECH_CODEC_LIST
	// SPEECH_CODEC
	// CALL_ID
	// CALL_ID_LIST
	// A_IF_SEL_FOR_RESET
	// KC_128
	// CSG_ID
	// REDIR_ATTEMPT_FLAG
	// REROUTE_REJ_CAUSE
	// SEND_SEQ_NUM
	// REROUTE_COMPL_OUTCOME
	// GLOBAL_CALL_REF
	// LCLS_CONFIG
	// LCLS_CONN_STATUS_CTRL
	// LCLS_CORR_NOT_NEEDED
	// LCLS_BSS_STATUS
	// LCLS_BREAK_REQ
	// CSFB_IND
	// CS_TO_PS_SRVCC
	// SRC_ENB_TO_TGT_ENB_TRANSP
	// CS_TO_PS_SRVCC_IND
	// CN_TO_MS_TRANSP_INFO
	// SELECTED_PLMN_ID
	// LAST_USED_EUTRAN_PLMN_ID
	// OLD_LAI
	// ATTACH_INDICATOR
	// SELECTED_OPERATOR
	// PS_REGISTERED_OPERATOR
	// CS_REGISTERED_OPERATOR
	default:
		return fmt.Sprintf("IE 0x%02X", int(ie))
	}
}
