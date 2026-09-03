package dtap

const (

	/* Table 10.3/3GPP TS 04.08 */
	MSG_CC_ALERTING    CC_Msg_Type = 0x01
	MSG_CC_CALL_CONF   CC_Msg_Type = 0x08
	MSG_CC_CALL_PROC   CC_Msg_Type = 0x02
	MSG_CC_CONNECT     CC_Msg_Type = 0x07
	MSG_CC_CONNECT_ACK CC_Msg_Type = 0x0f
	MSG_CC_EMERG_SETUP CC_Msg_Type = 0x0e
	MSG_CC_PROGRESS    CC_Msg_Type = 0x03
	MSG_CC_ESTAB       CC_Msg_Type = 0x04
	MSG_CC_ESTAB_CONF  CC_Msg_Type = 0x06
	MSG_CC_RECALL      CC_Msg_Type = 0x0b
	MSG_CC_START_CC    CC_Msg_Type = 0x09
	MSG_CC_SETUP       CC_Msg_Type = 0x05

	MSG_CC_MODIFY        CC_Msg_Type = 0x17
	MSG_CC_MODIFY_COMPL  CC_Msg_Type = 0x1f
	MSG_CC_MODIFY_REJECT CC_Msg_Type = 0x13
	MSG_CC_USER_INFO     CC_Msg_Type = 0x10
	MSG_CC_HOLD          CC_Msg_Type = 0x18
	MSG_CC_HOLD_ACK      CC_Msg_Type = 0x19
	MSG_CC_HOLD_REJ      CC_Msg_Type = 0x1a
	MSG_CC_RETR          CC_Msg_Type = 0x1c
	MSG_CC_RETR_ACK      CC_Msg_Type = 0x1d
	MSG_CC_RETR_REJ      CC_Msg_Type = 0x1e

	MSG_CC_DISCONNECT    CC_Msg_Type = 0x25
	MSG_CC_RELEASE       CC_Msg_Type = 0x2d
	MSG_CC_RELEASE_COMPL CC_Msg_Type = 0x2a

	MSG_CC_CONG_CTRL      CC_Msg_Type = 0x39
	MSG_CC_NOTIFY         CC_Msg_Type = 0x3e
	MSG_CC_STATUS         CC_Msg_Type = 0x3d
	MSG_CC_STATUS_ENQ     CC_Msg_Type = 0x34
	MSG_CC_START_DTMF     CC_Msg_Type = 0x35
	MSG_CC_STOP_DTMF      CC_Msg_Type = 0x31
	MSG_CC_STOP_DTMF_ACK  CC_Msg_Type = 0x32
	MSG_CC_START_DTMF_ACK CC_Msg_Type = 0x36
	MSG_CC_START_DTMF_REJ CC_Msg_Type = 0x37
	MSG_CC_FACILITY       CC_Msg_Type = 0x3a
)
