package retrystrategies

import (
	"time"

	"github.com/frain-dev/convoy/config"
	"github.com/frain-dev/convoy/datastore"
)

type RetryStrategy interface {
	// NextDuration is how long we should wait before next retry
	NextDuration(attempts uint64) time.Duration
}

func NewRetryStrategyFromMetadata(m datastore.Metadata) RetryStrategy {
	if string(m.Strategy) == string(config.ExponentialBackoffStrategyProvider) {
		// 10 seconds to 15 mins
		return NewExponential(m.BackoffTimes)
	}

	return NewDefault(m.IntervalSeconds)
}
