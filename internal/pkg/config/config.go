package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotAPIToken string
}

func NewConfig() *Config {
	env := os.Getenv("Environment")
	if env != "prod" && env != "test" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}

	return &Config{
		BotAPIToken: os.Getenv("BOT_API_TOKEN"),
	}
}
