package dtap

//go:generate go run golang.org/x/tools/cmd/stringer -type=RR_Msg_Type -trimprefix=MSG_ -output=msg-rr_string.go

const (
	MSG_RR_INIT_REQ     RR_Msg_Type = 0x3c
	MSG_RR_ADD_ASS      RR_Msg_Type = 0x3b
	MSG_RR_IMM_ASS      RR_Msg_Type = 0x3f
	MSG_RR_IMM_ASS_EXT  RR_Msg_Type = 0x39
	MSG_RR_IMM_ASS_REJ  RR_Msg_Type = 0x3a
	MSG_RR_DTM_ASS_FAIL RR_Msg_Type = 0x48
	MSG_RR_DTM_REJECT   RR_Msg_Type = 0x49
	MSG_RR_DTM_REQUEST  RR_Msg_Type = 0x4A
	MSG_RR_PACKET_ASS   RR_Msg_Type = 0x4B

	MSG_RR_CIPH_M_CMD   RR_Msg_Type = 0x35
	MSG_RR_CIPH_M_COMPL RR_Msg_Type = 0x32

	MSG_RR_CFG_CHG_CMD RR_Msg_Type = 0x30
	MSG_RR_CFG_CHG_ACK RR_Msg_Type = 0x31
	MSG_RR_CFG_CHG_REJ RR_Msg_Type = 0x33

	MSG_RR_ASS_CMD     RR_Msg_Type = 0x2e
	MSG_RR_ASS_COMPL   RR_Msg_Type = 0x29
	MSG_RR_ASS_FAIL    RR_Msg_Type = 0x2f
	MSG_RR_HANDO_CMD   RR_Msg_Type = 0x2b
	MSG_RR_HANDO_COMPL RR_Msg_Type = 0x2c
	MSG_RR_HANDO_FAIL  RR_Msg_Type = 0x28
	MSG_RR_HANDO_INFO  RR_Msg_Type = 0x2d
	MSG_RR_DTM_ASS_CMD RR_Msg_Type = 0x4c

	MSG_RR_CELL_CHG_ORDER RR_Msg_Type = 0x08
	MSG_RR_PDCH_ASS_CMD   RR_Msg_Type = 0x23

	MSG_RR_CHAN_REL      RR_Msg_Type = 0x0d
	MSG_RR_PART_REL      RR_Msg_Type = 0x0a
	MSG_RR_PART_REL_COMP RR_Msg_Type = 0x0f

	MSG_RR_PAG_REQ_1                      RR_Msg_Type = 0x21
	MSG_RR_PAG_REQ_2                      RR_Msg_Type = 0x22
	MSG_RR_PAG_REQ_3                      RR_Msg_Type = 0x24
	MSG_RR_PAG_RESP                       RR_Msg_Type = 0x27
	MSG_RR_NOTIF_NCH                      RR_Msg_Type = 0x20
	MSG_RR_NOTIF_FACCH                    RR_Msg_Type = 0x25 /* (Reserved) */
	MSG_RR_NOTIF_RESP                     RR_Msg_Type = 0x26
	MSG_RR_PACKET_NOTIF                   RR_Msg_Type = 0x4e
	MSG_RR_UTRAN_CLSM_CHG                 RR_Msg_Type = 0x60
	MSG_RR_CDMA2K_CLSM_CHG                RR_Msg_Type = 0x62
	MSG_RR_IS_TO_UTRAN_HANDO              RR_Msg_Type = 0x63
	MSG_RR_IS_TO_CDMA2K_HANDO             RR_Msg_Type = 0x64
	MSG_RR_GERAN_IU_MODE_CLASSMARK_CHANGE RR_Msg_Type = 0x65
	MSG_RR_INTER_SYS_TO_E_UTRAN_HANDO_CMD RR_Msg_Type = 0x66

	MSG_RR_SYSINFO_8 RR_Msg_Type = 0x18
	MSG_RR_SYSINFO_1 RR_Msg_Type = 0x19
	MSG_RR_SYSINFO_2 RR_Msg_Type = 0x1a
	MSG_RR_SYSINFO_3 RR_Msg_Type = 0x1b
	MSG_RR_SYSINFO_4 RR_Msg_Type = 0x1c
	MSG_RR_SYSINFO_5 RR_Msg_Type = 0x1d
	MSG_RR_SYSINFO_6 RR_Msg_Type = 0x1e
	MSG_RR_SYSINFO_7 RR_Msg_Type = 0x1f

	MSG_RR_SYSINFO_2bis    RR_Msg_Type = 0x02
	MSG_RR_SYSINFO_2ter    RR_Msg_Type = 0x03
	MSG_RR_SYSINFO_2quater RR_Msg_Type = 0x07
	MSG_RR_SYSINFO_5bis    RR_Msg_Type = 0x05
	MSG_RR_SYSINFO_5ter    RR_Msg_Type = 0x06
	MSG_RR_SYSINFO_9       RR_Msg_Type = 0x04
	MSG_RR_SYSINFO_13      RR_Msg_Type = 0x00

	MSG_RR_SYSINFO_16 RR_Msg_Type = 0x3d
	MSG_RR_SYSINFO_17 RR_Msg_Type = 0x3e

	MSG_RR_SYSINFO_18 RR_Msg_Type = 0x40
	MSG_RR_SYSINFO_19 RR_Msg_Type = 0x41
	MSG_RR_SYSINFO_20 RR_Msg_Type = 0x42

	MSG_RR_CHAN_MODE_MODIF     RR_Msg_Type = 0x10
	MSG_RR_STATUS              RR_Msg_Type = 0x12
	MSG_RR_CHAN_MODE_MODIF_ACK RR_Msg_Type = 0x17
	MSG_RR_FREQ_REDEF          RR_Msg_Type = 0x14
	MSG_RR_MEAS_REP            RR_Msg_Type = 0x15
	MSG_RR_CLSM_CHG            RR_Msg_Type = 0x16
	MSG_RR_CLSM_ENQ            RR_Msg_Type = 0x13
	MSG_RR_EXT_MEAS_REP        RR_Msg_Type = 0x36
	MSG_RR_EXT_MEAS_REP_ORD    RR_Msg_Type = 0x37
	MSG_RR_GPRS_SUSP_REQ       RR_Msg_Type = 0x34
	MSG_RR_DTM_INFO            RR_Msg_Type = 0x4d

	MSG_RR_VGCS_UPL_GRANT RR_Msg_Type = 0x09
	MSG_RR_UPLINK_RELEASE RR_Msg_Type = 0x0e
	MSG_RR_UPLINK_FREE    RR_Msg_Type = 0x0c
	MSG_RR_UPLINK_BUSY    RR_Msg_Type = 0x2a
	MSG_RR_TALKER_IND     RR_Msg_Type = 0x11

	MSG_RR_APP_INFO RR_Msg_Type = 0x38

	/* 3GPP TS 44.018 Table 10.4.2 */ // пока нет в getRRTagsOrder
	MSG_RR_SH_SI10                    RR_Msg_Type = 0x0
	MSG_RR_SH_FACCH                   RR_Msg_Type = 0x1
	MSG_RR_SH_UL_FREE                 RR_Msg_Type = 0x2
	MSG_RR_SH_MEAS_REP                RR_Msg_Type = 0x4
	MSG_RR_SH_MEAS_INFO               RR_Msg_Type = 0x5
	MSG_RR_SH_VGCS_RECON              RR_Msg_Type = 0x6
	MSG_RR_SH_VGCS_RECON2             RR_Msg_Type = 0x7
	MSG_RR_SH_VGCS_INFO               RR_Msg_Type = 0x8
	MSG_RR_SH_VGCS_SMS                RR_Msg_Type = 0x9
	MSG_RR_SH_SI10bis                 RR_Msg_Type = 0xA
	MSG_RR_SH_SI10ter                 RR_Msg_Type = 0xB
	MSG_RR_SH_VGCS_NEIGH              RR_Msg_Type = 0xC
	MSG_RR_SH_APP_DATA                RR_Msg_Type = 0xD
)
