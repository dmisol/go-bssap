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
	INTERFERENCE_BAND_TO_USED           BssmapIE = 20
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
	VELOCITY_ESTIMATE                   BssmapIE = 0x55
	NEW_BSS_TO_OLD_BSS_INFO             BssmapIE = 0x61
	GSM0800_IE_INTER_SYSTEM_INFO        BssmapIE = 0x63
	SNA_ACCESS_INFO                     BssmapIE = 0x64
	VSTK_RAND_INFO                      BssmapIE = 0x65
	VSTK_INFO                           BssmapIE = 0x66
	PAGING_INFO                         BssmapIE = 0x67
	IMEI                                BssmapIE = 0x68
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

// Format returns a length of the Information Element (IE)
// 1. If length > 0, it means that IE has FIXED length  (For example, "Number Of MSs" or "TMSI")
// 2. If length = 0, it means that IE has VARIABLE length
// 3. If length = -1, it means that IE is unsupported
func (ie BssmapIE) Format() (length int) {
	switch ie {
	case CIC:
		return 3
	case RESRC_AVAILABLE:
		return 21 // Need to check
	case CAUSE:
		return 3
	case CELL_ID:
		return 0
	case PRIORITY:
		return 3
	case LAYER_3_HEADER_INF:
		return 4
	case IMSI:
		return 0
	case TMSI:
		return 0
	case ENCRYPTION_INF:
		return 0
	case CHANNEL_TYPE:
		return 0
	case PERIODICITY:
		return 2
	case EXTENDED_RESRC_INDICATOR:
		return 2
	case NUMBER_OF_MSS:
		return 2
	case CLASSMARK_INF_T2:
		return 0
	case CLASSMARK_INF_T3:
		return 0
	case INTERFERENCE_BAND_TO_USED:
		return 2
	case RR_CAUSE:
		return 2
	case LAYER_3_INF:
		return 0
	case DLCI:
		return 2
	case DOWNLINK_DTX_FLAG:
		return 2
	case CELL_ID_LIST:
		return 0
	case RESPONSE_RQST:
		return 1
	case RESRC_IND_METHOD:
		return 2
	case CLASSMARK_INF_TYPE_1:
		return 2
	case CIC_LIST:
		return 0
	case DIAGNOSTIC:
		return 0
	case LAYER_3_MESSAGE_CONTENTS:
		return 0
	case CHOSEN_CHANNEL:
		return 2
	case TOTAL_RESRC_ACCESSIBLE:
		return 5
	case CIPHER_RESPONSE_MODE:
		return 2
	case CHANNEL_NEEDED:
		return 2
	case TRACE_TYPE:
		return 2
	case TRIGGERID:
		return 0
	case TRACE_REFERENCE:
		return 3
	case TRANSACTIONID:
		return 0
	case MOBILE_IDENTITY:
		return 0
	case OMCID:
		return 0
	case FORWARD_INDICATOR:
		return 2
	case CHOSEN_ENCR_ALG:
		return 2
	case CIRCUIT_POOL:
		return 2
	case CIRCUIT_POOL_LIST:
		return 0
	case TIME_IND:
		return 2
	case RESRC_SITUATION:
		return 0
	case CURRENT_CHANNEL_TYPE_1:
		return 2
	case QUEUEING_INDICATOR:
		return 2
	case SPEECH_VERSION:
		return 2
	case ASSIGNMENT_REQUIREMENT:
		return 2
	case TALKER_FLAG:
		return 1
	case CONNECTION_RELEASE_RQSTED:
		return 1
	case GROUP_CALL_REFERENCE:
		return 0
	case EMLPP_PRIORITY:
		return 2
	case CONFIG_EVO_INDI:
		return 2
	case OLD_BSS_TO_NEW_BSS_INF:
		return 0
	case LSA_ID:
		return 5
	case LSA_ID_LIST:
		return 0
	case LSA_INF:
		return 0
	case LCS_QOS:
		return 0
	case LSA_ACCESS_CTRL_SUPPR:
		return 2
	case LCS_PRIORITY:
		return 0
	case LOCATION_TYPE:
		return 0
	case LOCATION_ESTIMATE:
		return 0
	case POSITIONING_DATA:
		return 0
	case LCS_CAUSE:
		return 0
	case LCS_CLIENT_TYPE:
		return 0
	case APDU:
		return 0
	case NETW_ELEMENT_IDENTITY:
		return 0
	case GPS_ASSISTANCE_DATA:
		return 0
	case DECIPHERING_KEYS:
		return 0
	case RET_ERROR_RQST:
		return 0
	case RET_ERROR_CAUSE:
		return 0
	case SEGMENTATION:
		return 0
	case SERVICE_HANDOVER:
		return 0
	case SRC_RNC_TO_TARG_RNC_TRANSP_UMTS:
		return 0
	case SRC_RNC_TO_TARG_RNC_TRANSP_CDMA2000:
		return 0
	case GERAN_CLASSMARK:
		return 0
	case GERAN_BSC_CONTAINER:
		return 0
	case NEW_BSS_TO_OLD_BSS_INFO:
		return 0
	case GSM0800_IE_INTER_SYSTEM_INFO:
		return 0
	case SNA_ACCESS_INFO:
		return 0
	case VSTK_RAND_INFO:
		return 0
	case VSTK_INFO:
		return 0
	case PAGING_INFO:
		return 2
	case IMEI:
		return 0
	case VELOCITY_ESTIMATE:
		return 0
	case VGCS_FEATURE_FLAGS:
		return 0
	case TALKER_PRIORITY:
		return 2
	case EMERGENCY_SET_IND:
		return 1
	case TALKER_IDENTITY:
		return 0
	case CELL_ID_LIST_SEGMENT:
		return 0
	case SMS_TO_VGCS:
		return 0
	case VGCS_TALKER_MODE:
		return 0
	case VGCS_VBS_CELL_STATUS:
		return 0
	case CELL_ID_LIST_SEG_EST_CELLS:
		return 0
	case CELL_ID_LIST_SEG_CELLS_TBE:
		return 0
	case CELL_ID_LIST_SEG_REL_CELLS:
		return 0
	case CELL_ID_LIST_SEG_NE_CELLS:
		return 0
	case GANSS_ASSISTANCE_DATA:
		return 0
	case GANSS_POSITIONING_DATA:
		return 0
	case GANSS_LOCATION_TYPE:
		return 0
	case APP_DATA:
		return 0
	case DATA_IDENTITY:
		return 0
	case APP_DATA_INFO:
		return 0
	case MSISDN:
		return 0
	case AOIP_TRASP_ADDR:
		return 0
	case SPEECH_CODEC_LIST:
		return 0
	case SPEECH_CODEC:
		return 0
	case CALL_ID:
		return 5
	case CALL_ID_LIST:
		return 0
	case A_IF_SEL_FOR_RESET:
		return 2
	case KC_128:
		return 17
	case CSG_ID:
		return 0
	case REDIR_ATTEMPT_FLAG:
		return 1
	case REROUTE_REJ_CAUSE:
		return 2
	case SEND_SEQ_NUM:
		return 2
	case REROUTE_COMPL_OUTCOME:
		return 2
	case GLOBAL_CALL_REF:
		return 0
	case LCLS_CONFIG:
		return 2
	case LCLS_CONN_STATUS_CTRL:
		return 2
	case LCLS_CORR_NOT_NEEDED:
		return 1
	case LCLS_BSS_STATUS:
		return 2
	case LCLS_BREAK_REQ:
		return 1
	case CSFB_IND:
		return 1
	case CS_TO_PS_SRVCC:
		return 1
	case SRC_ENB_TO_TGT_ENB_TRANSP:
		return 0
	case CS_TO_PS_SRVCC_IND:
		return 1
	case CN_TO_MS_TRANSP_INFO:
		return 0
	case SELECTED_PLMN_ID:
		return 4
	case LAST_USED_EUTRAN_PLMN_ID:
		return 4
	case OLD_LAI:
		return 6
	case ATTACH_INDICATOR:
		return 1
	case SELECTED_OPERATOR:
		return 4
	case PS_REGISTERED_OPERATOR:
		return 4
	case CS_REGISTERED_OPERATOR:
		return 4
	// case RSVD_0: // Reserved
	// case RSVD_1: // Reserved
	// case RSVD_2: // Reserved
	// case RSVD_3: // Reserved
	// case RSVD_4: // Reserved
	// case RSVD_5: // Reserved
	// case RSVD_6: // Reserved
	default:
		return -1
	}
}

