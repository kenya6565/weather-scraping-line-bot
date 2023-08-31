package server

import (
	"log"
	"net/http"

	notification "github.com/kenya6565/weather-scraping-line-bot/app/notification"
)

func handleCallback(w http.ResponseWriter, r *http.Request) {
	events, _ := notification.Bot.ParseRequest(r)
	for _, event := range events {
		notification.HandleEvent(event)
	}
}

func StartServer() {
	// execute handleCallback when receiving request /callback
	http.HandleFunc("/callback", handleCallback)
	log.Println("Starting server on :8080...")
	// activate http server
	http.ListenAndServe(":8080", nil)
}
