package wal

import (
	"encoding/json"
	"fmt"
)

type wal struct {
	opts *options
}

func Open(options ...Option) (*wal, error) {
	opts := DefaultOptions()
	for _, option := range options {
		option(opts)
	}

	data, _ := json.MarshalIndent(opts, "", "  ")
	fmt.Println("Options :", string(data))

	return &wal{opts: opts}, nil
}
