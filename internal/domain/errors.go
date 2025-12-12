package domain

import (
	"errors"
	"fmt"
)

var ErrInvalidType = errors.New("invalid type")

type WalError struct {
	err      error
	msg      string
	kind     ErrorKind
	metadata map[string]any
}

type ErrorKind int

const (
	ErrKindInternal = iota + 1
	ErrKindValidation
	ErrKindRead
	ErrKindWrite
	ErrKindEncode
	ErrKindDecode
	ErrKindChecksum
	ErrKindOpenSegment
	ErrKindSeekSegment
	ErrKindCloseSegment
)

func (k ErrorKind) String() string {
	switch k {
	case ErrKindInternal:
		return "internal"
	case ErrKindValidation:
		return "validation"
	case ErrKindOpenSegment:
		return "open_segment"
	case ErrKindSeekSegment:
		return "seek"
	case ErrKindEncode:
		return "encode"
	case ErrKindDecode:
		return "decode"
	case ErrKindWrite:
		return "write"
	case ErrKindRead:
		return "read"
	case ErrKindChecksum:
		return "checksum"
	case ErrKindCloseSegment:
		return "close_segment"
	default:
		return "unknown"
	}
}

func MakeWalError(kind ErrorKind, err error, msg string) error {
	var we *WalError
	if errors.As(err, &we) && we.kind == kind {
		if we.msg == "" && msg != "" {
			return &WalError{
				err:      we.err,
				kind:     we.kind,
				msg:      msg,
				metadata: we.metadata,
			}
		}
		return err
	}

	return &WalError{
		err:  err,
		msg:  msg,
		kind: kind,
	}
}

func (e *WalError) Error() string {
	if e == nil {
		return "<nil>"
	}

	msg := e.kind.String()

	if e.msg != "" {
		msg = fmt.Sprintf("%s: %s", e.kind, e.msg)
	}
	if e.err != nil {
		msg = fmt.Sprintf("%s: %v", e.kind, e.err)
	}

	return msg
}

func (e *WalError) Unwrap() error {
	return e.err
}

func (e *WalError) IsKind(k ErrorKind) bool {
	return e != nil && e.kind == k
}

func (e *WalError) Metadata() map[string]any {
	if e == nil {
		return nil
	}
	return e.metadata
}

func (e *WalError) WithMetadata(key string, val any) *WalError {
	if e == nil {
		return nil
	}

	e.metadata[key] = val
	return e
}

func IsKind(err error, kind ErrorKind) bool {
	var e *WalError
	if errors.As(err, &e) {
		return e.kind == kind
	}
	return false
}

func IsInternalError(err error) bool {
	return IsKind(err, ErrKindInternal)
}

func IsOpenSegmentError(err error) bool {
	return IsKind(err, ErrKindOpenSegment)
}

func IsSeekError(err error) bool {
	return IsKind(err, ErrKindSeekSegment)
}

func IsEncodeError(err error) bool {
	return IsKind(err, ErrKindEncode)
}

func IsDecodeError(err error) bool {
	return IsKind(err, ErrKindDecode)
}

func IsWriteError(err error) bool {
	return IsKind(err, ErrKindWrite)
}

func IsReadError(err error) bool {
	return IsKind(err, ErrKindRead)
}

func IsChecksumError(err error) bool {
	return IsKind(err, ErrKindChecksum)
}

func IsCloseSegmentError(err error) bool {
	return IsKind(err, ErrKindCloseSegment)
}
