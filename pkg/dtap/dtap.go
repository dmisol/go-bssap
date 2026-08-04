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

func DtapDecode(rawData []byte, isL2PseudoLengthExist bool) (*Dtap, error) {
    if len(rawData) < 3 {
        return nil, errors.New("DTAP message too short: need at least 3 bytes")
    }

    dtap := &Dtap{
        Raw: rawData,
        IEs: make([]IE, 0, 10),
    }

    var offset int = 0
    if(isL2PseudoLengthExist) {
        offset = 1
    }

	dtap.Header.ProtocolDisc = PD_Type(rawData[offset] & 0x0F)
    dtap.Header.SkipInd = rawData[offset] & 0xF0
    offset++

    dtap.Header.MsgType = Msg_Type(rawData[offset] & 0x3F)
    offset++
    

    tagsOrder, err := GetTagsOrder(dtap.Header.ProtocolDisc, dtap.Header.MsgType)
    if err != nil {
        return dtap, fmt.Errorf("unsupported message: PD=0x%02X, MsgType=0x%02X: %w", 
            dtap.Header.ProtocolDisc, dtap.Header.MsgType, err)
    }

    if len(tagsOrder) == 0 {
        return dtap, nil
    }

    pendingNibble := -1

    skipTags := make([]int, len(tagsOrder))

    for i, expectedTag := range tagsOrder {
        if offset >= len(rawData) {
            break
        }

        if skipTags[i] == 1 {
            continue
        }
        ieDef := format(expectedTag, dtap.Header.ProtocolDisc)
        expectedTag = ieDef.Tag
        switch ieDef.Format {
        case FormatT:
            ie := IE {
                Tag: expectedTag,
                Value: []byte{},
            }
            dtap.IEs = append(dtap.IEs, ie)
            offset += 1
        case FormatV:
            if ieDef.FixedLen == 0 {
                var value byte
                if pendingNibble != -1 {
                    value = byte(pendingNibble)
                    pendingNibble = -1
                } else {
                    currentByte := rawData[offset]
                    value = currentByte & 0x0F
                    pendingNibble = int(currentByte >> 4)
                    offset += 1
                    
                }
                if(ieDef.SpecialHandling == 1) {
                    if value & 0x01 == 0x01 {
                        skipTags[i + 1] = 1
                    } else {
                        skipTags[i + 2] = 1
                    }
                }
                ie := IE {
                    Tag: expectedTag,
                    Value: []byte{value},
                }
                dtap.IEs = append(dtap.IEs, ie)

            } else {
                ie := IE {
                    Tag: expectedTag,
                    Value: rawData[offset : offset + ieDef.FixedLen],
                }
                dtap.IEs = append(dtap.IEs, ie)
                offset += ieDef.FixedLen
            }
        case FormatTV:
            if ieDef.FixedLen == -2 {
                value := rawData[offset] & 0x0F
                tagFromData := rawData[offset] >> 4
                if DtapIE(tagFromData) != expectedTag {
                    continue
                }

                ie := IE {
                    Tag: expectedTag,
                    Value: []byte{value},
                }
                dtap.IEs = append(dtap.IEs, ie)
            } else {

                tagFromData := rawData[offset]
                if DtapIE(tagFromData) != expectedTag {
                    continue
                }

                ie := IE {
                    Tag: expectedTag,
                    Value: rawData[offset + 1 : offset + ieDef.FixedLen],
                }
                dtap.IEs = append(dtap.IEs, ie)
                offset += ieDef.FixedLen + 1
            }
        case FormatLV:
            length := int(rawData[offset])
            offset += 1
            ie := IE {
                Tag: expectedTag,
                Value: rawData[offset : offset + length],
            }
            dtap.IEs = append(dtap.IEs, ie)
            offset += length
        case FormatTLV:

            tagFromData := rawData[offset]
            if DtapIE(tagFromData) != expectedTag {
                continue
            }
            offset += 1

            length := int(rawData[offset])
            offset += 1
            ie := IE{
                Tag:   DtapIE(tagFromData),
                Value: rawData[offset : offset + length],
            }
            dtap.IEs = append(dtap.IEs, ie)
            offset += length

        default:
            return nil, fmt.Errorf("invalid format %v for IE 0x%02X at offset %d", ieDef.Format, expectedTag, offset)
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

func (d *Dtap) GetIEs(tag DtapIE) [][]byte {
    var result [][]byte
    for _, ie := range d.IEs {
        if ie.Tag == tag {
            result = append(result, ie.Value)
        }
    }
    return result
}

// todo: fix
func (d *Dtap) Encode() []byte {
	return d.Raw
}
