package dtap

import "fmt"

type PD_Type byte

const (
	PD_GROUP_CC PD_Type = 0x00
	PD_BCAST_CC PD_Type = 0x01
	PD_PDSS1    PD_Type = 0x02
	PD_CC       PD_Type = 0x03
	PD_PDSS2    PD_Type = 0x04
	PD_GTTP     PD_Type = 0x04
	PD_MM       PD_Type = 0x05
	PD_RR       PD_Type = 0x06
	PD_MM_GPRS  PD_Type = 0x08
	PD_SMS      PD_Type = 0x09
	PD_SM_GPRS  PD_Type = 0x0a
	PD_NC_SS    PD_Type = 0x0b
	PD_LOC      PD_Type = 0x0c
	PD_EXTEND   PD_Type = 0x0e
	PD_TEST     PD_Type = 0x0f
	PD_MASK     PD_Type = 0x0f
)

func PD(b byte) PD_Type {
	return PD_Type(b & 0x0F)
}

type DTAP_MsgType byte

func Mt(b byte) (seq int, mt DTAP_MsgType) {
	seq = int(b >> 6)
	mt = DTAP_MsgType(b & 0x3F)
	return
}

const (
	MM_Info DTAP_MsgType = 0x032
)

func (d DTAP_MsgType) String() string {
	switch d {
	case MM_Info:
		return "MM Info"
	default:
		return fmt.Sprintf("MT_%02X", int(d&0x3F))
	}
}

