package server

import (
	"io"
	"log"
	"net/http"

	"github.com/kenya6565/weather-scraping-line-bot/app/presentation/notifications"
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	notifications.HandleLineEvent(string(body))
}

func ActivateLocalServer() {
	// execute handleCallback when receiving request /callback
	http.HandleFunc("/callback", HandleCallback)
	log.Println("Starting server on :8080...")
	// activate http server
	http.ListenAndServe(":8080", nil)
}
