package dtap

func DtapDecode(b []byte) (*DtapMsg, error) {
	d := &DtapMsg{pl: b}
	return d, nil
}

type DtapMsg struct {
	// @todo
	pl []byte
}

func (d *DtapMsg) Encode() []byte {
	return d.pl
}
