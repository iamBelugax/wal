package filesys

import (
	"os"
)

// Open opens or creates a file with read, write and append access using 0644 permissions.
func Open(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
}
