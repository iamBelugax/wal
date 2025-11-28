package wal

import (
	"strings"

	"github.com/iamBelugax/wal/internal/encoding"
)

// Encoder defines the interface for WAL record codecs.
type Encoder interface {
	Name() string
	Encode(*encoding.Record) ([]byte, error)
	Decode([]byte) (*encoding.Record, error)
}

// options configures the behavior of the Write-Ahead Log.
type options struct {
	// Encoder is used to encode and decode WAL records.
	//
	// Default: Protobuf Encoder
	Encoder Encoder

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

	// SyncOnWrite forces an fsync after every write when set to true.
	//
	// Default: false
	SyncOnWrite bool
}

// Option applies a configuration change to Options.
type Option func(*options)

// WithJSONEncoder sets the WAL encoder to JSON.
func WithJSONEncoder() Option {
	return func(o *options) {
		o.Encoder = encoding.NewJSONEncoder()
	}
}

// WithGOBEncoder sets the WAL encoder to Go's gob encoding.
func WithGOBEncoder() Option {
	return func(o *options) {
		o.Encoder = encoding.NewGobEncoder()
	}
}

// WithMsgPackEncoder sets the WAL encoder to MessagePack.
func WithMsgPackEncoder() Option {
	return func(o *options) {
		o.Encoder = encoding.NewMsgPackEncoder()
	}
}

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

// WithSyncOnWrite toggles per write fsync behavior.
func WithSyncOnWrite(sync bool) Option {
	return func(o *options) {
		o.SyncOnWrite = sync
	}
}

// DefaultOptions returns the default WAL options.
func DefaultOptions() *options {
	return &options{
		SyncOnWrite: false,
		DataDir:     DataDir,
		PageSize:    PageSize,
		BufferSize:  BufferSize,
		SegmentSize: SegmentSize,
		Encoder:     encoding.NewProtobufEncoder(),
	}
}
