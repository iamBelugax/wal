package domain

// Record is the physical unit written to and read from the WAL.
type Record struct {
	// Logical kind of this record.
	Kind RecordKind `json:"kind"`

	// Checksum for the record contents.
	Checksum uint32 `json:"checksum"`

	// Actual data carried by the record.
	Payload []byte `json:"payload"`

	// Optional padding bytes to align records on disk.
	Padded []byte `json:"padded,omitempty"`
}

// Header holds basic metadata for a record entry.
type Header struct {
	// Current log sequence number.
	LSN uint64 `json:"lsn"`

	// LSN of the previous record.
	PreviousLSN uint64 `json:"prevLSN"`

	// Size of the record payload in bytes.
	RecordSize uint64 `json:"recordSize"`

	// Time the record was created.
	Timestamp uint64 `json:"timestamp"`

	// Checksum for the header contents.
	Checksum uint32 `json:"checksum"`

	// Identifier to validate record format.
	Magic uint32 `json:"magic"`

	// Record format version.
	Version uint32 `json:"version"`
}

// Entry bundles a Header and a Record into a single unit.
type Entry struct {
	Header *Header
	Record *Record
}
