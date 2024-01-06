package server

import (
	"log"
	"net/http"

	notification "github.com/kenya6565/weather-scraping-line-bot/app/presentation/notification"
	config "github.com/kenya6565/weather-scraping-line-bot/app/utils"
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	events, _ := config.Bot.ParseRequest(r)
	for _, event := range events {
		notification.HandleEvent(event)
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
// func ActivateLambda() {

// }
