package gcppubsub

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func (c *client) GetOrCreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error) {
	topic := c.client.Topic(topicID)
	ok, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if topic exists: %v", err)
	}
	if !ok {
		createdTopic, err := c.client.CreateTopic(ctx, topicID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}

		log.Printf("Topic %s created.\n", topicID)

		return createdTopic, nil
	}

	return topic, nil
}

func (c *client) PublishMessage(ctx context.Context, topic *pubsub.Topic, msg []byte) error {
	result := topic.Publish(ctx,
		&pubsub.Message{
			Data: msg,
		},
	)

	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	log.Println("Published message", "id", id)

	return nil
}
