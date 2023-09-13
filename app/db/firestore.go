package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// FirebaseProjectID holds the project ID for the Firestore database.
const FirebaseProjectID = "weather-notification-line-dev"

// StoreUserID saves a given user ID to the Firestore.
// This function initializes a Firestore client, and stores the user ID
// into a "users" collection. If the operation is successful, it returns nil.
func StoreUserID(userID string) error {
	// Create a new context for Firestore operations.
	ctx := context.Background()

	// Initialize a Firestore client.
	client, err := firestore.NewClient(ctx, FirebaseProjectID)
	if err != nil {
		return err
	}
	// Ensure the client is properly closed after operations are complete.
	defer client.Close()

	// Add the provided userID to the "users" collection in Firestore.
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
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
