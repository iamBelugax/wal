package checksum

import (
	"crypto/sha256"
	"encoding/binary"
)

// sha represents a SHA-256 based checksumer.
type sha struct{}

// NewSHAChecksumer creates a new SHA-256 checksumer.
func NewSHAChecksumer() *sha {
	return &sha{}
}

func (c *sha) Name() string {
	return "SHA256"
}

func (c *sha) Checksum(data []byte) uint32 {
	var out uint32
	sum := sha256.Sum256(data)

	for i := 0; i < len(sum); i += 4 {
		out ^= binary.BigEndian.Uint32(sum[i : i+4])
	}

	return out
}
