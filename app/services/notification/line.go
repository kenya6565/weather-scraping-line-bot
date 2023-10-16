package line

import (
	"log"
	"strings"

	"github.com/kenya6565/weather-scraping-line-bot/app/db"
	weather "github.com/kenya6565/weather-scraping-line-bot/app/services/weather"
	config "github.com/kenya6565/weather-scraping-line-bot/app/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

// HandleEvent handles incoming Line bot events. It performs actions based on the event type.
// Currently, it supports 'follow' events where a user starts following the bot.
func HandleEvent(event *linebot.Event) {
	switch event.Type {
	case linebot.EventTypeFollow:
		// Store the user's ID to Firestore when the user follows the bot.
		err := db.StoreUserID(event.Source.UserID)
		if err != nil {
			log.Println("Failed to save user ID to Firestore:", err)
		}

		// Send a welcome message to the user.
		welcomeMessage := "Thank you for following our bot! We will provide you with weather updates."
		if _, err := config.Bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(welcomeMessage)).Do(); err != nil {
			log.Println("Failed to send welcome message:", err)
		}

		// Send a weather update to the user.
		NotifyWeatherToUser(event.Source.UserID)
	}
}

// NotifyWeatherToUser sends a weather report to a specified user.
// It fetches the latest weather report, processes it to generate a user-friendly message,
// and then sends this message to the provided user ID.
func NotifyWeatherToUser(userId string) {
	city := "yokohama"

	processor, err := weather.GetWeatherProcessorForCity(city)
	if err != nil {
		log.Println("Failed to get weather processor for city:", err)
		return
	}

	weatherReport, err := processor.FetchDataFromJMA()
	if err != nil {
		log.Println("Failed to fetch weather report:", err)
		return
	}
	log.Println("weatherReport:", weatherReport)

	areas, timeSeriesInfos := processor.FilterAreas(weatherReport)
	messages := processor.ProcessAreaInfos(areas, timeSeriesInfos)

	if len(messages) == 0 {
		log.Println("All precipitation probabilities are less than 50%. No notification sent.")
		return
	}

	combinedMessage := strings.Join(messages, "\n")
	if _, err := config.Bot.PushMessage(userId, linebot.NewTextMessage(combinedMessage)).Do(); err != nil {
		log.Println("Failed to send weather notification:", err)
	}
}
