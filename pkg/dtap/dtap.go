package dtap

type Dtap struct {
	PD  PD_Type
	Msg DTAP_MsgType
	IEs []IE

	raw []byte
}

// todo: fix
func (d *Dtap) Encode() []byte {
	return d.raw
}

// ToDO: fix!
func DtapDecode(b []byte) (*Dtap, error) {
	_, mt := Mt(b[1])
	d := &Dtap{
		PD:  PD(b[0]),
		Msg: mt,
		IEs: make([]IE, 0),
		raw: b,
	}
	return d, nil
}
