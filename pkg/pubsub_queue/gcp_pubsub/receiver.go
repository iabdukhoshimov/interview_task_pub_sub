package gcppubsub

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/iabdukhoshimov/interview_task_pub_sub/models"
)

func (c *client) GetOrCreateSubscription(
	ctx context.Context,
	topic *pubsub.Topic,
	subName string,
) (*pubsub.Subscription, error) {
	sub := c.client.Subscription(subName)
	ok, errSubExists := sub.Exists(ctx)
	if errSubExists != nil {
		log.Fatalf("Failed to check if subscription exists: %v", errSubExists)
	}
	if !ok {
		createdSub, errCreateSub := c.client.
			CreateSubscription(
				ctx, subName, pubsub.SubscriptionConfig{
					Topic: topic,
				},
			)
		if errCreateSub != nil {
			log.Fatalf("Failed to create subscription: %v", errCreateSub)
		}

		log.Printf("Subscription %s created.\n", subName)

		return createdSub, nil
	}

	return sub, nil
}

func (c *client) ConsumeMessages(ctx context.Context, sub *pubsub.Subscription, handler func([]byte) error) error {
	errReceive := sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		var m models.Message
		if errUnmarshal := json.Unmarshal(msg.Data, &m); errUnmarshal != nil {
			log.Printf("Failed to unmarshal message data: %v", errUnmarshal)
			msg.Nack()
			return
		}

		if errHandle := handler(m.Data); errHandle != nil {
			log.Printf("Failed to handle message: %v", errHandle)
			msg.Nack()
			return
		}

		msg.Ack()
	})
	if errReceive != nil {
		log.Fatalf("Failed to receive messages: %v", errReceive)
	}

	return nil
}

func (c *client) Close() error {
	return nil
}
