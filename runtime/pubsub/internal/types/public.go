package types

import (
	"context"
	"time"
)

// SubscriptionConfig is used when creating a subscription
type SubscriptionConfig struct {
	// Filter is a boolean expression using =, !=, IN, &&
	// It is used to filter which messages are forwarded from the
	// topic to a subscription
	// Filter string - Filters are not currently supported

	// AckDeadline is the time a consumer has to process a message
	// before it's returned to the subscription
	AckDeadline time.Duration

	// MessageRetention is how long an undelivered message is kept
	// on the topic before it's purged
	MessageRetention time.Duration

	// RetryPolicy defines how a message should be retried when
	// the subscriber returns an error
	RetryPolicy *RetryPolicy
}

// RetryPolicy defines how a subscription should handle retries
// after errors either delivering the message or processing the message.
//
// The values given to this structure are parsed at compile time, such that
// the correct Cloud resources can be provisioned to support the queue.
//
// As such the values given here may be clamped to the supported values by
// the target cloud. (i.e. min/max values brought within the supported range
// by the target cloud).
type RetryPolicy struct {
	// The minimum time to wait between retries. Defaults to 10 seconds.
	MinRetryDelay time.Duration

	// The maximum time to wait between retries. Defaults to 10 minutes.
	MaxRetryDelay time.Duration

	// MaxRetries is used to control deadletter queuing logic, when:
	//   n == 0: A default value of 100 retries will be used
	//   n > 0:  Encore will forward a message to a dead letter queue after n retries
	//   n == pubsub.InfiniteRetries: Messages will not be forwarded to the dead letter queue by the Encore framework
	MaxRetries int
}

const (
	NoRetries       = -2
	InfiniteRetries = -1
)

// Subscriber is a function reference
// The signature must be `func(context.Context, msg M) error` where M is either the
// message type of the topic or RawMessage
type Subscriber[T any] func(ctx context.Context, msg T) error

// DeliveryGuarantee is used to configure the delivery contract for a topic
type DeliveryGuarantee int

const (
	// AtLeastOnce guarantees that a message for a subscription is delivered to
	// a subscriber at least once
	AtLeastOnce DeliveryGuarantee = iota + 1

	// ExactlyOnce guarantees that a message for a subscription is delivered to
	// a subscriber exactly once
	// ExactlyOnce // - ExactlyOnce is currently not supported.
)

// TopicConfig is used when creating a Topic
type TopicConfig struct {
	// DeliveryGuarantee is used to configure the delivery guarantee of a Topic
	DeliveryGuarantee DeliveryGuarantee

	// OrderingKey is the name of the message attribute used to group
	// messages and delivery messages with the same OrderingKey value
	// in the order they were published.
	//
	// If OrderingKey is not set, messages can be delivered in any order.
	// OrderingKey string - OrderingKey is currently not supported.
}