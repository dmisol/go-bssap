package bssmap

import (
	"fmt"
)

type Msg_Type byte

const (
	MSG_RESERVED_0 Msg_Type = 0

	/* ASSIGNMENT MESSAGES */
	MSG_ASSIGNMENT_RQST    Msg_Type = 1
	MSG_ASSIGNMENT_CMPL    Msg_Type = 2
	MSG_ASSIGNMENT_FAILURE Msg_Type = 3
	MSG_CHAN_MOD_RQST      Msg_Type = 8

	/*  HO MESSAGES */
	MSG_HO_RQST             Msg_Type = 16
	MSG_HO_REQUIRED         Msg_Type = 17
	MSG_HO_RQST_ACK         Msg_Type = 18
	MSG_HO_CMD              Msg_Type = 19
	MSG_HO_CMPL             Msg_Type = 20
	MSG_HO_SUCCEEDED        Msg_Type = 21
	MSG_HO_FAILURE          Msg_Type = 22
	MSG_HO_PERFORMED        Msg_Type = 23
	MSG_HO_CAND_ENQ         Msg_Type = 24
	MSG_HO_CAND_RESP        Msg_Type = 25
	MSG_HO_REQUIRED_REJ     Msg_Type = 26
	MSG_HO_DETECT           Msg_Type = 27
	MSG_INT_HO_REQUIRED     Msg_Type = 0x70
	MSG_INT_HO_REQUIRED_REJ Msg_Type = 0x71
	MSG_INT_HO_CMD          Msg_Type = 0x72
	MSG_INT_HO_ENQUIRY      Msg_Type = 0x73

	/* RELEASE MESSAGES */
	MSG_CLEAR_CMD     Msg_Type = 32
	MSG_CLEAR_CMPL    Msg_Type = 33
	MSG_CLEAR_RQST    Msg_Type = 34
	MSG_RESERVED_1    Msg_Type = 35
	MSG_RESERVED_2    Msg_Type = 36
	MSG_SAPI_N_REJECT Msg_Type = 37
	MSG_CONFUSION     Msg_Type = 38

	/* OTHER CONNECTION RELATED MESSAGES */
	MSG_SUSPEND                 Msg_Type = 40
	MSG_RESUME                  Msg_Type = 41
	MSG_CONNECTION_ORIENTED_INF Msg_Type = 42
	MSG_PERFORM_LOCATION_RQST   Msg_Type = 43
	MSG_LSA_INFORMATION         Msg_Type = 44
	MSG_PERFORM_LOCATION_RESP   Msg_Type = 45
	MSG_PERFORM_LOCATION_ABORT  Msg_Type = 46
	MSG_COMMON_ID               Msg_Type = 47
	MSG_REROUTE_CMD             Msg_Type = 0x78
	MSG_REROUTE_CMPL            Msg_Type = 0x79

	/* GENERAL MESSAGES */
	MSG_RESET              Msg_Type = 48
	MSG_RESET_ACK          Msg_Type = 49
	MSG_OVERLOAD           Msg_Type = 50
	MSG_RESERVED_3         Msg_Type = 51
	MSG_RESET_CIRCUIT      Msg_Type = 52
	MSG_RESET_CIRCUIT_ACK  Msg_Type = 53
	MSG_MSC_INVOKE_TRACE   Msg_Type = 54
	MSG_BSS_INVOKE_TRACE   Msg_Type = 55
	MSG_CONNECTIONLESS_INF Msg_Type = 58
	MSG_RESET_IP_RSRC      Msg_Type = 0x3d
	MSG_RESET_IP_RSRC_ACK  Msg_Type = 0x3e

	/* TERRESTRIAL RESOURCE MESSAGES */
	MSG_BLK                Msg_Type = 64
	MSG_BLK_ACK            Msg_Type = 65
	MSG_UNBLK              Msg_Type = 66
	MSG_UNBLK_ACK          Msg_Type = 67
	MSG_CG_BLK             Msg_Type = 68
	MSG_CG_BLK_ACK         Msg_Type = 69
	MSG_CG_UNBLK           Msg_Type = 70
	MSG_CG_UNBLK_ACK       Msg_Type = 71
	MSG_UNEQUIPPED_CIRCUIT Msg_Type = 72
	MSG_CHANGE_CIRCUIT     Msg_Type = 78
	MSG_CHANGE_CIRCUIT_ACK Msg_Type = 79

	/* RADIO RESOURCE MESSAGES */
	MSG_RESOURCE_RQST       Msg_Type = 80
	MSG_RESOURCE_INDICATION Msg_Type = 81
	MSG_PAGING              Msg_Type = 82
	MSG_CIPHER_MODE_CMD     Msg_Type = 83
	MSG_CLASSMARK_UPDATE    Msg_Type = 84
	MSG_CIPHER_MODE_CMPL    Msg_Type = 85
	MSG_QUEUING_INDICATION  Msg_Type = 86
	MSG_CMPL_LAYER_3        Msg_Type = 87
	MSG_CLASSMARK_RQST      Msg_Type = 88
	MSG_CIPHER_MODE_REJECT  Msg_Type = 89
	MSG_LOAD_INDICATION     Msg_Type = 90

	/* VGCS/VBS */
	MSG_VGCS_VBS_SETUP              Msg_Type = 4
	MSG_VGCS_VBS_SETUP_ACK          Msg_Type = 5
	MSG_VGCS_VBS_SETUP_REFUSE       Msg_Type = 6
	MSG_VGCS_VBS_ASSIGNMENT_RQST    Msg_Type = 7
	MSG_VGCS_VBS_ASSIGNMENT_RESULT  Msg_Type = 28
	MSG_VGCS_VBS_ASSIGNMENT_FAILURE Msg_Type = 29
	MSG_VGCS_VBS_QUEUING_INDICATION Msg_Type = 30
	MSG_UPLINK_RQST                 Msg_Type = 31
	MSG_UPLINK_RQST_ACK             Msg_Type = 39
	MSG_VGCS_VBS_ASSIGNMENT_STATUS  Msg_Type = 59
	MSG_VGCS_VBS_AREA_CELL_INFO     Msg_Type = 60
	MSG_UPLINK_RQST_CONFIRMATION    Msg_Type = 73
	MSG_UPLINK_RELEASE_INDICATION   Msg_Type = 74
	MSG_UPLINK_REJECT_CMD           Msg_Type = 75
	MSG_UPLINK_RELEASE_CMD          Msg_Type = 76
	MSG_UPLINK_SEIZED_CMD           Msg_Type = 77
	MSG_VGCS_ADDL_INFO              Msg_Type = 0x60
	MSG_VGCS_SMS                    Msg_Type = 0x61
	MSG_NOTIFICATION_DATA           Msg_Type = 0x62
	MSG_UPLINK_APP_DATA             Msg_Type = 0x63

	/* LOCAL SWITCHING */
	MSG_LCLS_CONNECT_CTRL     Msg_Type = 0x74
	MSG_LCLS_CONNECT_CTRL_ACK Msg_Type = 0x75
	MSG_LCLS_NOTIFICATION     Msg_Type = 0x76
)

func (m Msg_Type) String() string {
	switch m {
	case MSG_ASSIGNMENT_RQST:
		return "AssRQST"
	case MSG_ASSIGNMENT_CMPL:
		return "AssCMPL"

	case MSG_HO_RQST:
		return "HO_RQST"
	case MSG_HO_RQST_ACK:
		return "HO_RQST_ACK"
	case MSG_HO_REQUIRED:
		return "HO_REQD"
	case MSG_HO_REQUIRED_REJ:
		return "HO_REQD_REJ"
	case MSG_HO_FAILURE:
		return "HO_FAILURE"

	case MSG_RESET:
		return "RESET"
	case MSG_RESET_ACK:
		return "RESET_ACK"

	case MSG_CLEAR_RQST:
		return "CLEAR RQST"
	case MSG_CLEAR_CMD:
		return "CLEAR CMD"
	case MSG_CLEAR_CMPL:
		return "CLEAR_CMPL"

	case MSG_CMPL_LAYER_3:
		return "CMPLT_L3_INFO"

	case MSG_PAGING:
		return "PAGING"
	default:
		return fmt.Sprintf("BSSMAP_0x%02X", int(m))
	}
}
