package dream

import "time"

type OptFunc func(*opts)

type opts struct {
	cleanUpInterval time.Duration
}

// WithCleanUp sets the key cleanup interval. Set an interval > 0 to
// automatically delete keys that exceed the interval. Use 0 to disable
// auto cleanup, keeping keys in memory indefinitely unless manually
// deleted.
func WithCleanUp(interval time.Duration) OptFunc {
	return func(o *opts) {
		o.cleanUpInterval = interval
	}
}

func defaultOpts() opts {
	return opts{
		cleanUpInterval: 0,
	}
}
