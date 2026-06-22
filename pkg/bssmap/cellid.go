package bssmap

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
)

type MNC struct {
	Val    uint16 // 0-999
	Digits uint8  // 2 or 3 digit (MNC 01 != 001)
}

func (c *MNC) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return errors.New("xml DecodeElement: " + err.Error())
	}

	val, err := strconv.Atoi(raw)
	if err != nil {
		return errors.New("atoi: " + err.Error())
	}
	if val < 0 || val > 999 {
		return errors.New("mnc should be 0-999")
	}

	c.Val = uint16(val)
	c.Digits = uint8(len(raw))
	return nil
}

func (c MNC) String() string {
	return fmt.Sprintf("%0*d", c.Digits, c.Val)
}

type CELL_IDENT_TYPE uint8

const (
	CELL_IDENT_WHOLE_GLOBAL       CELL_IDENT_TYPE = 0
	CELL_IDENT_LAC_AND_CI         CELL_IDENT_TYPE = 1
	CELL_IDENT_CI                 CELL_IDENT_TYPE = 2
	CELL_IDENT_NO_CELL            CELL_IDENT_TYPE = 3
	CELL_IDENT_LAI                CELL_IDENT_TYPE = 4
	CELL_IDENT_LAC                CELL_IDENT_TYPE = 5
	CELL_IDENT_BSS                CELL_IDENT_TYPE = 6
	CELL_IDENT_UTRAN_PLMN_LAC_RNC CELL_IDENT_TYPE = 8
	CELL_IDENT_UTRAN_RNC          CELL_IDENT_TYPE = 9
	CELL_IDENT_UTRAN_LAC_RNC      CELL_IDENT_TYPE = 10
	CELL_IDENT_SAI                CELL_IDENT_TYPE = 11

	/* Not in 03.03 nor 08.08. Place them > 0x0f (discr_id is 4 bits) */
	CELL_IDENT_WHOLE_GLOBAL_PS CELL_IDENT_TYPE = 128
)

func (i IE) ParseCellId() (identType CELL_IDENT_TYPE, mcc uint16, mnc MNC, ci, lac uint16, err error) {
	if i.Tag() != CELL_ID {
		err = fmt.Errorf("error: wrong IE %s", i.Tag().String())
		return
	}
	if len(i) < 3 {
		err = fmt.Errorf("error: CELL_ID is too short %s", hex.EncodeToString(i))
		return
	}
	if int(i[1]) != len(i)-2 {
		err = fmt.Errorf("error: CELL_ID has invalid len %s", hex.EncodeToString(i))
		return
	}

	identType = CELL_IDENT_TYPE(i[2])
	switch identType {
	case CELL_IDENT_WHOLE_GLOBAL:
		if len(i) != 10 {
			err = fmt.Errorf("error: CELL_ID (whole CGI) unexpected len %d %s", len(i), hex.EncodeToString(i))
			return
		}
		mcc = uint16(i[3]&0x0F)*100 + uint16(i[3]>>4)*10 + uint16(i[4]&0x0F)

		mnc = MNC{Val: uint16((i[5]&0x0F)*10 + i[5]>>4), Digits: 2}
		f := i[4] >> 4
		if f != 0x0F {
			mnc.Val = mnc.Val*10 + uint16(f)
			mnc.Digits = 3
		}

		lac = binary.BigEndian.Uint16(i[6:])
		ci = binary.BigEndian.Uint16(i[8:])
	case CELL_IDENT_LAC_AND_CI:
		if len(i) != 7 {
			err = fmt.Errorf("error: CELL_ID (LAC and CI) unexpected len %d %s", len(i), hex.EncodeToString(i))
			return
		}
		lac = binary.BigEndian.Uint16(i[3:])
		ci = binary.BigEndian.Uint16(i[5:])
	case CELL_IDENT_CI:
		if len(i) != 5 {
			err = fmt.Errorf("error: CELL_ID (CI) unexpected len %d %s", len(i), hex.EncodeToString(i))
			return
		}
		ci = binary.BigEndian.Uint16(i[3:])
	default:
		err = fmt.Errorf("error: CELL_ID option %d not supported %s", int(i[2]), hex.EncodeToString(i))
	}
	return
}
