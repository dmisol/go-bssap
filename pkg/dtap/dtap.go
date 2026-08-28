package dtap

import (
	"errors"
	"fmt"
)

type Dtap struct {
	IEs    []IE
	Header DtapHeader

	// WARNING: DO NOT use it if you can unmarshall it using dtap functionality
	Raw []byte
}

type DtapHeader struct {
	ProtocolDisc PD_Type
	HalfByte     byte
	MsgType      Msg_Type
}

type IE struct {
	Tag   DtapIE
	Value []byte
}

type Option func(*DecodeOptions)

type DecodeOptions struct {
	L3TotalLength         int
	IsL2PseudoLengthExist bool
	HalfByte              byte
}

func WithL2PseudoLength() Option {
	return func(opts *DecodeOptions) {
		opts.IsL2PseudoLengthExist = true
	}
}

func SetL2PseudoLength(exists bool) Option {
	return func(opts *DecodeOptions) {
		opts.IsL2PseudoLengthExist = exists
	}
}

func decodeHeader(rawData []byte, offset int) (DtapHeader, int, error) {
	var header DtapHeader

	if offset+2 > len(rawData) {
		return header, offset, fmt.Errorf("%w: need 2 bytes for header at offset %d",
			ErrUnexpectedEOF, offset)
	}

	header.ProtocolDisc = PD_Type(rawData[offset] & 0x0F)
	header.HalfByte = rawData[offset] & 0xF0
	offset++

	header.MsgType = Msg_Type(rawData[offset] & 0x3F)
	offset++

	return header, offset, nil
}

func decodeFormatT(data []byte, offset *int, expectedTag DtapIE, dtap *Dtap) error {

	tagFromData := data[*offset]

	if DtapIE(tagFromData) != expectedTag {
		return nil
	}

	ie := IE{
		Tag:   expectedTag,
		Value: []byte{},
	}
	dtap.IEs = append(dtap.IEs, ie)
	*offset++
	return nil
}

func decodeFormatV(data []byte, offset *int, ieDef *IEDefinition, expectedTag DtapIE,
	dtap *Dtap, pendingNibble *int, skipTags []int, idx int, options *DecodeOptions, l2PseudoLength int) error {

	if ieDef.SpecialHandling == 2 {
		return decodeFormatVVariableLength(data, offset, expectedTag, dtap, options, l2PseudoLength)
	}

	if ieDef.FixedLen == 0 {
		return decodeFormatVNibble(data, offset, ieDef, expectedTag, dtap, pendingNibble, skipTags, idx)
	}

	return decodeFormatVFixedLen(data, offset, ieDef, expectedTag, dtap)
}

func decodeFormatVVariableLength(data []byte, offset *int, expectedTag DtapIE,
	dtap *Dtap, options *DecodeOptions, l2PseudoLength int) error {

	if options == nil || options.L3TotalLength <= 0 {
		return fmt.Errorf("%w, L3TotalLength is required for variable length FormatV", ErrL3LengthWasNotProvided)
	}

	vLength := options.L3TotalLength - 1 - l2PseudoLength

	if vLength < 0 {
		return fmt.Errorf("%w, invalid negative length %d for LV at offset %d",
			ErrInvalidLength, vLength, *offset-1)
	}

	if *offset+vLength > len(data) {
		return fmt.Errorf("%w: need %d bytes for FormatV at offset %d",
			ErrUnexpectedEOF, vLength, *offset)
	}

	value := data[*offset : *offset+vLength]

	dtap.IEs = append(dtap.IEs, IE{
		Tag:   expectedTag,
		Value: value,
	})

	*offset += vLength
	return nil
}

func decodeFormatVNibble(data []byte, offset *int, ieDef *IEDefinition, expectedTag DtapIE,
	dtap *Dtap, pendingNibble *int, skipTags []int, idx int) error {

	value, err := getNextNibble(data, offset, pendingNibble)
	if err != nil {
		return err
	}

	if ieDef.SpecialHandling == 1 {
		handleSpecialSkip(value, idx, skipTags)
	}

	dtap.IEs = append(dtap.IEs, IE{
		Tag:   expectedTag,
		Value: []byte{value},
	})

	return nil
}

