package wal

import (
	"os"

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
	// Defaults to protobuf via NewProtobufEncoder().
	Encoder Encoder

	// DataDir is the directory where WAL segments are stored.
	DataDir string

	// SegmentSize is the maximum size (in bytes) of a WAL segment
	// before a new one is created.
	SegmentSize uint32

	// BufferSize controls the size of the internal write buffer in bytes.
	BufferSize uint32

	// PageSize controls low level alignment for WAL writes.
	// Defaults to os.Getpagesize().
	PageSize uint16

	// SyncOnWrite forces an fsync after every write when set to true.
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

// DefaultOptions returns the default WAL options.
func DefaultOptions() *options {
	return &options{
		SyncOnWrite: false,
		DataDir:     DataDir,
		BufferSize:  BufferSize,
		SegmentSize: SegmentSize,
		PageSize:    uint16(os.Getpagesize()),
		Encoder:     encoding.NewProtobufEncoder(),
	}
}
