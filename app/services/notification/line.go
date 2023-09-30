package line

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kenya6565/weather-scraping-line-bot/app/db"
	weather "github.com/kenya6565/weather-scraping-line-bot/app/services/weather"
	yokohama "github.com/kenya6565/weather-scraping-line-bot/app/services/weather/yokohama"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Bot represents the global LINE Bot client instance.
var Bot *linebot.Client

// YOKOHAMAWESTAREACODE is a constant defining the area code for Yokohama West for weather information.
const YOKOHAMAWESTAREACODE = "140020"

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
	city := "yokohama" // この部分を変更するだけで異なる都市の天気情報を取得可能

	processor, err := weather.GetWeatherProcessorForCity(city)
	if err != nil {
		log.Println("Failed to get weather processor for city:", err)
		return
	}

	// Fetch the latest weather report.
	weatherReport, err := processor.FetchDataFromJMA()
	log.Println("weatherReport:", weatherReport)
	if err != nil {
		log.Println("Failed to fetch weather report:", err)
		return
	}

	// Filter out relevant weather information based on the Yokohama West area code.
	areas, timeSeriesInfos := yokohama.FilterAreas(weatherReport, YOKOHAMAWESTAREACODE)

	// Process the fetched weather information to generate user-friendly messages.
	messages := yokohama.ProcessAreaInfos(areas, timeSeriesInfos)

	// if all precipitation probabilities do not meet the condition to line, log and return
	if len(messages) == 0 {
		log.Println("All precipitation probabilities are less than 50%. No notification sent.")
		return
	}

	// Combine the generated messages into a single message.
	combinedMessage := strings.Join(messages, "\n")

	// Send the combined weather message to the user.
	if _, err := Bot.PushMessage(userId, linebot.NewTextMessage(combinedMessage)).Do(); err != nil {
		log.Println("Failed to send weather notification:", err)
	}
}
