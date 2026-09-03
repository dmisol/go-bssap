package dtap

const (
	/* Table 10.2/3GPP TS 04.08 */
	MSG_MM_IMSI_DETACH_IND MM_Msg_Type = 0x01
	MSG_MM_LOC_UPD_ACCEPT  MM_Msg_Type = 0x02
	MSG_MM_LOC_UPD_REJECT  MM_Msg_Type = 0x04
	MSG_MM_LOC_UPD_REQUEST MM_Msg_Type = 0x08

	MSG_MM_AUTH_REJ         MM_Msg_Type = 0x11
	MSG_MM_AUTH_REQ         MM_Msg_Type = 0x12
	MSG_MM_AUTH_RESP        MM_Msg_Type = 0x14
	MSG_MM_AUTH_FAIL        MM_Msg_Type = 0x1c
	MSG_MM_ID_REQ           MM_Msg_Type = 0x18
	MSG_MM_ID_RESP          MM_Msg_Type = 0x19
	MSG_MM_TMSI_REALL_CMD   MM_Msg_Type = 0x1a
	MSG_MM_TMSI_REALL_COMPL MM_Msg_Type = 0x1b

	MSG_MM_CM_SERV_ACC    MM_Msg_Type = 0x21
	MSG_MM_CM_SERV_REJ    MM_Msg_Type = 0x22
	MSG_MM_CM_SERV_ABORT  MM_Msg_Type = 0x23
	MSG_MM_CM_SERV_REQ    MM_Msg_Type = 0x24
	MSG_MM_CM_SERV_PROMPT MM_Msg_Type = 0x25
	MSG_MM_CM_REEST_REQ   MM_Msg_Type = 0x28
	MSG_MM_ABORT          MM_Msg_Type = 0x29

	MSG_MM_NULL   MM_Msg_Type = 0x30
	MSG_MM_STATUS MM_Msg_Type = 0x31
	MSG_MM_INFO   MM_Msg_Type = 0x32
)
