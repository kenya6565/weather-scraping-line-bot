package db

import (
	"context"
	"log"
	"cloud.google.com/go/firestore"
)

const FirebaseProjectID = "weather-notification-line-dev"

// SaveUserIDToFirestore saves a given user ID to the Firestore.
func SaveUserIDToFirestore(userID string) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, FirebaseProjectID)
	if err != nil {
		return err
	}
	defer client.Close()

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"userId": userID,
	})
	if err != nil {
		log.Println("Failed adding user:", err)
		return err
	}
	return nil
}
