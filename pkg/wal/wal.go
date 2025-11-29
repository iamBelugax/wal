package wal

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/iamBelugax/wal/internal/domain"
	"github.com/iamBelugax/wal/internal/logger"
	"github.com/iamBelugax/wal/internal/segment"
	"go.uber.org/zap/zapcore"
)

const (
	seekOffset = 0
)

type wal struct {
	offset        uint64
	prevLSN       uint64
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

	// defer func() {
	// 	if err := wal.activeSegment.Close(); err != nil {
	// 		wal.log.Errorw("failed to close active segment", "segmentId", wal.activeSegment.ID(), "error", err)
	// 	}
	// }()

	if wal.activeSegment.Offset() != 0 {
		if _, err = wal.activeSegment.Seek(seekOffset, io.SeekEnd); err != nil {
			return nil, err
		}
	}

	return &wal, nil
}

func (w *wal) Append(ctx context.Context, payload []byte) (uint64, error) {
	// 1. I can't use domain.Header directly in Encoder.
	// 2. Create new types for header and record.
	// 3. Change the encoder interface to use new types and add a new method to encode the header also.
	// 4. First encode the header and record with new types.
	// 5. Then generate the checksum for both and append and again encode them

	record := &domain.Record{
		Payload: payload,
		Kind:    domain.RecordKindData,
	}

	recordBytes, err := w.opts.Encoder.Encode(record)
	if err != nil {
		return 0, err
	}

	record.Checksum = w.opts.Checksumer.Checksum(recordBytes)
	recordBytes, err = w.opts.Encoder.Encode(record)
	if err != nil {
		return 0, err
	}

	entrySize := len(recordBytes)
	padding := CalculatePadding(uint(w.opts.PageSize), entrySize)
	if padding > 0 {
		record.Padded = make([]byte, padding)
		recordBytes, err = w.opts.Encoder.Encode(record)
		if err != nil {
			return 0, err
		}
		entrySize = HeaderSize + len(recordBytes)
	}

	header := &domain.Header{
		Magic:       Magic,
		Version:     Version,
		LSN:         w.offset,
		PreviousLSN: w.prevLSN,
		RecordSize:  uint64(len(recordBytes)),
		Timestamp:   uint64(time.Now().UnixNano()),
	}

	headerBytes, err := w.opts.Encoder.Encode(header)
	if err != nil {
		return 0, err
	}

	header.Checksum = w.opts.Checksumer.Checksum(headerBytes)
	headerBytes, err = w.opts.Encoder.Encode(header)
	if err != nil {
		return 0, err
	}

	w.prevLSN = w.offset
	w.offset += uint64(entrySize)

	data := headerBytes
	data = append(data, recordBytes...)

	if _, err := w.activeSegment.Append(ctx, data); err != nil {
		fmt.Println("write error", err)
	}

	return header.LSN, nil
}

func (w *wal) ReadAt(ctx context.Context, offset uint64) (*domain.Entry, error) {
	buf := w.activeSegment.ReadAt(int64(offset), 51)
	header := &domain.Header{}
	if err := w.opts.Encoder.Decode(buf, header); err != nil {
		fmt.Println("header error", err)
		return nil, err
	}

	buf = w.activeSegment.ReadAt(int64(offset)+51, int(header.RecordSize))
	record := &domain.Record{}
	if err := w.opts.Encoder.Decode(buf, record); err != nil {
		fmt.Println("record error", err)
		return nil, err
	}

	return &domain.Entry{Header: header, Record: record}, nil
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
	return w.activeSegment.Close()
}
