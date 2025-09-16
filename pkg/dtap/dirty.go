package dtap

func AssResult(msg Msg_Type, c CAUSE) []byte {
	// TODO: fix second byte
	b := []byte{0x01, 0x00, 0x03, byte(PD_RR), byte(msg), byte(c)}
	return b
}
