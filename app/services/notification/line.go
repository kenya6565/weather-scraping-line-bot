package line

import (
	"log"
	"net/http"
	"strings"

	"github.com/kenya6565/weather-scraping-line-bot/app/db"
	"github.com/kenya6565/weather-scraping-line-bot/app/services/weather"
	config "github.com/kenya6565/weather-scraping-line-bot/app/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	events, _ := config.Bot.ParseRequest(r)
	for _, event := range events {
		HandleEvent(event)
	}
}

// HandleEvent handles incoming Line bot events based on event type.
// Supported events are 'follow' and 'text message'.
func HandleEvent(event *linebot.Event) {
	switch event.Type {
	// when user following me
	case linebot.EventTypeFollow:
		handleFollowEvent(event)

	// when user sending messages to me
	case linebot.EventTypeMessage:
		handleMessageEvent(event)
	}
}

// handleFollowEvent stores the user's ID to Firestore when the user starts following the bot.
func handleFollowEvent(event *linebot.Event) {
	ctx := config.CreateContext()
	err := db.StoreUserID(ctx, event.Source.UserID)
	if err != nil {
		log.Printf("Failed to save user ID %s to Firestore: %v", event.Source.UserID, err)
	}
}

// handleMessageEvent processes the message event.
// When a user sends a city name, it triggers the weather notification.
func handleMessageEvent(event *linebot.Event) {
	if message, ok := event.Message.(*linebot.TextMessage); ok {
		log.Printf("Received message from user %s: %s", event.Source.UserID, message.Text)

		processor, err := weather.GetWeatherProcessorForCity(message.Text)
		if err != nil {
			sendMessageToUser(event.Source.UserID, "申し訳ございませんが、その都市の天気情報はサポートされていません。他の都市名を入力してください。")
			return
		}
		// TODO: ここにDBに都市情報保存するメソッド追加
		// config, ok := processor.(*weather.CityWeatherConfig)
		// if !ok {
		// 	log.Println("Failed to assert type")
		// 	return
		// }
		// ctx := db.CreateContext()
		// // TODO: ここの部分でうまくいっていないので修正する
		// _, _, err = db.Client.Collection("users").Add(ctx, map[string]interface{}{
		// 	"userId":         event.Source.UserID,
		// 	"JmaApiEndpoint": config.JmaApiEndpoint,
		// 	"AreaCode":       config.AreaCode,
		// 	"AreaName":       config.AreaName,
		// })
		// if err != nil {
		// 	log.Println("Failed adding user:", err)
		// 	return
		// }

		NotifyWeatherToUser(event.Source.UserID, message.Text, processor)
	}
}

// NotifyWeatherToUser sends a weather report or an error message to the user.
func NotifyWeatherToUser(userID, city string, processor weather.WeatherProcessor) {
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
	if _, err := config.Bot.PushMessage(userID, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Printf("Failed to send message to user %s: %v", userID, err)
	}
}
