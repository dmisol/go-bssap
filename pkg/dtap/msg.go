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

type Msg_Type byte

func Mt(b byte) (seq int, mt Msg_Type) {
	seq = int(b >> 6)
	mt = Msg_Type(b & 0x3F)
	return
}

const (
	MM_Info Msg_Type = 0x032
)

func (d Msg_Type) String() string {
	switch d {
	case MM_Info:
		return "MM Info"
	default:
		return fmt.Sprintf("MT_%02X", int(d&0x3F))
	}
}

const (
	MSG_RR_INIT_REQ     Msg_Type = 0x3c
	MSG_RR_ADD_ASS      Msg_Type = 0x3b
	MSG_RR_IMM_ASS      Msg_Type = 0x3f
	MSG_RR_IMM_ASS_EXT  Msg_Type = 0x39
	MSG_RR_IMM_ASS_REJ  Msg_Type = 0x3a
	MSG_RR_DTM_ASS_FAIL Msg_Type = 0x48
	MSG_RR_DTM_REJECT   Msg_Type = 0x49
	MSG_RR_DTM_REQUEST  Msg_Type = 0x4A
	MSG_RR_PACKET_ASS   Msg_Type = 0x4B

	MSG_RR_CIPH_M_CMD   Msg_Type = 0x35
	MSG_RR_CIPH_M_COMPL Msg_Type = 0x32

	MSG_RR_CFG_CHG_CMD Msg_Type = 0x30
	MSG_RR_CFG_CHG_ACK Msg_Type = 0x31
	MSG_RR_CFG_CHG_REJ Msg_Type = 0x33

	MSG_RR_ASS_CMD     Msg_Type = 0x2e
	MSG_RR_ASS_COMPL   Msg_Type = 0x29
	MSG_RR_ASS_FAIL    Msg_Type = 0x2f
	MSG_RR_HANDO_CMD   Msg_Type = 0x2b
	MSG_RR_HANDO_COMPL Msg_Type = 0x2c
	MSG_RR_HANDO_FAIL  Msg_Type = 0x28
	MSG_RR_HANDO_INFO  Msg_Type = 0x2d
	MSG_RR_DTM_ASS_CMD Msg_Type = 0x4c

	MSG_RR_CELL_CHG_ORDER Msg_Type = 0x08
	MSG_RR_PDCH_ASS_CMD   Msg_Type = 0x23

	MSG_RR_CHAN_REL      Msg_Type = 0x0d
	MSG_RR_PART_REL      Msg_Type = 0x0a
	MSG_RR_PART_REL_COMP Msg_Type = 0x0f

	MSG_RR_PAG_REQ_1          Msg_Type = 0x21
	MSG_RR_PAG_REQ_2          Msg_Type = 0x22
	MSG_RR_PAG_REQ_3          Msg_Type = 0x24
	MSG_RR_PAG_RESP           Msg_Type = 0x27
	MSG_RR_NOTIF_NCH          Msg_Type = 0x20
	MSG_RR_NOTIF_FACCH        Msg_Type = 0x25 /* (Reserved) */
	MSG_RR_NOTIF_RESP         Msg_Type = 0x26
	MSG_RR_PACKET_NOTIF       Msg_Type = 0x4e
	MSG_RR_UTRAN_CLSM_CHG     Msg_Type = 0x60
	MSG_RR_CDMA2K_CLSM_CHG    Msg_Type = 0x62
	MSG_RR_IS_TO_UTRAN_HANDO  Msg_Type = 0x63
	MSG_RR_IS_TO_CDMA2K_HANDO Msg_Type = 0x64

	MSG_RR_SYSINFO_8 Msg_Type = 0x18
	MSG_RR_SYSINFO_1 Msg_Type = 0x19
	MSG_RR_SYSINFO_2 Msg_Type = 0x1a
	MSG_RR_SYSINFO_3 Msg_Type = 0x1b
	MSG_RR_SYSINFO_4 Msg_Type = 0x1c
	MSG_RR_SYSINFO_5 Msg_Type = 0x1d
	MSG_RR_SYSINFO_6 Msg_Type = 0x1e
	MSG_RR_SYSINFO_7 Msg_Type = 0x1f

	MSG_RR_SYSINFO_2bis    Msg_Type = 0x02
	MSG_RR_SYSINFO_2ter    Msg_Type = 0x03
	MSG_RR_SYSINFO_2quater Msg_Type = 0x07
	MSG_RR_SYSINFO_5bis    Msg_Type = 0x05
	MSG_RR_SYSINFO_5ter    Msg_Type = 0x06
	MSG_RR_SYSINFO_9       Msg_Type = 0x04
	MSG_RR_SYSINFO_13      Msg_Type = 0x00

	MSG_RR_SYSINFO_16 Msg_Type = 0x3d
	MSG_RR_SYSINFO_17 Msg_Type = 0x3e

	MSG_RR_SYSINFO_18 Msg_Type = 0x40
	MSG_RR_SYSINFO_19 Msg_Type = 0x41
	MSG_RR_SYSINFO_20 Msg_Type = 0x42

	MSG_RR_CHAN_MODE_MODIF     Msg_Type = 0x10
	MSG_RR_STATUS              Msg_Type = 0x12
	MSG_RR_CHAN_MODE_MODIF_ACK Msg_Type = 0x17
	MSG_RR_FREQ_REDEF          Msg_Type = 0x14
	MSG_RR_MEAS_REP            Msg_Type = 0x15
	MSG_RR_CLSM_CHG            Msg_Type = 0x16
	MSG_RR_CLSM_ENQ            Msg_Type = 0x13
	MSG_RR_EXT_MEAS_REP        Msg_Type = 0x36
	MSG_RR_EXT_MEAS_REP_ORD    Msg_Type = 0x37
	MSG_RR_GPRS_SUSP_REQ       Msg_Type = 0x34
	MSG_RR_DTM_INFO            Msg_Type = 0x4d

	MSG_RR_VGCS_UPL_GRANT Msg_Type = 0x09
	MSG_RR_UPLINK_RELEASE Msg_Type = 0x0e
	MSG_RR_UPLINK_FREE    Msg_Type = 0x0c
	MSG_RR_UPLINK_BUSY    Msg_Type = 0x2a
	MSG_RR_TALKER_IND     Msg_Type = 0x11

	MSG_RR_APP_INFO Msg_Type = 0x38

	/* 3GPP TS 44.018 Table 10.4.2 */
	MSG_RR_SH_SI10        Msg_Type = 0x0
	MSG_RR_SH_FACCH       Msg_Type = 0x1
	MSG_RR_SH_UL_FREE     Msg_Type = 0x2
	MSG_RR_SH_MEAS_REP    Msg_Type = 0x4
	MSG_RR_SH_MEAS_INFO   Msg_Type = 0x5
	MSG_RR_SH_VGCS_RECON  Msg_Type = 0x6
	MSG_RR_SH_VGCS_RECON2 Msg_Type = 0x7
	MSG_RR_SH_VGCS_INFO   Msg_Type = 0x8
	MSG_RR_SH_VGCS_SMS    Msg_Type = 0x9
	MSG_RR_SH_SI10bis     Msg_Type = 0xA
	MSG_RR_SH_SI10ter     Msg_Type = 0xB
	MSG_RR_SH_VGCS_NEIGH  Msg_Type = 0xC
	MSG_RR_SH_APP_DATA    Msg_Type = 0xD

	/* Table 10.2/3GPP TS 04.08 */
	MSG_MM_IMSI_DETACH_IND Msg_Type = 0x01
	MSG_MM_LOC_UPD_ACCEPT  Msg_Type = 0x02
	MSG_MM_LOC_UPD_REJECT  Msg_Type = 0x04
	MSG_MM_LOC_UPD_REQUEST Msg_Type = 0x08

	MSG_MM_AUTH_REJ         Msg_Type = 0x11
	MSG_MM_AUTH_REQ         Msg_Type = 0x12
	MSG_MM_AUTH_RESP        Msg_Type = 0x14
	MSG_MM_AUTH_FAIL        Msg_Type = 0x1c
	MSG_MM_ID_REQ           Msg_Type = 0x18
	MSG_MM_ID_RESP          Msg_Type = 0x19
	MSG_MM_TMSI_REALL_CMD   Msg_Type = 0x1a
	MSG_MM_TMSI_REALL_COMPL Msg_Type = 0x1b

	MSG_MM_CM_SERV_ACC    Msg_Type = 0x21
	MSG_MM_CM_SERV_REJ    Msg_Type = 0x22
	MSG_MM_CM_SERV_ABORT  Msg_Type = 0x23
	MSG_MM_CM_SERV_REQ    Msg_Type = 0x24
	MSG_MM_CM_SERV_PROMPT Msg_Type = 0x25
	MSG_MM_CM_REEST_REQ   Msg_Type = 0x28
	MSG_MM_ABORT          Msg_Type = 0x29

	MSG_MM_NULL   Msg_Type = 0x30
	MSG_MM_STATUS Msg_Type = 0x31
	MSG_MM_INFO   Msg_Type = 0x32

	/* Table 10.3/3GPP TS 04.08 */
	MSG_CC_ALERTING    Msg_Type = 0x01
	MSG_CC_CALL_CONF   Msg_Type = 0x08
	MSG_CC_CALL_PROC   Msg_Type = 0x02
	MSG_CC_CONNECT     Msg_Type = 0x07
	MSG_CC_CONNECT_ACK Msg_Type = 0x0f
	MSG_CC_EMERG_SETUP Msg_Type = 0x0e
	MSG_CC_PROGRESS    Msg_Type = 0x03
	MSG_CC_ESTAB       Msg_Type = 0x04
	MSG_CC_ESTAB_CONF  Msg_Type = 0x06
	MSG_CC_RECALL      Msg_Type = 0x0b
	MSG_CC_START_CC    Msg_Type = 0x09
	MSG_CC_SETUP       Msg_Type = 0x05

	MSG_CC_MODIFY        Msg_Type = 0x17
	MSG_CC_MODIFY_COMPL  Msg_Type = 0x1f
	MSG_CC_MODIFY_REJECT Msg_Type = 0x13
	MSG_CC_USER_INFO     Msg_Type = 0x10
	MSG_CC_HOLD          Msg_Type = 0x18
	MSG_CC_HOLD_ACK      Msg_Type = 0x19
	MSG_CC_HOLD_REJ      Msg_Type = 0x1a
	MSG_CC_RETR          Msg_Type = 0x1c
	MSG_CC_RETR_ACK      Msg_Type = 0x1d
	MSG_CC_RETR_REJ      Msg_Type = 0x1e

	MSG_CC_DISCONNECT    Msg_Type = 0x25
	MSG_CC_RELEASE       Msg_Type = 0x2d
	MSG_CC_RELEASE_COMPL Msg_Type = 0x2a

	MSG_CC_CONG_CTRL      Msg_Type = 0x39
	MSG_CC_NOTIFY         Msg_Type = 0x3e
	MSG_CC_STATUS         Msg_Type = 0x3d
	MSG_CC_STATUS_ENQ     Msg_Type = 0x34
	MSG_CC_START_DTMF     Msg_Type = 0x35
	MSG_CC_STOP_DTMF      Msg_Type = 0x31
	MSG_CC_STOP_DTMF_ACK  Msg_Type = 0x32
	MSG_CC_START_DTMF_ACK Msg_Type = 0x36
	MSG_CC_START_DTMF_REJ Msg_Type = 0x37
	MSG_CC_FACILITY       Msg_Type = 0x3a
)

type CAUSE byte

// TS 144.018 10.5.2.31
const (
	CauseNORMAL CAUSE = iota
	CauseAbnormalUnspec
	CauseAbnormalChanUnaccept
	CauseAbnormalTimerExpired
	CauseAbnormalNoRadio
)
