package server

import (
	"log"
	"net/http"

	line "github.com/kenya6565/weather-scraping-line-bot/app/services/notification"
)

func ActivateLocalServer() {
	// execute handleCallback when receiving request /callback
	http.HandleFunc("/callback", line.HandleCallback)
	log.Println("Starting server on :8080...")
	// activate http server
	http.ListenAndServe(":8080", nil)
}

// TODO: Lambda環境で必要なネットワーク処理を記載
// func ActivateLambda() {

// }
