package dtap

import (
	"errors"
)

var (
	ErrNotImplemented = errors.New("unsupported DTAP message type")
	ErrInvalidLengthIE = errors.New("invalid DTAP message IE length")

	ErrMsgTypeNotExist = errors.New("message type does not exist for this protocol")
	ErrProtocolDiscNotExist = errors.New("protocol disc does not exist for this protocol")

	ErrUnsupportedMsgType = errors.New("message type not implemented yet")
	ErrUnsupportedProtocolDisc = errors.New("protocol discriminator not implemented yet")

	ErrInvalidFormat = errors.New("invalid format for IE")
	ErrUnexpectedEOF = errors.New("unexpected end of data")
	
	ErrInvalidLength = errors.New("Invalid Length")
	ErrL3LengthWasNotProvided = errors.New("L3 Length was not provided")
)