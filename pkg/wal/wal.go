package wal

import (
	"context"

	"github.com/iamBelugax/wal/internal/domain"
	"github.com/iamBelugax/wal/internal/logger"
	"github.com/iamBelugax/wal/internal/segment"
	"go.uber.org/zap/zapcore"
)

type wal struct {
	opts          *options
	log           *logger.Logger
	activeSegment *segment.Segment
}

func Open(options ...Option) (*wal, error) {
	log, err := logger.New("wal.log", zapcore.InfoLevel)
	if err != nil {
		return nil, err
	}

	opts := DefaultOptions()
	for _, option := range options {
		option(opts)
	}

	lastSegmentId, err := LastSegmentID(opts.DataDir)
	if err != nil {
		return nil, err
	}

	var newSegmentId uint64 = 1
	if lastSegmentId != 0 {
		newSegmentId = lastSegmentId
	}

	segmentName, err := MakeSegmentName(newSegmentId)
	if err != nil {
		return nil, err
	}

	segment, err := segment.Open(newSegmentId, segmentName)
	if err != nil {
		return nil, err
	}

	wal := wal{log: log, opts: opts, activeSegment: segment}
	wal.log.Infow("segment opened successfully", "name", segmentName, "id", newSegmentId)

	defer func() {
		if err := wal.activeSegment.Close(); err != nil {
			wal.log.Errorw("failed to close active segment", "segmentId", wal.activeSegment.ID(), "error", err)
		}
	}()

	return &wal, nil
}

func (w *wal) Append(ctx context.Context, payload []byte) (uint64, error) {
	return w.activeSegment.Append(ctx)
}

func (w *wal) ReadAt(ctx context.Context, offset uint64) (*domain.Entry, error) {
	return &domain.Entry{}, nil
}

func (w *wal) ReadAll(ctx context.Context) ([]*domain.Entry, error) {
	return []*domain.Entry{}, nil
}

func (w *wal) Checkpoint(ctx context.Context) error {
	return nil
}

func (w *wal) Replay(ctx context.Context) error {
	return nil
}

func (w *wal) Close(ctx context.Context) error {
	return nil
}
