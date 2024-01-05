package notification

import (
	"net/http"

	"github.com/kenya6565/weather-scraping-line-bot/app/services/notification"
	config "github.com/kenya6565/weather-scraping-line-bot/app/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	events, _ := config.Bot.ParseRequest(r)
	for _, event := range events {
		HandleEvent(event)
	}
}

// HandleEvent handles incoming Line bot events based on event type.
// Supported events are 'follow' and 'text message'.
func HandleEvent(event *linebot.Event) {
	switch event.Type {
	// when user following me
	case linebot.EventTypeFollow:
		notification.HandleFollowEvent(event)

	// when user sending messages to me
	case linebot.EventTypeMessage:
		notification.HandleMessageEvent(event)
	}
}
