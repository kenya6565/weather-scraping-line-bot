package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

var Bot *linebot.Client

// executed if config package is imported
func init() {
	InitLineBot()
}

func InitLineBot() {
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the LINE Bot client using credentials from environment variables.
	Bot, err = linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatalf("Failed to create LINE Bot client: %v", err)
	}
}
