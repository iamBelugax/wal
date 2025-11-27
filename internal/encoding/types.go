package encoding

import walpb "github.com/iamBelugax/wal/internal/encoding/proto/__gen__"

// RecordKind describes the logical type of a WAL record.
type RecordKind uint32

const (
	RecordKindUnspecified RecordKind = iota
	RecordKindSegmentHeader
	RecordKindData
	RecordKindRotation
	RecordKindCheckpoint
)

// FromPBKind converts a protobuf RecordKind into the local RecordKind.
func FromPBKind(kind walpb.RecordKind) RecordKind {
	switch kind {
	case walpb.RecordKind_SEGMENT_HEADER:
		return RecordKindSegmentHeader
	case walpb.RecordKind_DATA:
		return RecordKindData
	case walpb.RecordKind_ROTATION:
		return RecordKindRotation
	case walpb.RecordKind_CHECKPOINT:
		return RecordKindCheckpoint
	default:
		return RecordKindUnspecified
	}
}

// ToPBKind converts the local RecordKind into a protobuf RecordKind.
func ToPBKind(kind RecordKind) walpb.RecordKind {
	switch kind {
	case RecordKindSegmentHeader:
		return walpb.RecordKind_SEGMENT_HEADER
	case RecordKindData:
		return walpb.RecordKind_DATA
	case RecordKindRotation:
		return walpb.RecordKind_ROTATION
	case RecordKindCheckpoint:
		return walpb.RecordKind_CHECKPOINT
	default:
		return walpb.RecordKind_UNSPECIFIED
	}
}

// Record is the physical unit written to and read from the WAL.
type Record struct {
	// Logical kind of this record (see RecordKind).
	Kind RecordKind `json:"kind" msgpack:"kind" gob:"kind"`

	// Checksum for the record contents.
	Checksum uint32 `json:"checksum" msgpack:"checksum" gob:"checksum"`

	// Actual data carried by the record.
	Payload []byte `json:"payload" msgpack:"payload" gob:"payload"`

	// Optional padding bytes to align records on disk.
	Padded []byte `json:"padded,omitempty" msgpack:"padded,omitempty" gob:"padded"`
}
