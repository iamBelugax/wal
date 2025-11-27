package encoding

import (
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
func (pbe *protoBufEncoder) Encode(record *Record) ([]byte, error) {
	pb := &walpb.Record{}
	pb.SetChecksum(record.Checksum)
	pb.SetKind(ToPBKind(record.Kind))
	pb.SetPadded(record.Padded)
	pb.SetPayload(record.Payload)
	return pbe.encoder.Marshal(pb)
}

// Decode deserializes protobuf binary data into a WAL record.
func (pbe *protoBufEncoder) Decode(encoded []byte) (*Record, error) {
	pb := &walpb.Record{}
	if err := pbe.decoder.Unmarshal(encoded, pb); err != nil {
		return nil, err
	}

	return &Record{
		Kind:     FromPBKind(pb.GetKind()),
		Checksum: pb.GetChecksum(),
		Payload:  pb.GetPayload(),
		Padded:   pb.GetPadded(),
	}, nil
}
