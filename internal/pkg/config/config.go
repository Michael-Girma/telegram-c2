package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BotAPIToken   string
	ClientAPPID   int
	ClientAPPHash string
	Phone         string
	Password      string
}

func NewConfig() *Config {
	env := os.Getenv("Environment")
	if env != "prod" && env != "test" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}

	var appId, err = strconv.Atoi(os.Getenv("CLIENT_APP_ID"))
	if err != nil {
		log.Fatalf("Couldn't parse the appID for telegram client")
	}

	return &Config{
		BotAPIToken:   os.Getenv("BOT_API_TOKEN"),
		ClientAPPID:   appId,
		ClientAPPHash: os.Getenv("CLIENT_APP_HASH"),
		Phone:         os.Getenv("PHONE"),
		Password:      os.Getenv("PASSWORD"),
	}
}
