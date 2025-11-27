package wal

const (
	// BufferSize is the default size (in bytes) of the internal write buffer.
	// Currently set to 4MB.
	BufferSize uint32 = 4 * 1024 * 1024

	// SegmentSize is the default maximum size (in bytes) of a WAL segment
	// before rotation occurs. Currently set to 64MB.
	SegmentSize uint32 = 64 * 1024 * 1024
)
