package bssmap

import (
	"errors"
	"fmt"
)

// https://www.etsi.org/deliver/etsi_ts/148000_148099/148008/19.00.00_60/ts_148008v190000p.pdf
// page 149 3.2.2.11 Channel Type

//go:generate go run golang.org/x/tools/cmd/stringer -type=BSSMAP_SDI --output=channel_sdi_string.go

// Speech / data indicator
type BSSMAP_SDI uint8

const (
	SDISpeech BSSMAP_SDI = iota + 1
	SDIData
	SDISignaling
	SDISpeechCTMTelephony
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=ChType --output=channel_type_string.go

// Channel type. Only for counters and estimation what channel it is, not matches actual BSSMAP
type ChType uint8

const (
	CT_Invalid ChType = iota
	CT_TCH_F
	CT_TCH_H
	CT_TCH_F_H // TCH/F or TCH/H
	CT_SDCCH
	CT_SDCCH_TCH_F   // SDCCH or TCH/F
	CT_SDCCH_TCH_H   // SDCCH or TCH/H
	CT_SDCCH_TCH_F_H // SDCCH or TCH/F or TCH/H
)

func (i IE) ChannelType() (BSSMAP_SDI, ChType, error) {
	if i.Tag() != CHANNEL_TYPE {
		return 0, 0, errors.New("not ChannelType IE")
	}
	if len(i) < 4 {
		return 0, 0, errors.New("too short ChannelType IE")
	}
	sdi := BSSMAP_SDI(i[2] & 0x0F)
	ct := CT_Invalid
	crtByte := i[3]

	if sdi == SDISpeech || sdi == SDISpeechCTMTelephony {
		/* speech */
		switch crtByte {
		case 0x08: //"Full rate TCH channel Bm.  Prefer full rate TCH"
			ct = CT_TCH_F
		case 0x09: //"Half rate TCH channel Lm.  Prefer half rate TCH"
			ct = CT_TCH_H
		case 0x0a: //"Full or Half rate channel, Full rate preferred changes allowed after first allocation"
			ct = CT_TCH_F_H
		case 0x0b: //"Full or Half rate channel, Half rate preferred changes allowed after first allocation"
			ct = CT_TCH_F_H
		case 0x0f: //"Full or Half rate channel, changes allowed after first allocation"
			ct = CT_TCH_F_H
		case 0x1a: //"Full or Half rate channel, Full rate preferred changes between full and half rate not allowed after first allocation"
			ct = CT_TCH_F_H
		case 0x1b: //"Full or Half rate channel, Half rate preferred changes between full and half rate not allowed after first allocation"
			ct = CT_TCH_F_H
		case 0x1f: //"Full or Half rate channel, changes between full and half rate not allowed after first allocation"
			ct = CT_TCH_F_H
		}
	} else if sdi == 0x02 {
		/* data */
		switch crtByte {
		case 0x08: //"Full rate TCH channel Bm"
			ct = CT_TCH_F
		case 0x09: //"Half rate TCH channel Lm"
			ct = CT_TCH_H
		case 0x0a: //"Full or Half rate TCH channel, Full rate preferred, changes allowed also after first channel allocation as a result of the request"
			ct = CT_TCH_F_H
		case 0x0b: //"Full or Half rate TCH channel, Half rate preferred, changes allowed also after first channel allocation as a result of the request"
			ct = CT_TCH_F_H
		case 0x1a: //"Full or Half rate TCH channel, Full rate preferred, changes not allowed after first channel allocation as a result of the request"
			ct = CT_TCH_F_H
		case 0x1b: //"Full or Half rate TCH channel. Half rate preferred, changes not allowed after first channel allocation as a result of the request"
			ct = CT_TCH_F_H
		case 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27: //"Full rate TCH channels in a multislot configuration, changes by the BSS of the number of TCHs and if applicable the used radio interface rate per channel allowed after first channel allocation as a result of the request"
			ct = CT_TCH_F
		case 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37: //"Full rate TCH channels in a multislot configuration, changes by the BSS of the number of TCHs or the used radio interface rate per channel not allowed after first channel allocation as a result of the request"
			ct = CT_TCH_F
		}
	} else if sdi == 0x03 {
		/* signalling */
		switch crtByte {
		case 0x00: //"SDCCH or Full rate TCH channel Bm or Half rate TCH channel Lm"
			ct = CT_SDCCH_TCH_F_H
		case 0x01: //"SDCCH"
			ct = CT_SDCCH
		case 0x02: //"SDCCH or Full rate TCH channel Bm"
			ct = CT_SDCCH_TCH_F
		case 0x03: //"SDDCH or Half rate TCH channel Lm"
			ct = CT_SDCCH_TCH_H
		case 0x08: //"Full rate TCH channel Bm"
			ct = CT_TCH_F
		case 0x09: //"Half rate TCH channel Lm"
			ct = CT_TCH_H
		case 0x0a: //"Full or Half rate TCH channel, Full rate preferred, changes allowed also after first channel allocation as a result of the request"
			ct = CT_TCH_F_H
		case 0x0b: //"Full or Half rate TCH channel, Half rate preferred, changes allowed also after first channel allocation as a result of the request"
			ct = CT_TCH_F_H
		case 0x1a: //"Full or Half rate TCH channel, Full rate preferred, changes allowed also after first channel allocation as a result of the request"
			ct = CT_TCH_F_H
		case 0x1b: //"Full or Half rate TCH channel, Half rate preferred, changes allowed also after first channel allocation as a result of the request"
			ct = CT_TCH_F_H
		}
	} else {
		return 0, 0, errors.New("invalid SDI in ChannelType IE")
	}

	if ct == CT_Invalid {
		return 0, 0, fmt.Errorf("invalid CRT for SDI %v in ChannelType IE", sdi.String())
	}

	return sdi, ct, nil
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=BSSMAPChannelMode --output=channel_mode_string.go

// https://www.etsi.org/deliver/etsi_ts/148000_148099/148008/19.00.00_60/ts_148008v190000p.pdf
// 3.2.2.33 Chosen Channel
type BSSMAPChannelMode int

const (
	CM_NO_MODE        BSSMAPChannelMode = 0
	CM_SPEECH         BSSMAPChannelMode = 9
	CM_DATA_14_5_14_5 BSSMAPChannelMode = 14
	CM_DATA_12_0_12_0 BSSMAPChannelMode = 11
	CM_DATA_6_0_6_0   BSSMAPChannelMode = 12
	CM_DATA_3_6_3_6   BSSMAPChannelMode = 13
	CM_SIGNALLING     BSSMAPChannelMode = 8
	CM_DATA_29_0_29_0 BSSMAPChannelMode = 1
	CM_DATA_32_0_32_0 BSSMAPChannelMode = 2
	CM_DATA_43_5_43_5 BSSMAPChannelMode = 3
	CM_DATA_43_5_14_5 BSSMAPChannelMode = 4
	CM_DATA_29_0_14_5 BSSMAPChannelMode = 5
	CM_DATA_43_5_29_0 BSSMAPChannelMode = 6
	CM_DATA_14_5_43_5 BSSMAPChannelMode = 7
	CM_DATA_14_5_29_0 BSSMAPChannelMode = 10
	CM_DATA_29_0_43_5 BSSMAPChannelMode = 15
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=BSSMAPChosenChannel --output=chosen_channel_string.go

// https://www.etsi.org/deliver/etsi_ts/148000_148099/148008/19.00.00_60/ts_148008v190000p.pdf
// 3.2.2.33 Chosen Channel
type BSSMAPChosenChannel int

const (
	CHANNEL_NONE    BSSMAPChosenChannel = 0
	CHANNEL_SDCCH   BSSMAPChosenChannel = 1
	CHANNEL_TCH_F   BSSMAPChosenChannel = 8
	CHANNEL_TCH_H   BSSMAPChosenChannel = 9
	CHANNEL_2_TCH_F BSSMAPChosenChannel = 10
	CHANNEL_3_TCH_F BSSMAPChosenChannel = 11
	CHANNEL_4_TCH_F BSSMAPChosenChannel = 12
	CHANNEL_5_TCH_F BSSMAPChosenChannel = 13
	CHANNEL_6_TCH_F BSSMAPChosenChannel = 14
	CHANNEL_7_TCH_F BSSMAPChosenChannel = 15
	CHANNEL_8_TCH_F BSSMAPChosenChannel = 4
)

func (i IE) ChosenChannel() (BSSMAPChannelMode, BSSMAPChosenChannel, error) {
	if i.Tag() != CHOSEN_CHANNEL {
		return 0, 0, errors.New("not Chosen Channel IE")
	}
	if len(i) < 2 {
		return 0, 0, errors.New("too short Chosen Channel IE")
	}

	return BSSMAPChannelMode((i[1] >> 4) & 0xF), BSSMAPChosenChannel(i[1] & 0xF), nil
}
