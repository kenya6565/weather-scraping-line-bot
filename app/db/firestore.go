package db

import (
	"context"
	"log"

	config "github.com/kenya6565/weather-scraping-line-bot/app/utils"
)

// StoreUserID saves a given user ID to the Firestore.
// This function initializes a Firestore client, and stores the user ID
// into a "users" collection. If the operation is successful, it returns nil.
func StoreUserID(ctx context.Context, userID string) error {

	// Add the provided userID to the "users" collection in Firestore.
	_, _, err := config.Client.Collection("users").Add(ctx, map[string]interface{}{
		"userId": userID,
	})
	if err != nil {
		// Log the error if unable to add the user to Firestore.
		log.Println("Failed adding user:", err)
		return err
	}

	// Return nil if the user ID is successfully added.
	return nil
}

// TODO: ユーザーから送られた都市情報を保存するメソッド書く
