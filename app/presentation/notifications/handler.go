package notifications

import (
	"encoding/json"
	"log"

	domain "github.com/kenya6565/weather-scraping-line-bot/app/domain/notification"
	"github.com/kenya6565/weather-scraping-line-bot/app/services/notification"
	"github.com/line/line-bot-sdk-go/linebot"
)

func HandleEvent(body string) {
	var lineEvent domain.LineEvent
	err := json.Unmarshal([]byte(body), &lineEvent)
	if err != nil {
		log.Printf("Error unmarshalling event: %v", err)
		return
	}

	for _, event := range lineEvent.Events {
		switch event.Type {
		case "follow":
			lineEvent := &linebot.Event{
				Type: linebot.EventTypeFollow,
				Source: &linebot.EventSource{
					UserID: event.Source.UserID,
				},
			}
			notification.HandleFollowEvent(lineEvent)

		case "message":
			if event.Message.Type == "text" {
				lineEvent := &linebot.Event{
					Type: linebot.EventTypeMessage,
					Message: &linebot.TextMessage{
						Text: event.Message.Text,
					},
					Source: &linebot.EventSource{
						UserID: event.Source.UserID,
					},
				}
				notification.HandleMessageEvent(lineEvent)
			}
		}
	}
}
