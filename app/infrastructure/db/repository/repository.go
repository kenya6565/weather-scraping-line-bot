package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/db"
)

func StoreUserID(ctx context.Context, userID string) (*firestore.DocumentRef, error) {
	docRef := db.Client.Collection("users").Doc(userID)
	_, err := docRef.Set(ctx, map[string]interface{}{
		"UserId": userID,
	}, firestore.MergeAll)
	if err != nil {
		log.Println("Failed adding user:", err)
		return nil, err
	}
	return docRef, nil
}

func StoreCityInfo(userID string, details map[string]interface{}) (*firestore.DocumentRef, error) {
	ctx := db.CreateContext()
	docRef := db.Client.Collection("users").Doc(userID)
	_, err := docRef.Set(ctx, details, firestore.MergeAll)
	if err != nil {
		log.Println("Failed adding user details:", err)
		return nil, err
	}
	return docRef, nil
}

func GetAllUsers() ([]map[string]interface{}, error) {
	ctx := db.CreateContext()
	iter := db.Client.Collection("users").Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		log.Printf("Failed to get user documents: %v", err)
		return nil, err
	}

	var users []map[string]interface{}
	for _, doc := range docs {
		users = append(users, doc.Data())
	}

	return users, nil
}
