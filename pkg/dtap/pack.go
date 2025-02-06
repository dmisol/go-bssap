package dtap

import "fmt"

func PackCB(s string) []byte {
	b := make([]byte, 1)

	offs := 0
	tail := uint16(0)
	bits := 0
	for {
		if offs >= len(s) {
			break
		}
		c := uint16(s[offs])
		offs++

		c <<= uint16(bits)
		tail |= c
		bits += 7

		if bits < 8 && len(s) > offs {
			c = uint16(s[offs])
			offs++

			c <<= uint16(bits)
			tail |= c
			bits += 7
		}
		b = append(b, byte(tail&0xFF))
		tail >>= 8
		bits -= 8
	}
	if bits > 0 {
		b = append(b, byte(tail&0xFF))
	}
	b[0] = (8-byte(bits))&0x0F | 0x80
	fmt.Println(tail, bits, b)
	return b
}
