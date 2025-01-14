package bssmap

import "fmt"

type BSSMAP_IE byte

const (
	CIC                                 BSSMAP_IE = 1
	RSVD_0                              BSSMAP_IE = 2
	RESRC_AVAILABLE                     BSSMAP_IE = 3
	CAUSE                               BSSMAP_IE = 4
	CELL_ID                             BSSMAP_IE = 5
	PRIORITY                            BSSMAP_IE = 6
	LAYER_3_HEADER_INF                  BSSMAP_IE = 7
	IMSI                                BSSMAP_IE = 8
	TMSI                                BSSMAP_IE = 9
	ENCRYPTION_INF                      BSSMAP_IE = 10
	CHANNEL_TYPE                        BSSMAP_IE = 11
	PERIODICITY                         BSSMAP_IE = 12
	EXTENDED_RESRC_INDICATOR            BSSMAP_IE = 13
	NUMBER_OF_MSS                       BSSMAP_IE = 14
	RSVD_1                              BSSMAP_IE = 15
	RSVD_2                              BSSMAP_IE = 16
	RSVD_3                              BSSMAP_IE = 17
	CLASSMARK_INF_T2                    BSSMAP_IE = 18
	CLASSMARK_INF_T3                    BSSMAP_IE = 19
	INTERFERENCE_BAND_TO_USE            BSSMAP_IE = 20
	RR_CAUSE                            BSSMAP_IE = 21
	RSVD_4                              BSSMAP_IE = 22
	LAYER_3_INF                         BSSMAP_IE = 23
	DLCI                                BSSMAP_IE = 24
	DOWNLINK_DTX_FLAG                   BSSMAP_IE = 25
	CELL_ID_LIST                        BSSMAP_IE = 26
	RESPONSE_RQST                       BSSMAP_IE = 27
	RESRC_IND_METHOD                    BSSMAP_IE = 28
	CLASSMARK_INF_TYPE_1                BSSMAP_IE = 29
	CIC_LIST                            BSSMAP_IE = 30
	DIAGNOSTIC                          BSSMAP_IE = 31
	LAYER_3_MESSAGE_CONTENTS            BSSMAP_IE = 32
	CHOSEN_CHANNEL                      BSSMAP_IE = 33
	TOTAL_RESRC_ACCESSIBLE              BSSMAP_IE = 34
	CIPHER_RESPONSE_MODE                BSSMAP_IE = 35
	CHANNEL_NEEDED                      BSSMAP_IE = 36
	TRACE_TYPE                          BSSMAP_IE = 37
	TRIGGERID                           BSSMAP_IE = 38
	TRACE_REFERENCE                     BSSMAP_IE = 39
	TRANSACTIONID                       BSSMAP_IE = 40
	MOBILE_IDENTITY                     BSSMAP_IE = 41
	OMCID                               BSSMAP_IE = 42
	FORWARD_INDICATOR                   BSSMAP_IE = 43
	CHOSEN_ENCR_ALG                     BSSMAP_IE = 44
	CIRCUIT_POOL                        BSSMAP_IE = 45
	CIRCUIT_POOL_LIST                   BSSMAP_IE = 46
	TIME_IND                            BSSMAP_IE = 47
	RESRC_SITUATION                     BSSMAP_IE = 48
	CURRENT_CHANNEL_TYPE_1              BSSMAP_IE = 49
	QUEUEING_INDICATOR                  BSSMAP_IE = 50
	SPEECH_VERSION                      BSSMAP_IE = 64
	ASSIGNMENT_REQUIREMENT              BSSMAP_IE = 51
	TALKER_FLAG                         BSSMAP_IE = 53
	CONNECTION_RELEASE_RQSTED           BSSMAP_IE = 54
	GROUP_CALL_REFERENCE                BSSMAP_IE = 55
	EMLPP_PRIORITY                      BSSMAP_IE = 56
	CONFIG_EVO_INDI                     BSSMAP_IE = 57
	OLD_BSS_TO_NEW_BSS_INF              BSSMAP_IE = 58
	LSA_ID                              BSSMAP_IE = 59
	LSA_ID_LIST                         BSSMAP_IE = 60
	LSA_INF                             BSSMAP_IE = 61
	LCS_QOS                             BSSMAP_IE = 62
	LSA_ACCESS_CTRL_SUPPR               BSSMAP_IE = 63
	LCS_PRIORITY                        BSSMAP_IE = 67
	LOCATION_TYPE                       BSSMAP_IE = 68
	LOCATION_ESTIMATE                   BSSMAP_IE = 69
	POSITIONING_DATA                    BSSMAP_IE = 70
	LCS_CAUSE                           BSSMAP_IE = 71
	LCS_CLIENT_TYPE                     BSSMAP_IE = 72
	APDU                                BSSMAP_IE = 73
	NETW_ELEMENT_IDENTITY               BSSMAP_IE = 74
	GPS_ASSISTANCE_DATA                 BSSMAP_IE = 75
	DECIPHERING_KEYS                    BSSMAP_IE = 76
	RET_ERROR_RQST                      BSSMAP_IE = 77
	RET_ERROR_CAUSE                     BSSMAP_IE = 78
	SEGMENTATION                        BSSMAP_IE = 79
	SERVICE_HANDOVER                    BSSMAP_IE = 80
	SRC_RNC_TO_TARG_RNC_TRANSP_UMTS     BSSMAP_IE = 81
	SRC_RNC_TO_TARG_RNC_TRANSP_CDMA2000 BSSMAP_IE = 82
	RSVD_5                              BSSMAP_IE = 65
	RSVD_6                              BSSMAP_IE = 66
	GERAN_CLASSMARK                     BSSMAP_IE = 0x53
	GERAN_BSC_CONTAINER                 BSSMAP_IE = 0x54
	NEW_BSS_TO_OLD_BSS_INFO             BSSMAP_IE = 0x61
	GSM0800_IE_INTER_SYSTEM_INFO        BSSMAP_IE = 0x63
	SNA_ACCESS_INFO                     BSSMAP_IE = 0x64
	VSTK_RAND_INFO                      BSSMAP_IE = 0x65
	VSTK_INFO                           BSSMAP_IE = 0x66
	PAGING_INFO                         BSSMAP_IE = 0x67
	IMEI                                BSSMAP_IE = 0x68
	VELOCITY_ESTIMATE                   BSSMAP_IE = 0x55
	VGCS_FEATURE_FLAGS                  BSSMAP_IE = 0x69
	TALKER_PRIORITY                     BSSMAP_IE = 0x6a
	EMERGENCY_SET_IND                   BSSMAP_IE = 0x6b
	TALKER_IDENTITY                     BSSMAP_IE = 0x6c
	CELL_ID_LIST_SEGMENT                BSSMAP_IE = 0x6d
	SMS_TO_VGCS                         BSSMAP_IE = 0x6e
	VGCS_TALKER_MODE                    BSSMAP_IE = 0x6f
	VGCS_VBS_CELL_STATUS                BSSMAP_IE = 0x70
	CELL_ID_LIST_SEG_EST_CELLS          BSSMAP_IE = 0x71
	CELL_ID_LIST_SEG_CELLS_TBE          BSSMAP_IE = 0x72
	CELL_ID_LIST_SEG_REL_CELLS          BSSMAP_IE = 0x73
	CELL_ID_LIST_SEG_NE_CELLS           BSSMAP_IE = 0x74
	GANSS_ASSISTANCE_DATA               BSSMAP_IE = 0x75
	GANSS_POSITIONING_DATA              BSSMAP_IE = 0x76
	GANSS_LOCATION_TYPE                 BSSMAP_IE = 0x77
	APP_DATA                            BSSMAP_IE = 0x78
	DATA_IDENTITY                       BSSMAP_IE = 0x79
	APP_DATA_INFO                       BSSMAP_IE = 0x7a
	MSISDN                              BSSMAP_IE = 0x7b
	AOIP_TRASP_ADDR                     BSSMAP_IE = 0x7c
	SPEECH_CODEC_LIST                   BSSMAP_IE = 0x7d
	SPEECH_CODEC                        BSSMAP_IE = 0x7e
	CALL_ID                             BSSMAP_IE = 0x7f
	CALL_ID_LIST                        BSSMAP_IE = 0x80
	A_IF_SEL_FOR_RESET                  BSSMAP_IE = 0x81
	KC_128                              BSSMAP_IE = 0x83
	CSG_ID                              BSSMAP_IE = 0x84
	REDIR_ATTEMPT_FLAG                  BSSMAP_IE = 0x85
	REROUTE_REJ_CAUSE                   BSSMAP_IE = 0x86
	SEND_SEQ_NUM                        BSSMAP_IE = 0x87
	REROUTE_COMPL_OUTCOME               BSSMAP_IE = 0x88
	GLOBAL_CALL_REF                     BSSMAP_IE = 0x89
	LCLS_CONFIG                         BSSMAP_IE = 0x8a
	LCLS_CONN_STATUS_CTRL               BSSMAP_IE = 0x8b
	LCLS_CORR_NOT_NEEDED                BSSMAP_IE = 0x8c
	LCLS_BSS_STATUS                     BSSMAP_IE = 0x8d
	LCLS_BREAK_REQ                      BSSMAP_IE = 0x8e
	CSFB_IND                            BSSMAP_IE = 0x8f
	CS_TO_PS_SRVCC                      BSSMAP_IE = 0x90
	SRC_ENB_TO_TGT_ENB_TRANSP           BSSMAP_IE = 0x91
	CS_TO_PS_SRVCC_IND                  BSSMAP_IE = 0x92
	CN_TO_MS_TRANSP_INFO                BSSMAP_IE = 0x93
	SELECTED_PLMN_ID                    BSSMAP_IE = 0x94
	LAST_USED_EUTRAN_PLMN_ID            BSSMAP_IE = 0x95
	OLD_LAI                             BSSMAP_IE = 0x96
	ATTACH_INDICATOR                    BSSMAP_IE = 0x97
	SELECTED_OPERATOR                   BSSMAP_IE = 0x98
	PS_REGISTERED_OPERATOR              BSSMAP_IE = 0x99
	CS_REGISTERED_OPERATOR              BSSMAP_IE = 0x9a
)

func (ie BSSMAP_IE) Format() (isVariable bool, fixed int) {
	switch ie {
	default:
		return true, 0
	}
}

func (ie BSSMAP_IE) String() string {
	switch ie {
	case CIC:
		return "CIC"
	default:
		return fmt.Sprintf("IE 0x%02X", int(ie))
	}
}
