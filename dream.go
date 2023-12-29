package dream

import "context"

// New instantiates a new instance of the Dream store.
func New(optfns ...OptFunc) *Store {
	o := defaultOpts()
	for _, fn := range optfns {
		fn(&o)
	}

	ctx, cancel := context.WithCancel(context.Background())
	s := &Store{
		kv:     make(map[string]data),
		cancel: cancel,
	}

	if o.cleanUpInterval > 0 {
		go s.cleanUp(ctx, o.cleanUpInterval)
	}

	return s
}
