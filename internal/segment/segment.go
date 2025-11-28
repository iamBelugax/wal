package segment

import (
	"context"
	"fmt"
	"os"

	"github.com/iamBelugax/wal/internal/filesys"
)

type Segment struct {
	id   uint64
	file *os.File
}

func Open(id uint64, name string) (*Segment, error) {
	file, err := filesys.Open(name)
	if err != nil {
		return nil, fmt.Errorf("file open : %w", err)
	}

	return &Segment{id: id, file: file}, nil
}

func (s *Segment) Append(ctx context.Context) (uint64, error) {
	fmt.Println("Writing to file :", s.file.Name())
	return 0, nil
}

func (s *Segment) ID() uint64 {
	return s.id
}

func (s *Segment) Close() error {
	return s.file.Close()
}
