package domain

type Checksumer interface {
	Name() string
	Checksum([]byte) uint32
}
