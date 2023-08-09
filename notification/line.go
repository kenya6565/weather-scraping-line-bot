package line

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/api/iterator"
)

var Bot *linebot.Client

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
		log.Println("Received follow event!")

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
			log.Printf("Failed adding user: %v", err)
		} else {
			log.Println("Successfully added user to Firestore.")
		}
	}
}

func NotifyRain(userId, message string) error {
	if _, err := Bot.PushMessage(userId, linebot.NewTextMessage(message)).Do(); err != nil {
		return err
	}
	return nil
}

func NotifyAllUsers(message string) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, FirebaseProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		userId := doc.Data()["lineUserId"].(string)
		NotifyRain(userId, message)
	}
}
