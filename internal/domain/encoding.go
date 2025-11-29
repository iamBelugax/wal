package domain

// Encoder defines the interface for WAL record codecs.
type Encoder interface {
	Name() string
	Encode(v any) ([]byte, error)
	Decode(encoded []byte, v any) error
}
