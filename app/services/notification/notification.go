package notification

import (
	"log"
	"strings"

	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/db"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/db/repository"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/notification"
	"github.com/kenya6565/weather-scraping-line-bot/app/services/weather"
	"github.com/line/line-bot-sdk-go/linebot"
)

// handleFollowEvent stores the user's ID to Firestore when the user starts following the bot.
func HandleFollowEvent(event *linebot.Event) {
	ctx := db.CreateContext()
	_, err := repository.StoreUserID(ctx, event.Source.UserID)
	if err != nil {
		log.Printf("Failed to save user ID %s to Firestore: %v", event.Source.UserID, err)
	}
}

// handleMessageEvent processes the message event.
// When a user sends a city name, it triggers the weather notification.
func HandleMessageEvent(event *linebot.Event) {
	if message, ok := event.Message.(*linebot.TextMessage); ok {
		log.Printf("Received message from user %s: %s", event.Source.UserID, message.Text)

		processor, err := weather.GetWeatherProcessorForCity(message.Text)
		if err != nil {
			sendMessageToUser(event.Source.UserID, "申し訳ございませんが、その都市の天気情報はサポートされていません。他の都市名を入力してください。")
			return
		}

		cityInfo := map[string]interface{}{
			"CityName":       processor.GetCityName(),
			"JmaApiEndpoint": processor.GetJmaApiEndpoint(),
			"AreaCode":       processor.GetAreaCode(),
			"AreaName":       processor.GetAreaName(),
		}
		if _, err = repository.StoreCityInfo(event.Source.UserID, cityInfo); err != nil {
			log.Println("Failed adding city info:", err)
			return
		}

		NotifyWeatherToUser(event.Source.UserID, message.Text, processor)
			}
}

// NotifyWeatherToUser sends a weather report or an error message to the user.
func NotifyWeatherToUser(userID, city string, processor weather.WeatherProcessor) {
	log.Printf("NotifyWeatherToUser called for user %s and city %s", userID, city)
	weatherReport, err := processor.FetchDataFromJMA()
	if err != nil {
		log.Printf("Failed to fetch weather report for city %s: %v", city, err)
		return
	}
	timeSeriesInfos := processor.TransformWeatherData(weatherReport)
	messages := weather.GenerateRainMessages(timeSeriesInfos)

	// when no precipitation
	if len(messages) == 0 {
		log.Print("All precipitation probabilities for city are less than 50%")
	}

	combinedMessage := city + "の天気予報：\n" + strings.Join(messages, "\n")
	sendMessageToUser(userID, combinedMessage)
}

// sendMessageToUser sends a text message to the specified user.
func sendMessageToUser(userID, message string) {
	if _, err := notification.Bot.PushMessage(userID, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Printf("Failed to send message to user %s: %v", userID, err)
	}
}

func NotifyWeatherToAllUsers() {
	users, err := repository.GetAllUsers()
	if err != nil {
		log.Printf("Failed to get all users: %v", err)
		return
	}

	for _, user := range users {
		userID := user["UserId"].(string)
		city := user["CityName"].(string)
		processor, err := weather.GetWeatherProcessorForCity(city)
		if err != nil {
			log.Printf("Failed to get weather processor for city %s: %v", city, err)
			continue
		}
		NotifyWeatherToUser(userID, city, processor)
	}
}
