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

func CreateContext() context.Context {
	return context.Background()
}

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
		googleApplicationCredentials = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	}

	// Create a new context for Firestore operations.
	ctx := CreateContext()

	// Initialize a Firestore client.
	// ここでGOOGLE_APPLICATION_CREDENTIALSを読み取っている
	opt := option.WithCredentialsJSON([]byte(googleApplicationCredentials))
	client, err := firestore.NewClient(ctx, FirebaseProjectID, opt)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
		return err
	}

	Client = client
	return nil
}
