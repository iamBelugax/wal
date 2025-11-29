package segment

import (
	"context"
	"fmt"
	"os"

	"github.com/iamBelugax/wal/internal/filesys"
)

type Segment struct {
	offset int64
	id     uint64
	file   *os.File
}

func Open(id uint64, name string) (*Segment, error) {
	file, err := filesys.Open(name)
	if err != nil {
		return nil, fmt.Errorf("file open : %w", err)
	}

	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &Segment{
		id:     id,
		file:   file,
		offset: stats.Size(),
	}, nil
}

func (s *Segment) Append(ctx context.Context, record []byte) (uint64, error) {
	if _, err := s.file.Write(record); err != nil {
		return 0, err
	}

	if err := s.file.Sync(); err != nil {
		return 0, err
	}

	return 0, nil
}

func (s *Segment) ReadAt(offset int64, buffSize int) []byte {
	buf := make([]byte, buffSize)
	s.file.ReadAt(buf, offset)
	return buf
}

func (s *Segment) ID() uint64 {
	return s.id
}

func (s *Segment) Offset() uint64 {
	return uint64(s.offset)
}

func (s *Segment) Seek(offset int64, whence int) (int64, error) {
	return s.file.Seek(offset, whence)
}

func (s *Segment) Close() error {
	return s.file.Close()
}
