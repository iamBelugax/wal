package checksum

import "hash/crc32"

// crc represents a CRC32 checksum calculator.
type crc struct{}

// NewCRCChecksumer creates and returns a new CRC32 calculator.
func NewCRCChecksumer() *crc {
	return &crc{}
}

func (c *crc) Name() string {
	return "CRC32"
}

func (c *crc) Checksum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
