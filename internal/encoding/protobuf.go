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
func (pbe *protoBufEncoder) Encode(v any) ([]byte, error) {
	if header, ok := v.(*domain.Header); ok && header != nil {
		pb := &walpb.RecordHeader{}
		pb.SetLsn(header.LSN)
		pb.SetMagic(header.Magic)
		pb.SetVersion(header.Version)
		pb.SetChecksum(header.Checksum)
		pb.SetTimestamp(header.Timestamp)
		pb.SetRecordSize(header.RecordSize)
		pb.SetPreviousLsn(header.PreviousLSN)
		return pbe.encoder.Marshal(pb)
	}

	if record, ok := v.(*domain.Record); ok && record != nil {
		pb := &walpb.Record{}
		pb.SetChecksum(record.Checksum)
		pb.SetKind(domain.ToPBKind(record.Kind))
		pb.SetPadded(record.Padded)
		pb.SetPayload(record.Payload)
		return pbe.encoder.Marshal(pb)
	}

	return nil, domain.ErrInvalidType
}

// Decode deserializes protobuf binary data into a WAL record.
func (pbe *protoBufEncoder) Decode(encoded []byte, v any) error {
	if header, ok := v.(*domain.Header); ok && header != nil {
		pb := &walpb.RecordHeader{}
		if err := pbe.decoder.Unmarshal(encoded, pb); err != nil {
			return err
		}

		header.LSN = pb.GetLsn()
		header.Magic = pb.GetMagic()
		header.Version = pb.GetVersion()
		header.Checksum = pb.GetChecksum()
		header.Timestamp = pb.GetTimestamp()
		header.RecordSize = pb.GetRecordSize()
		header.PreviousLSN = pb.GetPreviousLsn()
		return nil
	}

	if record, ok := v.(*domain.Record); ok && record != nil {
		pb := &walpb.Record{}
		if err := pbe.decoder.Unmarshal(encoded, pb); err != nil {
			return err
		}

		record.Padded = pb.GetPadded()
		record.Payload = pb.GetPayload()
		record.Checksum = pb.GetChecksum()
		record.Kind = domain.FromPBKind(pb.GetKind())
		return nil
	}

	return domain.ErrInvalidType
}
