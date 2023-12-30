package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirestoreClient() error {

	var FirebaseProjectID, googleApplicationCredentials string
	var err error

	if os.Getenv("AWS_EXECUTION_ENV") == "AWS_Lambda" {
		// AWS Lambdaの場合、Parameter Storeから環境変数を読み込む
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		svc := ssm.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

		// FIREBASE_PROJECT_IDの読み込み
		secretName := "/app/firebase_project_id"
		param, err := svc.GetParameter(&ssm.GetParameterInput{
			Name:           aws.String(secretName),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			log.Fatalf("Failed to fetch parameter %s: %v", secretName, err)
		}
		FirebaseProjectID = *param.Parameter.Value

		// GOOGLE_APPLICATION_CREDENTIALSの読み込み
		secretName = "/app/google_application_credentials"
		param, err = svc.GetParameter(&ssm.GetParameterInput{
			Name:           aws.String(secretName),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			log.Fatalf("Failed to fetch parameter %s: %v", secretName, err)
		}
		googleApplicationCredentials = *param.Parameter.Value
	} else {
		// ローカル環境の場合、.envファイルから環境変数を読み込む
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		FirebaseProjectID = os.Getenv("FIREBASE_PROJECT_ID")
	}

	// Create a new context for Firestore operations.
	ctx := CreateContext()

	// Initialize a Firestore client.
	// ここでGOOGLE_APPLICATION_CREDENTIALSを読み取っている
	opt := option.WithCredentialsJSON([]byte(googleApplicationCredentials))
	client, err := firestore.NewClient(ctx, FirebaseProjectID, opt)
	if err != nil {
		return err
	}
	// Ensure the client is properly closed after operations are complete.
	defer client.Close()

	Client = client
	return nil
}

func CreateContext() context.Context {
	return context.Background()
}

// StoreUserID saves a given user ID to the Firestore.
// This function initializes a Firestore client, and stores the user ID
// into a "users" collection. If the operation is successful, it returns nil.
func StoreUserID(ctx context.Context, userID string) error {

	// Add the provided userID to the "users" collection in Firestore.
	_, _, err := Client.Collection("users").Add(ctx, map[string]interface{}{
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
