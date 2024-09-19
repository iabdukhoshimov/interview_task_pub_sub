package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v9"
	gcppubsub "github.com/iabdukhoshimov/interview_task_pub_sub/pkg/pubsub_queue/gcp_pubsub"
	"github.com/joho/godotenv"
)

type Config struct {
	ProjectName string `env:"PROJECT_NAME" envDefault:"pubsub-queue"`
	GcpPubsub   gcppubsub.Config
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file", err)
	}

	var cfg Config
	errParse := env.Parse(&cfg)
	if errParse != nil {
		log.Fatalf("Failed to parse config: %v", errParse)
	}

	return &cfg
}
