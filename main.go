package main

import (
	"log"
	"net/http"

	notification "weather-scraping-line-bot/notification"
)

const YokohamaWestAreaCode = "140020"

func handleCallback(w http.ResponseWriter, r *http.Request) {
	events, _ := notification.Bot.ParseRequest(r)
	for _, event := range events {
		notification.HandleEvent(event)
	}
}

func main() {
	// weatherReport, err := weather.FetchWeatherReport()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// areas, timeSeriesInfos := weather.FilterAreas(weatherReport, YokohamaWestAreaCode)
	// weather.ProcessAreaInfos(areas, timeSeriesInfos)

	// execute handleCallback when receiving request /callback
	http.HandleFunc("/callback", handleCallback)
	log.Println("Starting server on :8080...")
	// activate http server 
	http.ListenAndServe(":8080", nil)
}
