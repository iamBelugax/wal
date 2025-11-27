package wal

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// BufferSize is the default size (in bytes) of the internal write buffer.
	// Currently set to 4MB.
	BufferSize uint32 = 4 * 1024 * 1024

	// SegmentSize is the default maximum size (in bytes) of a WAL segment
	// before rotation occurs. Currently set to 64MB.
	SegmentSize uint32 = 64 * 1024 * 1024

	// SegmentNameFormat defines the filename format of WAL segment files.
	// "%016d.wal" produces a 16 digit, zero padded decimal ID followed by ".wal".
	//
	// Example: 12345 → "0000000000012345.wal"
	SegmentNameFormat = "%016d.wal"

	// SegmentDigits is the expected fixed length of the numeric ID inside a segment filename.
	SegmentDigits = 16

	// MaxSegmentID is the largest numeric value that can fit into a 16 digit
	// zero padded decimal segment filename.
	MaxSegmentID = 9999999999999999

	// SegmentSuffix is the fixed file extension used for WAL segment files.
	SegmentSuffix = ".wal"

	// DataDir is the directory where the WAL data segments are stored.
	DataDir = "/var/lib/wal"

	// MetaDir is the directory where the metadata for WAL files is stored.
	MetaDir = DataDir + "/meta"
)

// MakeSegmentName returns the filename for a WAL segment based on its numeric ID.
func MakeSegmentName(id uint64) (string, error) {
	if id > MaxSegmentID {
		return "", fmt.Errorf("segment id exceeds %d digit limit: %d", SegmentDigits, id)
	}
	return fmt.Sprintf(SegmentNameFormat, id), nil
}

// ExtractSegmentID parses a WAL segment filename and returns its numeric ID.
func ExtractSegmentID(name string) (uint64, error) {
	if !strings.HasSuffix(name, SegmentSuffix) {
		return 0, fmt.Errorf("invalid segment filename: missing %s suffix", SegmentSuffix)
	}

	base := strings.TrimSuffix(name, SegmentSuffix)
	if len(base) != SegmentDigits {
		return 0, fmt.Errorf("invalid segment filename: wrong ID length")
	}

	id, err := strconv.ParseUint(base, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid segment ID: %w", err)
	}

	return id, nil
}
