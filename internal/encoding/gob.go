package encoding

import (
	"bytes"
	"encoding/gob"

	"github.com/iamBelugax/wal/internal/domain"
)

type gobEncoder struct{}

// NewGobEncoder returns a new gob based encoder.
func NewGobEncoder() *gobEncoder {
	return &gobEncoder{}
}

func (*gobEncoder) Name() string {
	return "GOB"
}

// Encode serializes a WAL record into gob encoded binary form.
func (*gobEncoder) Encode(record *domain.Record) ([]byte, error) {
	buffer := bytes.Buffer{}

	if err := gob.NewEncoder(&buffer).Encode(record); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Decode deserializes gob encoded binary data into a WAL record.
func (*gobEncoder) Decode(encoded []byte) (*domain.Record, error) {
	record := &domain.Record{}
	buffer := bytes.NewBuffer(encoded)

	if err := gob.NewDecoder(buffer).Decode(record); err != nil {
		return nil, err
	}
	return record, nil
}
