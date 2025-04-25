package bssmap

import "encoding/binary"

func NewCellId(lac, ci int) IE {
	i := IE([]byte{byte(CELL_ID), 5, 1, 0, 0, 0, 0})
	binary.BigEndian.PutUint16(i[3:], uint16(lac))
	binary.BigEndian.PutUint16(i[5:], uint16(ci))
	return i
}

func NewCellIdList(mcc, mnc int, lac, ci []int) IE {
	n := len(lac)
	if n != len(ci) {
		return nil
	}

	tag := byte(0)

	step := 7
	lo := 3
	co := 5

	// TODOS: pack mcc/mnc prperly
	x := []byte{0, 0xF1, 0x10}

	if mcc == 0 || mnc == 0 {
		x = x[:0]
		tag = 1
		step = 4
		lo = 0
		co = 2
	}

	b := make([]byte, 3+step*n)
	b[0] = byte(CELL_ID_LIST)
	b[1] = byte(1 + step*n)
	b[2] = tag

	for i := 0; i < n; i++ {
		for j := 0; j < len(x); j++ {
			b[3+step*i+j] = x[j]
		}
		binary.BigEndian.PutUint16(b[3+step*i+lo:], uint16(lac[i]))
		binary.BigEndian.PutUint16(b[3+step*i+co:], uint16(ci[i]))
	}
	return b
}
