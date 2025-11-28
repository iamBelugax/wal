package domain

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

// Header holds basic metadata for a record entry.
type Header struct {
	// Current log sequence number.
	LSN uint64 `json:"lsn" msgpack:"lsn" gob:"lsn"`

	// LSN of the previous record.
	PreviousLSN uint64 `json:"prevLSN" msgpack:"prevLSN" gob:"prevLSN"`

	// Time the record was created.
	Timestamp uint64 `json:"timestamp" msgpack:"timestamp" gob:"timestamp"`

	// Checksum for the header contents.
	Checksum uint32 `json:"checksum" msgpack:"checksum" gob:"checksum"`

	// Identifier to validate record format.
	Magic uint32 `json:"magic" msgpack:"magic" gob:"magic"`

	// Record format version.
	Version uint32 `json:"version" msgpack:"version" gob:"version"`
}

// Entry bundles a Header and a Record into a single unit.
type Entry struct {
	Header *Header
	Record *Record
}
