package encoding

import (
	"github.com/iamBelugax/wal/internal/domain"
	"github.com/vmihailenco/msgpack/v5"
)

type msgPackEncoder struct{}

// NewMsgPackEncoder returns a new MessagePack based encoder.
func NewMsgPackEncoder() *msgPackEncoder {
	return &msgPackEncoder{}
}

func (*msgPackEncoder) Name() string {
	return "MSG_PACK"
}

// Encode serializes a WAL record into MessagePack format.
func (e *msgPackEncoder) Encode(record *domain.Record) ([]byte, error) {
	return msgpack.Marshal(record)
}

// Decode deserializes MessagePack encoded data into a WAL record.
func (e *msgPackEncoder) Decode(data []byte) (*domain.Record, error) {
	result := &domain.Record{}

	if err := msgpack.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return result, nil
}
