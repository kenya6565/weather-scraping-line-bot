package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

// StoreUserID saves a given user ID to the Firestore.
// This function initializes a Firestore client, and stores the user ID
// into a "users" collection. If the operation is successful, it returns nil.
func StoreUserID(userID string) error {
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get FirebaseProjectID from the environment variables.
	FirebaseProjectID := os.Getenv("TF_VAR_FIREBASE_PROJECT_ID")
	if FirebaseProjectID == "" {
		log.Fatal("TF_VAR_FIREBASE_PROJECT_ID must be set in the environment")
	}

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
