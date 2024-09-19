package pubsubqueue

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Client interface {
	GetOrCreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error)
	GetOrCreateSubscription(ctx context.Context, topic *pubsub.Topic, subName string) (*pubsub.Subscription, error)
	PublishMessage(ctx context.Context, topic *pubsub.Topic, msg []byte) error
	ConsumeMessages(ctx context.Context, sub *pubsub.Subscription, handler func([]byte) error) error
	Close() error
}
