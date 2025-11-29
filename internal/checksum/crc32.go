package checksum

import "hash/crc32"

type crc struct{}

func NewCRC() *crc {
	return &crc{}
}

func (c *crc) Name() string {
	return "CRC32"
}

func (c *crc) Checksum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
