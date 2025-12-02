package wal

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

const (
	// BufferSize is the default size (4MB in bytes) of the internal write buffer.
	BufferSize = 4 * 1024 * 1024

	// SegmentSize is the default maximum size (64MB in bytes) of a WAL segment
	// before rotation occurs.
	SegmentSize = 64 * 1024 * 1024

	// PageSize represents the I/O alignment size (4KB in bytes) used
	// when writing WAL pages.
	PageSize = 4 * 1024

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

	// Version represents the version of the WAL format.
	Version = 1

	// Magic is a magic number used to identify WAL files.
	Magic = 0xDEADC0DE

	// HeaderSize is the size of the WAL header in bytes.
	HeaderSize = 51
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
		return 0, fmt.Errorf("invalid segment filename: wrong id length")
	}

	id, err := strconv.ParseUint(base, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %w", err)
	}

	return id, nil
}

// LastSegmentID returns the highest segment ID found in the directory.
func LastSegmentID(dir string) (uint64, error) {
	pattern := filepath.Join(dir, "*"+SegmentSuffix)

	fileNames, err := filepath.Glob(pattern)
	if err != nil {
		return 0, err
	}

	if len(fileNames) == 0 {
		return 0, nil
	}

	slices.Sort(fileNames)
	return ExtractSegmentID(filepath.Base(fileNames[len(fileNames)-1]))
}

// CalculatePadding returns how many padding bytes are needed to make
// HeaderSize + bytes align to pageSize. If it't already aligned, it returns 0.
func CalculatePadding(pageSize uint, bytes int) uint32 {
	totalBytes := HeaderSize + bytes
	mod := totalBytes % int(pageSize)

	if mod == 0 {
		return 0
	}
	return uint32(pageSize) - uint32(mod) - 1
}
