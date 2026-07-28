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

func (msgType Msg_Type) GetTagsOrder() []DtapIE {
    switch msgType {
    case MSG_MM_IMSI_DETACH_IND:
        return []DtapIE{
            MS_CLASSMARK_1,
            M_IDENTITY_1,
        }

    case MSG_MM_LOC_UPD_ACCEPT:
        return []DtapIE{
            LOC_AREA_ID,
            M_IDENTITY_2,
            FLLW_ON_PROC,
            CTS_PERM,
            PLMN_LST,
            EMER_NUM_LST,
            GPRS_TIM_3,
            NON_3GPP,
        }

    case MSG_MM_LOC_UPD_REJECT:
        return []DtapIE{
            REJ_CAUSE,
            MM_TIMER,
        }

    case MSG_MM_LOC_UPD_REQUEST:
        return []DtapIE{
            LOC_UPD_TYPE,
            CIPH_KEY_SEQ_NUM,
            LOC_AREA_ID,
            MS_CLASSMARK_1,
            M_IDENTITY_1,
            MS_CLASSMARK_UMTS,
            ADD_UPD_PARAMS,
            DEVICE_PROPS,
            MS_NET_FEAT_SUP,
        }

    case MSG_MM_AUTH_REQ:
        return []DtapIE{
            CIPH_KEY_SEQ_NUM,
            SPARE_HALF_OCT,
            AUTH_PARAM_RAND,
            AUTH_PARAM_AUTN,
        }

    case MSG_MM_AUTH_RESP:
        return []DtapIE{
            AUTH_RESP_PARAM,
            AUTH_RESP_PARAM_EXT,
        }

    case MSG_MM_AUTH_REJ:
        return []DtapIE{}

    case MSG_MM_AUTH_FAIL:
        return []DtapIE{
            REJ_CAUSE,
            AUTH_FAIL_PARAM,
        }

    case MSG_MM_ID_REQ:
        return []DtapIE{
            ID_TYPE,
            SPARE_HALF_OCT,
        }

    case MSG_MM_ID_RESP:
        return []DtapIE{
            M_IDENTITY_1,
            P_TMSI_TYPE,
            ROUT_AREA_ID_2,
            P_TMSI_SIGN_2,
        }

    case MSG_MM_TMSI_REALL_CMD:
        return []DtapIE{
            LOC_AREA_ID,
            M_IDENTITY_1,
        }

    case MSG_MM_TMSI_REALL_COMPL:
        return []DtapIE{}

    case MSG_MM_CM_SERV_ACC:
        return []DtapIE{}

    case MSG_MM_CM_SERV_REJ:
        return []DtapIE{
            REJ_CAUSE,
            MM_TIMER,
        }

    case MSG_MM_CM_SERV_ABORT:
        return []DtapIE{}

    case MSG_MM_CM_SERV_REQ:
        return []DtapIE{
            CM_SERVICE_TYPE,
            CIPH_KEY_SEQ_NUM,
            MS_CLASSMARK_2,
            M_IDENTITY_1,
            PRIOR_LVL,
            ADD_UPD_PARAMS,
            DEVICE_PROPS,
        }

    case MSG_MM_CM_SERV_PROMPT:
        return []DtapIE{PD_AND_SAPI}

    case MSG_MM_CM_REEST_REQ:
        return []DtapIE{
            CIPH_KEY_SEQ_NUM,
            SPARE_HALF_OCT,
            MS_CLASSMARK_2,
            M_IDENTITY_1,
            LOC_AREA_ID,
            DEVICE_PROPS,
        }

    case MSG_MM_ABORT:
        return []DtapIE{REJ_CAUSE}

    case MSG_MM_NULL:
        return []DtapIE{}
    
    case MSG_MM_STATUS:
        return []DtapIE{REJ_CAUSE}

    case MSG_MM_INFO:
        return  []DtapIE{
            FNAME_F_NET,
            SNAME_F_NET,
            TIME_ZONE,
            TIME_ZONE_AND_TIME,
            LSA_IDEN,
            DAY_SAVING_TIME,
        }

    default:
        return []DtapIE{}
    }
}

func DtapDecode(rawData []byte) (*Dtap, error) {
    if len(rawData) < 3 {
        return nil, errors.New("DTAP message too short: need at least 3 bytes")
    }

    dtap := &Dtap{
        Raw: rawData,
        IEs: make([]IE, 0, 10),
    }

	dtap.Header.ProtocolDisc = PD_Type(rawData[0] & 0x0F)
    dtap.Header.SkipInd = rawData[0] & 0xF0
    dtap.Header.MsgType = Msg_Type(rawData[1] & 0x3F)

    offset := 2

    tagsOrder := dtap.Header.MsgType.GetTagsOrder()

    if len(tagsOrder) == 0 {
        return dtap, nil
    }

    pendingNibble := -1

    for _, expectedTag := range tagsOrder {
        if offset >= len(rawData) {
            break
        }

        ieDef := expectedTag.format()

        if ieDef.Format == FormatUnsupported {
            return nil, fmt.Errorf("unsupported IE 0x%02X at offset %d", expectedTag, offset)
        }

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
                    Value: rawData[offset + 1 : offset + 1 + ieDef.FixedLen],
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
                Value: rawData[offset : offset+length],
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
