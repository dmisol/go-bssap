package abisrsl

// from IE_CHAN_NR
func (r *RSL) DecodeChannel() (ch int, err error) {
	ie, ok := Get(r.IEs, IE_CHAN_NR)
	if !ok {
		err = ErrWrongIE
		return
	}
	if len(ie) != 2 {
		err = ErrInvalidLen
		return
	}

	ch = int(ie[1])
	return
}
