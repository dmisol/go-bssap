package dtap

import(
	"errors"
	"fmt"
)

type Dtap struct {
	IEs []IE
	Raw []byte
	Header DtapHeader
}

type DtapHeader struct {
    ProtocolDisc PD_Type
    SkipInd      byte
    MsgType      Msg_Type
}

type IE struct {
	Tag DtapIE
	Value []byte
}

func DtapDecode(rawData []byte) (*Dtap, error) {
    if len(rawData) < 3 {
        return nil, errors.New("DTAP message too short: need at least 3 bytes")
    }

    dtap := &Dtap{
        Raw: rawData,
        IEs: make([]IE, 0, 10),
    }

	dtap.Header.ProtocolDisc = PD_Type(rawData[0])
    dtap.Header.SkipInd = rawData[1]
    dtap.Header.MsgType = Msg_Type(rawData[2])

    offset := 3
    for offset < len(rawData) {
        if offset >= len(rawData) {
            break
        }

		tag := DtapIE(rawData[offset])
        length := tag.format()

		switch {
			case length == -1:

			case length == -2:

			case length == 0:

			case length > 0:

			default:
            	return nil, fmt.Errorf("invalid length %d for IE 0x%02X at offset %d", length, tag, offset-1)
		}
	}

	return dtap, nil
}

func (d *Dtap) GetIE(tag DtapIE) ([]byte, bool) {
    for _, ie := range d.IEs {
        if ie.Tag == tag {
            return ie.Value, true
        }
    }
    return nil, false
}

// todo: fix
func (d *Dtap) Encode() []byte {
	return d.Raw
}