func (ie BssmapIE) String() string {
	switch ie {
	case CIC:
		return "Cic"
	case RSVD_0:
		return "Reserved"
	case RESRC_AVAILABLE:
		return "ResAvail" // Resource Available
	case CAUSE:
		return "Cause"
	case CELL_ID:
		return "CellId" // Cell Identifier
	case PRIORITY:
		return "Priority"
	case LAYER_3_HEADER_INF:
		return "L3HeaderInfo" // Layer 3 Header Information
	case IMSI:
		return "IMSI"
	case TMSI:
		return "TMSI"
	case ENCRYPTION_INF:
		return "EncInfo" // Encryption Information
	case CHANNEL_TYPE:
		return "ChanType" // Channel Type
	case PERIODICITY:
		return "Periodicity"
	case EXTENDED_RESRC_INDICATOR:
		return "ExtResInd" // Extended Resource Indicator
	case NUMBER_OF_MSS:
		return "NumOfMSs" // Number Of MSs
	case RSVD_1:
		return "Reserved"
	case RSVD_2:
		return "Reserved"
	case RSVD_3:
		return "Reserved"
	case CLASSMARK_INF_T2:
		return "ClsmarkInfoType2" // Classmark Information Type 2
	case CLASSMARK_INF_T3:
		return "ClsmarkInfoType3" // Classmark Information Type 3
	case INTERFERENCE_BAND_TO_USED:
		return "InterferBand2BeUsed" // Interference Band To Be Used
	case RR_CAUSE:
		return "RRCause"
	case RSVD_4:
		return "Reserved"
	case LAYER_3_INF:
		return "L3Info" // Layer 3 Information
	case DLCI:
		return "DLCI"
	case DOWNLINK_DTX_FLAG:
		return "DlDtxFlag" // Downlink DTX Flag
	case CELL_ID_LIST:
		return "CellIdList"
	case RESPONSE_RQST:
		return "RespRqst"
	case RESRC_IND_METHOD:
		return "ResIndMethod" // Resource Indication Method
	case CLASSMARK_INF_TYPE_1:
		return "ClsmarkInfoType1" // Classmark Information Type 1
	case CIC_LIST:
		return "CicList" // Circuit Identity Code List
	case DIAGNOSTIC:
		return "Diagnostic"
	case LAYER_3_MESSAGE_CONTENTS:
		return "L3MsgCont" // Layer 3 Message Contents
	case CHOSEN_CHANNEL:
		return "ChosenChan" // Chosen Channel
	case TOTAL_RESRC_ACCESSIBLE:
		return "TtlResAccessible" // Total Resource Accessible
	case CIPHER_RESPONSE_MODE:
		return "CipherRespMode" // Cipher Response Mode
	case CHANNEL_NEEDED:
		return "ChanNeeded" // Channel Needed
	case TRACE_TYPE:
		return "TraceType"
	case TRIGGERID:
		return "Triggerid"
	case TRACE_REFERENCE:
		return "TraceRef" // Trace Reference
	case TRANSACTIONID:
		return "TransactId" // Transactionid
	case MOBILE_IDENTITY:
		return "MobIdty" // Mobile Identity
	case OMCID:
		return "OmcId"
	case FORWARD_INDICATOR:
		return "FwdInd" // Forward Indicator
	case CHOSEN_ENCR_ALG:
		return "ChosenEncAlgo" // Chosen Encryption Algorithm
	case CIRCUIT_POOL:
		return "CircPool" // Circuit Pool
	case CIRCUIT_POOL_LIST:
		return "CircPoolList" // Circuit Pool List
	case TIME_IND:
		return "TimeInd" // Time Indication
	case RESRC_SITUATION:
		return "ResSituat" // Resource Situation
	case CURRENT_CHANNEL_TYPE_1:
		return "CurChanType1"
	case QUEUEING_INDICATOR:
		return "QueueingInd" // Queueing Indicator
	case SPEECH_VERSION:
		return "SpeechVer" // Speech Version
	case ASSIGNMENT_REQUIREMENT:
		return "AssReqrmnt" // Assignment Requirement
	case TALKER_FLAG:
		return "TalkerFlg" // Talker Flag
	case CONNECTION_RELEASE_RQSTED:
		return "ConnReleaseReq" // Connection Release Requested
	case GROUP_CALL_REFERENCE:
		return "GroupCallRef" // Group Call Reference
	case EMLPP_PRIORITY:
		return "eMlppPrior" // eMLPP Priority
	case CONFIG_EVO_INDI:
		return "CfgEvoInd" // Configuration Evolution Indication
	case OLD_BSS_TO_NEW_BSS_INF:
		return "Old2NewBssInfo"
	case LSA_ID:
		return "LsaId" // LSA Identifier
	case LSA_ID_LIST:
		return "LsaIdList" // LSA Identifier List
	case LSA_INF:
		return "LsaInfo" // LSA Information
	case LCS_QOS:
		return "LcsQoS" // LCS QoS
	case LSA_ACCESS_CTRL_SUPPR:
		return "LsaAccessCtrlSuppr" // LSA access control suppression
	case LCS_PRIORITY:
		return "LcsPrio" // LCS Priority
	case LOCATION_TYPE:
		return "LocType" // Location Type
	case LOCATION_ESTIMATE:
		return "LocEst" // Location Estimate
	case POSITIONING_DATA:
		return "PosData" // Positioning Data
	case LCS_CAUSE:
		return "LcsCause" // LCS Cause
	case LCS_CLIENT_TYPE:
		return "LcsClientType" // LCS Client Type
	case APDU:
		return "APDU"
	case NETW_ELEMENT_IDENTITY:
		return "NetwrkElemId" // Network Element Identity
	case GPS_ASSISTANCE_DATA:
		return "GpsAssistData" // GPS Assistance Data
	case DECIPHERING_KEYS:
		return "DecipheringKeys" // Deciphering Keys
	case RET_ERROR_RQST:
		return "RetErrRqst" // Return Error Request
	case RET_ERROR_CAUSE:
		return "RetErrCause" // Return Error Cause
	case SEGMENTATION:
		return "Segmentat" // Segmentation
	case SERVICE_HANDOVER:
		return "ServHO" // Service Handover
	case SRC_RNC_TO_TARG_RNC_TRANSP_UMTS:
		return "SrcRnc2TrgRncTranspInfoUmts" // Source RNC to target RNC transparent information (UMTS)
	case SRC_RNC_TO_TARG_RNC_TRANSP_CDMA2000:
		return "SrcRnc2TrgRncTranspInfoCdma2000" // Source RNC to target RNC transparent information (cdma2000)
	case RSVD_5:
		return "Reserved"
	case RSVD_6:
		return "Reserved"
	case GERAN_CLASSMARK:
		return "GeranClsmrk" // GERAN Classmark
	case GERAN_BSC_CONTAINER:
		return "GeranBscCont" // GERAN BSC Container
	case NEW_BSS_TO_OLD_BSS_INFO:
		return "NewBss2OldBssInfo" // New BSS to Old BSS Information
	case GSM0800_IE_INTER_SYSTEM_INFO:
		return "InterSysInfo" // Inter-System Information
	case SNA_ACCESS_INFO:
		return "SnaAccessInfo" // SNA Access Information
	case VSTK_RAND_INFO:
		return "VSTK_RANDInfo" // VSTK_RAND Information
	case VSTK_INFO:
		return "VSTKInfo" // VSTK Information
	case PAGING_INFO:
		return "PagingInfo" // Paging Information
	case IMEI:
		return "IMEI"
	case VELOCITY_ESTIMATE:
		return "VelocityEst" // Velocity Estimate
	case VGCS_FEATURE_FLAGS:
		return "VgcsFeatFlags" // VGCS Feature Flags
	case TALKER_PRIORITY:
		return "TalkerPrio" // Talker Priority
	case EMERGENCY_SET_IND:
		return "EmergSetInd" // Emergency Set Indication
	case TALKER_IDENTITY:
		return "TalkerId" // Talker Identity
	case CELL_ID_LIST_SEGMENT:
		return "CellIdListSegm" // Cell Identifier List Segment
	case SMS_TO_VGCS:
		return "Sms2Vgcs" // SMS to VGCS
	case VGCS_TALKER_MODE:
		return "VgcsTalkMode" // VGCS Talker Mode
	case VGCS_VBS_CELL_STATUS:
		return "Vgcs/VbsCellStatus" // VGCS/VBS Cell Status
	case CELL_ID_LIST_SEG_EST_CELLS:
		return "CellIdListSegmForEstabCells" // Cell Identifier List Segment for established cells
	case CELL_ID_LIST_SEG_CELLS_TBE:
		return "CellIdListSegmForCells2BeEstab" // Cell Identifier List Segment for cells to be established
	case CELL_ID_LIST_SEG_REL_CELLS:
		return "CellIdListSegmForRlsdCellsNoUsrPresent" // Cell Identifier List Segment for released cells - no user present
	case CELL_ID_LIST_SEG_NE_CELLS:
		return "CellIdListSegmForNotEstabCellsNoEstabPsbl" // Cell Identifier List Segment for not established cells - no establishment possible
	case GANSS_ASSISTANCE_DATA:
		return "GanssAssistData" // GANSS Assistance Data
	case GANSS_POSITIONING_DATA:
		return "GanssPosData" // GANSS Positioning Data
	case GANSS_LOCATION_TYPE:
		return "GanssLocType" // GANSS Location Type
	case APP_DATA:
		return "AppData" // Application Data
	case DATA_IDENTITY:
		return "DataId" // Data Identity
	case APP_DATA_INFO:
		return "AppDataInfo" // Application Data Information
	case MSISDN:
		return "MSISDN"
	case AOIP_TRASP_ADDR:
		return "AoIpTranspLayerAddr" // AoIP Transport Layer Address
	case SPEECH_CODEC_LIST:
		return "SpeechCodecList" // Speech Codec List
	case SPEECH_CODEC:
		return "SpeechCodec" // Speech Codec
	case CALL_ID:
		return "CallId" // Call Identifier
	case CALL_ID_LIST:
		return "CallIdList" // Call Identifier List
	case A_IF_SEL_FOR_RESET:
		return "A-IfaceSelForReset" // A-Interface Selector for RESET
	case KC_128:
		return "Kc128"
	case CSG_ID:
		return "CsgId" // CSG Identifier
	case REDIR_ATTEMPT_FLAG:
		return "RedirAttemptFlag" // Redirect Attempt Flag
	case REROUTE_REJ_CAUSE:
		return "RerouteRejectCause" // Reroute Reject Cause
	case SEND_SEQ_NUM:
		return "SendSeqNum" // Send Sequence Number
	case REROUTE_COMPL_OUTCOME:
		return "RerouteCompleteOutcome" // Reroute complete outcome
	case GLOBAL_CALL_REF:
		return "GlobalCallRef" // Global Call Reference
	case LCLS_CONFIG:
		return "LclsCfg" // LCLS-Configuration
	case LCLS_CONN_STATUS_CTRL:
		return "LclsConnStatusCtrl" // LCLS-Connection-Status-Control
	case LCLS_CORR_NOT_NEEDED:
		return "LclsCorrelNotNeed" // LCLS-Correlation-Not-Needed
	case LCLS_BSS_STATUS:
		return "LclsBssStat" // LCLS-BSS-Status
	case LCLS_BREAK_REQ:
		return "LclsBreakReq" // LCLS-Break-Request
	case CSFB_IND:
		return "CsfbInd" // CSFB Indication
	case CS_TO_PS_SRVCC:
		return "Cs2PsSrvcc" // CS to PS SRVCC
	case SRC_ENB_TO_TGT_ENB_TRANSP:
		return "SrcEnb2TrgEnbTransparInfoEutran" // Source eNB to target eNB transparent information (E-UTRAN)
	case CS_TO_PS_SRVCC_IND:
		return "Cs2PsSrvccInd" // CS to PS SRVCC Indication
	case CN_TO_MS_TRANSP_INFO:
		return "Cn2MsTransparInfo" // CN to MS transparent information
	case SELECTED_PLMN_ID:
		return "SlctdPlmnId" // Selected PLMN ID
	case LAST_USED_EUTRAN_PLMN_ID:
		return "LastUsedEutranPlmnId" // Last used E-UTRAN PLMN ID
	case OLD_LAI:
		return "OldLocAreaId" // Old Location Area Identification
	case ATTACH_INDICATOR:
		return "AttachInd" // Attach Indicator
	case SELECTED_OPERATOR:
		return "SlctdOper" // Selected Operator
	case PS_REGISTERED_OPERATOR:
		return "PsRegOper" // PS Registered Operator
	case CS_REGISTERED_OPERATOR:
		return "CsRegOper" // CS Registered Operator
	default:
		return fmt.Sprintf("IE_0x%02x", int(ie))
	}
}
