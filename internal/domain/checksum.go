package domain

// Checksumer defines the interface for checksum algorithms.
type Checksumer interface {
	Name() string
	Checksum([]byte) uint32
}
