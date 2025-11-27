package wal

type wal struct {
	opts *options
}

func Open(options ...Option) (*wal, error) {
	opts := DefaultOptions()
	for _, option := range options {
		option(opts)
	}

	return &wal{opts: opts}, nil
}
