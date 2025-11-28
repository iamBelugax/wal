package encoding

import (
	"github.com/iamBelugax/wal/internal/domain"
	walpb "github.com/iamBelugax/wal/internal/encoding/proto/__gen__"
	"google.golang.org/protobuf/proto"
)

type protoBufEncoder struct {
	encoder *proto.MarshalOptions
	decoder *proto.UnmarshalOptions
}

// NewProtobufEncoder returns a new encoder configured for use with the WAL.
func NewProtobufEncoder() *protoBufEncoder {
	return &protoBufEncoder{
		encoder: &proto.MarshalOptions{
			Deterministic: true,
			AllowPartial:  false,
		},
		decoder: &proto.UnmarshalOptions{
			DiscardUnknown: true,
			AllowPartial:   false,
		},
	}
}

func (*protoBufEncoder) Name() string {
	return "PROTOBUF"
}

// Encode deterministically serializes a WAL record into protobuf binary form.
func (pbe *protoBufEncoder) Encode(record *domain.Record) ([]byte, error) {
	pb := &walpb.Record{}
	pb.SetChecksum(record.Checksum)
	pb.SetKind(domain.ToPBKind(record.Kind))
	pb.SetPadded(record.Padded)
	pb.SetPayload(record.Payload)
	return pbe.encoder.Marshal(pb)
}

// Decode deserializes protobuf binary data into a WAL record.
func (pbe *protoBufEncoder) Decode(encoded []byte) (*domain.Record, error) {
	pb := &walpb.Record{}
	if err := pbe.decoder.Unmarshal(encoded, pb); err != nil {
		return nil, err
	}

	return &domain.Record{
		Padded:   pb.GetPadded(),
		Payload:  pb.GetPayload(),
		Checksum: pb.GetChecksum(),
		Kind:     domain.FromPBKind(pb.GetKind()),
	}, nil
}
