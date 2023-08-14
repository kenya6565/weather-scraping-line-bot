package line

import (
	"context"
	"log"
	"os"
	"strings"
	"weather-scraping-line-bot/weather"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

var Bot *linebot.Client

const YOKOHAMAWESTAREACODE = "140020"
const FirebaseProjectID = "weather-notification-line-dev"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Bot, err = linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatalf("Failed to create LINE Bot client: %v", err)
	}
}

func HandleEvent(event *linebot.Event) {
	switch event.Type {
	case linebot.EventTypeFollow:
		ctx := context.Background()

		// Firestore clientの作成
		client, err := firestore.NewClient(ctx, FirebaseProjectID)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		defer client.Close()

		// Firestoreへの保存処理
		_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
			"userId": event.Source.UserID,
		})
		if err != nil {
			log.Fatalf("Failed adding user: %v", err)
		}

		// フォロー通知
		welcomeMessage := "Thank you for following our bot! We will provide you with weather updates."
		if _, err := Bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(welcomeMessage)).Do(); err != nil {
			log.Println("Failed to send welcome message:", err)
		}

		// 天気予報の通知
		NotifyWeatherToUser(event.Source.UserID)
	}
}

func NotifyWeatherToUser(userId string) {
	weatherReport, err := weather.FetchWeatherReport()
	if err != nil {
		log.Println("Failed to fetch weather report:", err)
		return
	}

	areas, timeSeriesInfos := weather.FilterAreas(weatherReport, YOKOHAMAWESTAREACODE)
	messages := weather.ProcessAreaInfos(areas, timeSeriesInfos) // この関数は文字列のスライスを返すように修正する必要があります

	// メッセージを結合して1つのメッセージとして送信
	combinedMessage := strings.Join(messages, "\n")

	if _, err := Bot.PushMessage(userId, linebot.NewTextMessage(combinedMessage)).Do(); err != nil {
		log.Println("Failed to send weather notification:", err)
	}
}
