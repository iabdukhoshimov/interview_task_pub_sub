package gcppubsub

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/caarlos0/env/v9"
	pubsubqueue "github.com/iabdukhoshimov/interview_task_pub_sub/pkg/pubsub_queue"
	"go.uber.org/zap"
)

type Config struct {
	Endpiont       string        `env:"ENDPOINT,required"`
	ProjectID      string        `env:"PROJECT_ID,required"`
	TopicID        string        `env:"TOPIC_ID,required"`
	SubName        string        `env:"SUB_NAME,required"`
	DefaultTimeout time.Duration `env:"DEFAULT_TIMEOUT" envDefault:"10s"`
}

type client struct {
	client *pubsub.Client
	cfg    *Config
}

/*
New creates a new instance of the GCP PubSub client
and it receives everythig through the environment variables .
*/
func New(ctx context.Context) pubsubqueue.Client {
	var (
		cfg Config
	)

	errParse := env.Parse(&cfg)
	if errParse != nil {
		zap.L().Error("failed to parse env", zap.Error(errParse))
		panic(errParse)
	}

	pubsubClient, errNewClient := pubsub.NewClient(ctx, cfg.ProjectID)
	if errNewClient != nil {
		zap.L().Error("failed to create pubsub client", zap.Error(errNewClient))
		panic(errNewClient)
	}

	return &client{
		cfg:    &cfg,
		client: pubsubClient,
	}
}
