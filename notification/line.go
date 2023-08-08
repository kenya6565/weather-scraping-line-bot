package line

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bot, err = linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatalf("Failed to create LINE Bot client: %v", err)
	}
}

var userIds []string

func HandleEvent(event *linebot.Event) {
	switch event.Type {
	case linebot.EventTypeFollow:
		userIds = append(userIds, event.Source.UserID)
	}
}

func NotifyRain(userId, message string) error {
	if _, err := bot.PushMessage(userId, linebot.NewTextMessage(message)).Do(); err != nil {
		return err
	}
	return nil
}

func NotifyAllUsers(message string) {
	for _, userId := range userIds {
		NotifyRain(userId, message)
	}
}
