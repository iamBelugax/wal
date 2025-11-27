package encoding

import (
	"encoding/json"
)

type jsonEncoder struct{}

// NewJSONEncoder returns a new JSON based encoder.
func NewJSONEncoder() *jsonEncoder {
	return &jsonEncoder{}
}

func (*jsonEncoder) Name() string {
	return "JSON"
}

// Encode serializes a WAL record into JSON form.
func (*jsonEncoder) Encode(record *Record) ([]byte, error) {
	return json.Marshal(record)
}

// Decode deserializes JSON data into a WAL record.
func (*jsonEncoder) Decode(encoded []byte) (*Record, error) {
	record := &Record{}

	if err := json.Unmarshal(encoded, record); err != nil {
		return nil, err
	}
	return record, nil
}
