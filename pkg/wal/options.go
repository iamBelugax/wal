package wal

import (
	"strings"

	"github.com/iamBelugax/wal/internal/checksum"
	"github.com/iamBelugax/wal/internal/domain"
	"github.com/iamBelugax/wal/internal/encoding"
)

// options configures the behavior of the Write-Ahead Log.
type options struct {
	// Encoder is used to encode and decode WAL records.
	//
	// Default: Protobuf Encoder
	Encoder domain.Encoder

	Checksumer domain.Checksumer

	// DataDir is the directory where WAL segments are stored.
	//
	// Default: "/var/lib/wal"
	DataDir string

	// SegmentSize is the maximum size (in bytes) of a WAL segment
	// before a new one is created.
	//
	// Default: 64MB
	SegmentSize uint32

	// BufferSize controls the size of the internal write buffer in bytes.
	//
	// Default: 4MB
	BufferSize uint32

	// PageSize represents the I/O alignment size (in bytes) used
	// when writing WAL pages.
	//
	// Default: 4KB
	PageSize uint16
}

// Option applies a configuration change to Options.
type Option func(*options)

// WithDataDir sets the directory in which WAL segments will be stored.
func WithDataDir(dir string) Option {
	return func(o *options) {
		dir = strings.TrimSpace(dir)
		if dir != "" {
			o.DataDir = dir
		}
	}
}

// WithSegmentSize overrides the default WAL segment size.
func WithSegmentSize(size uint32) Option {
	return func(o *options) {
		if size > 0 {
			o.SegmentSize = size
		}
	}
}

// WithBufferSize overrides the default write buffer size.
func WithBufferSize(size uint32) Option {
	return func(o *options) {
		if size > 0 {
			o.BufferSize = size
		}
	}
}

// WithPageSize overrides the WAL page alignment size.
func WithPageSize(size uint16) Option {
	return func(o *options) {
		if size > 0 {
			o.PageSize = size
		}
	}
}

// DefaultOptions returns the default WAL options.
func DefaultOptions() *options {
	return &options{
		DataDir:     DataDir,
		PageSize:    PageSize,
		BufferSize:  BufferSize,
		SegmentSize: SegmentSize,
		Checksumer:  checksum.NewCRC(),
		Encoder:     encoding.NewProtobufEncoder(),
	}
}
