package main

import (
	"fmt"
	"net/http"
	// notification "github.com/kenya6565/weather-scraping-line-bot/app/notification"

)

func LineWebhookFunction(w http.ResponseWriter, r *http.Request) {
	// events, _ := notification.Bot.ParseRequest(r)
	// for _, event := range events {
	// 	notification.HandleEvent(event)
	// }
	// 応答が必要な場合
	fmt.Fprint(w, "Line Webhook function executed!")
}
