package cloud_functions

import (
	"fmt"
	"net/http"
	notification "weather-scraping-line-bot/notification"
)

func LineWebhookFunction(w http.ResponseWriter, r *http.Request) {
	events, _ := notification.Bot.ParseRequest(r)
	for _, event := range events {
		notification.HandleEvent(event)
	}
	// 応答が必要な場合
	fmt.Fprint(w, "Webhook processed!")
}