const (
	MSG_RR_INIT_REQ     DTAP_MsgType = 0x3c
	MSG_RR_ADD_ASS      DTAP_MsgType = 0x3b
	MSG_RR_IMM_ASS      DTAP_MsgType = 0x3f
	MSG_RR_IMM_ASS_EXT  DTAP_MsgType = 0x39
	MSG_RR_IMM_ASS_REJ  DTAP_MsgType = 0x3a
	MSG_RR_DTM_ASS_FAIL DTAP_MsgType = 0x48
	MSG_RR_DTM_REJECT   DTAP_MsgType = 0x49
	MSG_RR_DTM_REQUEST  DTAP_MsgType = 0x4A
	MSG_RR_PACKET_ASS   DTAP_MsgType = 0x4B

	MSG_RR_CIPH_M_CMD   DTAP_MsgType = 0x35
	MSG_RR_CIPH_M_COMPL DTAP_MsgType = 0x32

	MSG_RR_CFG_CHG_CMD DTAP_MsgType = 0x30
	MSG_RR_CFG_CHG_ACK DTAP_MsgType = 0x31
	MSG_RR_CFG_CHG_REJ DTAP_MsgType = 0x33

	MSG_RR_ASS_CMD     DTAP_MsgType = 0x2e
	MSG_RR_ASS_COMPL   DTAP_MsgType = 0x29
	MSG_RR_ASS_FAIL    DTAP_MsgType = 0x2f
	MSG_RR_HANDO_CMD   DTAP_MsgType = 0x2b
	MSG_RR_HANDO_COMPL DTAP_MsgType = 0x2c
	MSG_RR_HANDO_FAIL  DTAP_MsgType = 0x28
	MSG_RR_HANDO_INFO  DTAP_MsgType = 0x2d
	MSG_RR_DTM_ASS_CMD DTAP_MsgType = 0x4c

	MSG_RR_CELL_CHG_ORDER DTAP_MsgType = 0x08
	MSG_RR_PDCH_ASS_CMD   DTAP_MsgType = 0x23

	MSG_RR_CHAN_REL      DTAP_MsgType = 0x0d
	MSG_RR_PART_REL      DTAP_MsgType = 0x0a
	MSG_RR_PART_REL_COMP DTAP_MsgType = 0x0f

	MSG_RR_PAG_REQ_1          DTAP_MsgType = 0x21
	MSG_RR_PAG_REQ_2          DTAP_MsgType = 0x22
	MSG_RR_PAG_REQ_3          DTAP_MsgType = 0x24
	MSG_RR_PAG_RESP           DTAP_MsgType = 0x27
	MSG_RR_NOTIF_NCH          DTAP_MsgType = 0x20
	MSG_RR_NOTIF_FACCH        DTAP_MsgType = 0x25 /* (Reserved) */
	MSG_RR_NOTIF_RESP         DTAP_MsgType = 0x26
	MSG_RR_PACKET_NOTIF       DTAP_MsgType = 0x4e
	MSG_RR_UTRAN_CLSM_CHG     DTAP_MsgType = 0x60
	MSG_RR_CDMA2K_CLSM_CHG    DTAP_MsgType = 0x62
	MSG_RR_IS_TO_UTRAN_HANDO  DTAP_MsgType = 0x63
	MSG_RR_IS_TO_CDMA2K_HANDO DTAP_MsgType = 0x64

	MSG_RR_SYSINFO_8 DTAP_MsgType = 0x18
	MSG_RR_SYSINFO_1 DTAP_MsgType = 0x19
	MSG_RR_SYSINFO_2 DTAP_MsgType = 0x1a
	MSG_RR_SYSINFO_3 DTAP_MsgType = 0x1b
	MSG_RR_SYSINFO_4 DTAP_MsgType = 0x1c
	MSG_RR_SYSINFO_5 DTAP_MsgType = 0x1d
	MSG_RR_SYSINFO_6 DTAP_MsgType = 0x1e
	MSG_RR_SYSINFO_7 DTAP_MsgType = 0x1f

	MSG_RR_SYSINFO_2bis    DTAP_MsgType = 0x02
	MSG_RR_SYSINFO_2ter    DTAP_MsgType = 0x03
	MSG_RR_SYSINFO_2quater DTAP_MsgType = 0x07
	MSG_RR_SYSINFO_5bis    DTAP_MsgType = 0x05
	MSG_RR_SYSINFO_5ter    DTAP_MsgType = 0x06
	MSG_RR_SYSINFO_9       DTAP_MsgType = 0x04
	MSG_RR_SYSINFO_13      DTAP_MsgType = 0x00

	MSG_RR_SYSINFO_16 DTAP_MsgType = 0x3d
	MSG_RR_SYSINFO_17 DTAP_MsgType = 0x3e

	MSG_RR_SYSINFO_18 DTAP_MsgType = 0x40
	MSG_RR_SYSINFO_19 DTAP_MsgType = 0x41
	MSG_RR_SYSINFO_20 DTAP_MsgType = 0x42

	MSG_RR_CHAN_MODE_MODIF     DTAP_MsgType = 0x10
	MSG_RR_STATUS              DTAP_MsgType = 0x12
	MSG_RR_CHAN_MODE_MODIF_ACK DTAP_MsgType = 0x17
	MSG_RR_FREQ_REDEF          DTAP_MsgType = 0x14
	MSG_RR_MEAS_REP            DTAP_MsgType = 0x15
	MSG_RR_CLSM_CHG            DTAP_MsgType = 0x16
	MSG_RR_CLSM_ENQ            DTAP_MsgType = 0x13
	MSG_RR_EXT_MEAS_REP        DTAP_MsgType = 0x36
	MSG_RR_EXT_MEAS_REP_ORD    DTAP_MsgType = 0x37
	MSG_RR_GPRS_SUSP_REQ       DTAP_MsgType = 0x34
	MSG_RR_DTM_INFO            DTAP_MsgType = 0x4d

	MSG_RR_VGCS_UPL_GRANT DTAP_MsgType = 0x09
	MSG_RR_UPLINK_RELEASE DTAP_MsgType = 0x0e
	MSG_RR_UPLINK_FREE    DTAP_MsgType = 0x0c
	MSG_RR_UPLINK_BUSY    DTAP_MsgType = 0x2a
	MSG_RR_TALKER_IND     DTAP_MsgType = 0x11

	MSG_RR_APP_INFO DTAP_MsgType = 0x38

	/* 3GPP TS 44.018 Table 10.4.2 */
	MSG_RR_SH_SI10        DTAP_MsgType = 0x0
	MSG_RR_SH_FACCH       DTAP_MsgType = 0x1
	MSG_RR_SH_UL_FREE     DTAP_MsgType = 0x2
	MSG_RR_SH_MEAS_REP    DTAP_MsgType = 0x4
	MSG_RR_SH_MEAS_INFO   DTAP_MsgType = 0x5
	MSG_RR_SH_VGCS_RECON  DTAP_MsgType = 0x6
	MSG_RR_SH_VGCS_RECON2 DTAP_MsgType = 0x7
	MSG_RR_SH_VGCS_INFO   DTAP_MsgType = 0x8
	MSG_RR_SH_VGCS_SMS    DTAP_MsgType = 0x9
	MSG_RR_SH_SI10bis     DTAP_MsgType = 0xA
	MSG_RR_SH_SI10ter     DTAP_MsgType = 0xB
	MSG_RR_SH_VGCS_NEIGH  DTAP_MsgType = 0xC
	MSG_RR_SH_APP_DATA    DTAP_MsgType = 0xD

	/* Table 10.2/3GPP TS 04.08 */
	MSG_MM_IMSI_DETACH_IND DTAP_MsgType = 0x01
	MSG_MM_LOC_UPD_ACCEPT  DTAP_MsgType = 0x02
	MSG_MM_LOC_UPD_REJECT  DTAP_MsgType = 0x04
	MSG_MM_LOC_UPD_REQUEST DTAP_MsgType = 0x08

	MSG_MM_AUTH_REJ         DTAP_MsgType = 0x11
	MSG_MM_AUTH_REQ         DTAP_MsgType = 0x12
	MSG_MM_AUTH_RESP        DTAP_MsgType = 0x14
	MSG_MM_AUTH_FAIL        DTAP_MsgType = 0x1c
	MSG_MM_ID_REQ           DTAP_MsgType = 0x18
	MSG_MM_ID_RESP          DTAP_MsgType = 0x19
	MSG_MM_TMSI_REALL_CMD   DTAP_MsgType = 0x1a
	MSG_MM_TMSI_REALL_COMPL DTAP_MsgType = 0x1b

	MSG_MM_CM_SERV_ACC    DTAP_MsgType = 0x21
	MSG_MM_CM_SERV_REJ    DTAP_MsgType = 0x22
	MSG_MM_CM_SERV_ABORT  DTAP_MsgType = 0x23
	MSG_MM_CM_SERV_REQ    DTAP_MsgType = 0x24
	MSG_MM_CM_SERV_PROMPT DTAP_MsgType = 0x25
	MSG_MM_CM_REEST_REQ   DTAP_MsgType = 0x28
	MSG_MM_ABORT          DTAP_MsgType = 0x29

	MSG_MM_NULL   DTAP_MsgType = 0x30
	MSG_MM_STATUS DTAP_MsgType = 0x31
	MSG_MM_INFO   DTAP_MsgType = 0x32

	/* Table 10.3/3GPP TS 04.08 */
	MSG_CC_ALERTING    DTAP_MsgType = 0x01
	MSG_CC_CALL_CONF   DTAP_MsgType = 0x08
	MSG_CC_CALL_PROC   DTAP_MsgType = 0x02
	MSG_CC_CONNECT     DTAP_MsgType = 0x07
	MSG_CC_CONNECT_ACK DTAP_MsgType = 0x0f
	MSG_CC_EMERG_SETUP DTAP_MsgType = 0x0e
	MSG_CC_PROGRESS    DTAP_MsgType = 0x03
	MSG_CC_ESTAB       DTAP_MsgType = 0x04
	MSG_CC_ESTAB_CONF  DTAP_MsgType = 0x06
	MSG_CC_RECALL      DTAP_MsgType = 0x0b
	MSG_CC_START_CC    DTAP_MsgType = 0x09
	MSG_CC_SETUP       DTAP_MsgType = 0x05

	MSG_CC_MODIFY        DTAP_MsgType = 0x17
	MSG_CC_MODIFY_COMPL  DTAP_MsgType = 0x1f
	MSG_CC_MODIFY_REJECT DTAP_MsgType = 0x13
	MSG_CC_USER_INFO     DTAP_MsgType = 0x10
	MSG_CC_HOLD          DTAP_MsgType = 0x18
	MSG_CC_HOLD_ACK      DTAP_MsgType = 0x19
	MSG_CC_HOLD_REJ      DTAP_MsgType = 0x1a
	MSG_CC_RETR          DTAP_MsgType = 0x1c
	MSG_CC_RETR_ACK      DTAP_MsgType = 0x1d
	MSG_CC_RETR_REJ      DTAP_MsgType = 0x1e

	MSG_CC_DISCONNECT    DTAP_MsgType = 0x25
	MSG_CC_RELEASE       DTAP_MsgType = 0x2d
	MSG_CC_RELEASE_COMPL DTAP_MsgType = 0x2a

	MSG_CC_CONG_CTRL      DTAP_MsgType = 0x39
	MSG_CC_NOTIFY         DTAP_MsgType = 0x3e
	MSG_CC_STATUS         DTAP_MsgType = 0x3d
	MSG_CC_STATUS_ENQ     DTAP_MsgType = 0x34
	MSG_CC_START_DTMF     DTAP_MsgType = 0x35
	MSG_CC_STOP_DTMF      DTAP_MsgType = 0x31
	MSG_CC_STOP_DTMF_ACK  DTAP_MsgType = 0x32
	MSG_CC_START_DTMF_ACK DTAP_MsgType = 0x36
	MSG_CC_START_DTMF_REJ DTAP_MsgType = 0x37
	MSG_CC_FACILITY       DTAP_MsgType = 0x3a
)