func decodeFormatVFixedLen(data []byte, offset *int, ieDef *IEDefinition, expectedTag DtapIE,
	dtap *Dtap) error {

	if *offset+ieDef.FixedLen > len(data) {
		return fmt.Errorf("%w: need %d bytes for FormatV at offset %d",
			ErrUnexpectedEOF, ieDef.FixedLen, *offset)
	}

	value := data[*offset : *offset+ieDef.FixedLen]

	dtap.IEs = append(dtap.IEs, IE{
		Tag:   expectedTag,
		Value: value,
	})

	*offset += ieDef.FixedLen
	return nil
}

func getNextNibble(data []byte, offset *int, pendingNibble *int) (byte, error) {

	if *pendingNibble != -1 {
		value := byte(*pendingNibble)
		*pendingNibble = -1
		return value, nil
	}

	currentByte := data[*offset]
	value := currentByte & 0x0F
	*pendingNibble = int(currentByte >> 4)
	*offset++

	return value, nil
}

func handleSpecialSkip(value byte, idx int, skipTags []int) {
	if value&0x01 == 0x01 {
		if idx+1 < len(skipTags) {
			skipTags[idx+1] = 1
		}
	} else {
		if idx+2 < len(skipTags) {
			skipTags[idx+2] = 1
		}
	}
}

func decodeFormatTLV(data []byte, offset *int, expectedTag DtapIE, dtap *Dtap) error {

	tagFromData := data[*offset]

	if DtapIE(tagFromData) != expectedTag {
		return nil
	}

	if *offset+1 >= len(data) {
		return fmt.Errorf("%w: need tag byte for FormatTLV at offset %d",
			ErrUnexpectedEOF, *offset)
	}
	*offset++

	length := int(data[*offset])

	if length < 0 {
		return fmt.Errorf("%w, invalid negative length %d for LV at offset %d",
			ErrInvalidLength, length, *offset-1)
	}

	if *offset+length+1 > len(data) {
		return fmt.Errorf("%w: need %d bytes for TLV value at offset %d",
			ErrUnexpectedEOF, length, *offset)
	}
	*offset++

	value := data[*offset : *offset+length]
	*offset += length

	dtap.IEs = append(dtap.IEs, IE{
		Tag:   DtapIE(tagFromData),
		Value: value,
	})

	return nil
}

func decodeFormatLV(data []byte, offset *int, expectedTag DtapIE, dtap *Dtap) error {

	length := int(data[*offset])
	*offset++

	if length < 0 {
		return fmt.Errorf("%w, invalid negative length %d for LV at offset %d",
			ErrInvalidLength, length, *offset-1)
	}

	if *offset+length > len(data) {
		return fmt.Errorf("%w: need %d bytes for LV value at offset %d",
			ErrUnexpectedEOF, length, *offset)
	}

	value := data[*offset : *offset+length]
	*offset += length

	dtap.IEs = append(dtap.IEs, IE{
		Tag:   expectedTag,
		Value: value,
	})

	return nil
}

func decodeFormatTV(data []byte, offset *int, ieDef *IEDefinition, expectedTag DtapIE, dtap *Dtap) error {
	if ieDef.FixedLen == -2 {
		return decodeFormatTVHalfByte(data, offset, expectedTag, dtap)
	}
	return decodeFormatTVRegular(data, offset, ieDef, expectedTag, dtap)
}

func decodeFormatTVHalfByte(data []byte, offset *int, expectedTag DtapIE, dtap *Dtap) error {

	b := data[*offset]

	tagFromData := b >> 4
	value := b & 0x0F
	if DtapIE(tagFromData) != expectedTag {
		return nil
	}

	dtap.IEs = append(dtap.IEs, IE{
		Tag:   expectedTag,
		Value: []byte{value},
	})

	*offset++
	return nil
}

