package server

import (
	"log"
	"net/http"

	line "github.com/kenya6565/weather-scraping-line-bot/app/services/notification"
	config "github.com/kenya6565/weather-scraping-line-bot/app/utils" // 追加
)

func handleCallback(w http.ResponseWriter, r *http.Request) {
	events, _ := config.Bot.ParseRequest(r)
	for _, event := range events {
		line.HandleEvent(event)
	}
}

func StartServer() {
	// execute handleCallback when receiving request /callback
	http.HandleFunc("/callback", handleCallback)
	log.Println("Starting server on :8080...")
	// activate http server
	http.ListenAndServe(":8080", nil)
}
