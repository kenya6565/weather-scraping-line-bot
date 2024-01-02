package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	config "github.com/kenya6565/weather-scraping-line-bot/app/utils"
)

func StoreUserID(ctx context.Context, userID string) (*firestore.DocumentRef, error) {
	docRef := config.Client.Collection("users").Doc(userID)
	_, err := docRef.Set(ctx, map[string]interface{}{
		"UserId": userID,
	}, firestore.MergeAll)
	if err != nil {
		log.Println("Failed adding user:", err)
		return nil, err
	}
	return docRef, nil
}
