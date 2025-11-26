package dtap

type Dtap struct {
	PD  PD_Type
	Msg Msg_Type
	IEs []IE

	Raw []byte
}

// todo: fix
func (d *Dtap) Encode() []byte {
	return d.Raw
}

// ToDO: fix!
func DtapDecode(b []byte) (*Dtap, error) {
	_, mt := Mt(b[1])
	d := &Dtap{
		PD:  PD(b[0]),
		Msg: mt,
		IEs: make([]IE, 0),
		Raw: b,
	}
	return d, nil
}
