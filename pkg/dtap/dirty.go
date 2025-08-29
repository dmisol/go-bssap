package dtap

func AssResult(msg Msg_Type, c CAUSE) []byte {
	b := []byte{byte(PD_RR), byte(msg), byte(c)}
	return b
}