func decodeFormatTVRegular(data []byte, offset *int, ieDef *IEDefinition, expectedTag DtapIE, dtap *Dtap) error {

	tagFromData := data[*offset]

	if DtapIE(tagFromData) != expectedTag {
		return nil
	}
	*offset++

	if *offset+ieDef.FixedLen-1 > len(data) {
		return fmt.Errorf("%w: need %d bytes for TV value at offset %d",
			ErrUnexpectedEOF, ieDef.FixedLen, *offset)
	}
	value := data[*offset : *offset+ieDef.FixedLen-1]
	*offset += ieDef.FixedLen - 1

	dtap.IEs = append(dtap.IEs, IE{
		Tag:   expectedTag,
		Value: value,
	})

	return nil
}

func DtapDecode(rawData []byte, opts ...Option) (*Dtap, error) {
	if len(rawData) < 2 {
		return nil, errors.New("DTAP message too short: need at least 2 bytes")
	}

	dtap := &Dtap{
		IEs: make([]IE, 0, 10),
		Raw: rawData,
	}

	var l2PseudoLength int = 0
	var offset int = 0

	options := &DecodeOptions{
		L3TotalLength:         len(rawData),
		IsL2PseudoLengthExist: false,
	}

	for _, opt := range opts {
		opt(options)
	}

	if options.IsL2PseudoLengthExist {
		l2PseudoLength = int(rawData[offset]) >> 2
		offset++
	}

	if l2PseudoLength < 0 {
		return nil, fmt.Errorf("decoding IE L2 Pseudo Length (FormatV): %w", ErrInvalidLengthIE)
	}

	header, newOffset, err := decodeHeader(rawData, offset)
	if err != nil {
		return nil, err
	}

	dtap.Header = header
	offset = newOffset

	tagsOrder, err := GetTagsOrder(dtap.Header.ProtocolDisc, dtap.Header.MsgType)
	if err != nil {
		return dtap, err
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
			if err := decodeFormatT(rawData, &offset, expectedTag, dtap); err != nil {
				return nil, fmt.Errorf("decoding IE 0x%02X (FormatT): %w", expectedTag, err)
			}
		case FormatV:
			if err := decodeFormatV(rawData, &offset, &ieDef, expectedTag,
				dtap, &pendingNibble, skipTags, i, options, l2PseudoLength); err != nil {
				return nil, fmt.Errorf("decoding IE 0x%02X (FormatV): %w", expectedTag, err)
			}
		case FormatTV:
			if err := decodeFormatTV(rawData, &offset, &ieDef, expectedTag, dtap); err != nil {
				return nil, fmt.Errorf("decoding IE 0x%02X (FormatTV): %w", expectedTag, err)
			}
		case FormatLV:
			if err := decodeFormatLV(rawData, &offset, expectedTag, dtap); err != nil {
				return nil, fmt.Errorf("decoding IE 0x%02X (FormatLV): %w", expectedTag, err)
			}
		case FormatTLV:
			if err := decodeFormatTLV(rawData, &offset, expectedTag, dtap); err != nil {
				return nil, fmt.Errorf("decoding IE 0x%02X (FormatTLV): %w", expectedTag, err)
			}
		default:
			return nil, fmt.Errorf("%w: format %v for IE 0x%02X at offset %d",
				ErrInvalidFormat, ieDef.Format, expectedTag, offset)
		}
	}

	return dtap, nil
}

func (d *Dtap) GetIEValue(tag DtapIE) ([]byte, bool) {
	for _, ie := range d.IEs {
		if ie.Tag == tag {
			return ie.Value, true
		}
	}
	return nil, false
}

func (d *Dtap) GetIEValueN(tag DtapIE, n int) ([]byte, bool) {
	count := 0
	for _, ie := range d.IEs {
		if ie.Tag == tag {
			if count == n {
				return ie.Value, true
			}
			count++
		}
	}
	return nil, false
}

func (d *Dtap) GetAllIEValues(tag DtapIE) ([][]byte, bool) {
	var result [][]byte
	for _, ie := range d.IEs {
		if ie.Tag == tag {
			val := make([]byte, len(ie.Value))
			copy(val, ie.Value)
			result = append(result, val)
		}
	}
	return result, len(result) > 0
}
