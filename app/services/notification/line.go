package line

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kenya6565/weather-scraping-line-bot/app/db"
	weather "github.com/kenya6565/weather-scraping-line-bot/app/services/weather"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Bot represents the global LINE Bot client instance.
var Bot *linebot.Client

// init is a special Go function that is executed upon package initialization.
// This function is responsible for loading the .env configuration and initializing the LINE Bot client.
func init() {
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the LINE Bot client using credentials from environment variables.
	Bot, err = linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatalf("Failed to create LINE Bot client: %v", err)
	}
}

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
		if _, err := Bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(welcomeMessage)).Do(); err != nil {
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

	areas, timeSeriesInfos := processor.FilterAreas(*weatherReport)
	messages := processor.ProcessAreaInfos(areas, timeSeriesInfos)

	if len(messages) == 0 {
		log.Println("All precipitation probabilities are less than 50%. No notification sent.")
		return
	}

	combinedMessage := strings.Join(messages, "\n")
	if _, err := Bot.PushMessage(userId, linebot.NewTextMessage(combinedMessage)).Do(); err != nil {
		log.Println("Failed to send weather notification:", err)
	}
}
