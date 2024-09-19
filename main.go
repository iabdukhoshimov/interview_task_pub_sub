package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/iabdukhoshimov/interview_task_pub_sub/config"
	"github.com/iabdukhoshimov/interview_task_pub_sub/models"
	gcppubsub "github.com/iabdukhoshimov/interview_task_pub_sub/pkg/pubsub_queue/gcp_pubsub"
)

func main() {
	cfg := config.LoadConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	os.Setenv("PUBSUB_EMULATOR_HOST", cfg.GcpPubsub.Endpiont)

	gcpClient := gcppubsub.New(ctx)
	defer gcpClient.Close()

	topic, errTopic := gcpClient.GetOrCreateTopic(ctx, cfg.GcpPubsub.TopicID)
	if errTopic != nil {
		log.Fatal(errTopic)
	}

	sub, errSub := gcpClient.GetOrCreateSubscription(ctx, topic, cfg.GcpPubsub.SubName)
	if errSub != nil {
		log.Fatal(errSub)
	}

	var (
		messaage models.Message
		product  = models.ProductMessage{
			ID:   "1",
			Type: "product",
			Name: "Product 1",
			Image: models.Image{
				URL:    "https://example.com/image.jpg",
				Width:  100,
				Height: 100,
			},
			Thumbnail: models.Image{
				URL:    "https://example.com/thumbnail.jpg",
				Width:  50,
				Height: 50,
			},
		}
	)

	productData, errMarshal := json.Marshal(product)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}

	messaage.Data = productData

	bytes, errMarshal := json.Marshal(messaage)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}

	errPubliish := gcpClient.PublishMessage(ctx, topic, bytes)
	if errPubliish != nil {
		log.Fatal(errPubliish)
	}

	errConsume := gcpClient.ConsumeMessages(ctx, sub, func(data []byte) error {
		var product models.ProductMessage
		if errUnmarshal := json.Unmarshal(data, &product); errUnmarshal != nil {
			return errUnmarshal
		}

		log.Printf("Received product: %v\n", product)

		return nil
	})
	if errConsume != nil {
		log.Fatal(errConsume)
	}
}
