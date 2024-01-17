package server

import (
	"log"
	"net/http"

	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/notification"
	"github.com/kenya6565/weather-scraping-line-bot/app/presentation/notifications"
	n "github.com/kenya6565/weather-scraping-line-bot/app/services/notification"
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	events, _ := notification.Bot.ParseRequest(r)
	for _, event := range events {
		notifications.HandleEvent(event)
	}
}

func ActivateLocalServer() {
	// execute handleCallback when receiving request /callback
	http.HandleFunc("/callback", HandleCallback)
	log.Println("Starting server on :8080...")
	// activate http server
	http.ListenAndServe(":8080", nil)
}

// TODO: Lambda環境で必要なネットワーク処理を記載
func ActivateLambda() {
	n.NotifyWeatherToAllUsers()
}
