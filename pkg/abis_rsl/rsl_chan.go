package abisrsl


type CbitsType byte

const (
    // Выделенные каналы (без L2 Pseudo Length)
    CBits_TCH_F    CbitsType = 0b00001 // Bm + ACCH's (TCH/F)
    CBits_TCH_H_0  CbitsType = 0b00010 // Lm + ACCH's (TCH/H) с T=0
    CBits_TCH_H_1  CbitsType = 0b00011 // Lm + ACCH's (TCH/H) с T=1
    CBits_SDCCH4_0 CbitsType = 0b00100 // SDCCH/4 + ACCH с T=0
    CBits_SDCCH4_1 CbitsType = 0b00101 // SDCCH/4 + ACCH с T=1
    CBits_SDCCH4_2 CbitsType = 0b00110 // SDCCH/4 + ACCH с T=2
    CBits_SDCCH4_3 CbitsType = 0b00111 // SDCCH/4 + ACCH с T=3
    CBits_SDCCH8_0 CbitsType = 0b01000 // SDCCH/8 + ACCH с T=0
    CBits_SDCCH8_1 CbitsType = 0b01001 // SDCCH/8 + ACCH с T=1
    CBits_SDCCH8_2 CbitsType = 0b01010 // SDCCH/8 + ACCH с T=2
    CBits_SDCCH8_3 CbitsType = 0b01011 // SDCCH/8 + ACCH с T=3
    CBits_SDCCH8_4 CbitsType = 0b01100 // SDCCH/8 + ACCH с T=4
    CBits_SDCCH8_5 CbitsType = 0b01101 // SDCCH/8 + ACCH с T=5
    CBits_SDCCH8_6 CbitsType = 0b01110 // SDCCH/8 + ACCH с T=6
    CBits_SDCCH8_7 CbitsType = 0b01111 // SDCCH/8 + ACCH с T=7

    // Общие каналы (с L2 Pseudo Length)
    CBits_BCCH      CbitsType = 0b10000 // BCCH
    CBits_RACH      CbitsType = 0b10001 // Uplink CCCH (RACH)
    CBits_CCCH      CbitsType = 0b10010 // Downlink CCCH (PCH + AGCH)
)

const (
    CbitsMask   byte = 0xF8
    CbitsShift  int  = 3
)

func (c CbitsType) IsDedicatedChannel() bool {
    return c >= CBits_TCH_F && c <= CBits_SDCCH8_7
}

func (c CbitsType) IsCommonChannel() bool {
    return c >= CBits_BCCH && c <= CBits_CCCH
}

func GetCbits(channelOctet byte) CbitsType {
    return CbitsType((channelOctet & CbitsMask) >> CbitsShift)
}

func (c CbitsType) HasL2PseudoLength() bool {
    return c.IsCommonChannel()
}

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
