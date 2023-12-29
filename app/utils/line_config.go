package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

var Bot *linebot.Client

// executed if config package is imported
func init() {
	InitLineBot()
}

func InitLineBot() {
	var err error
	if os.Getenv("AWS_EXECUTION_ENV") == "AWS_Lambda" {
		// AWS Lambdaの場合、Parameter Storeから環境変数を読み込む
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		svc := ssm.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

		// LINE_CHANNEL_SECRETの読み込み
		secretName := "/app/line_channel_secret"
		param, err := svc.GetParameter(&ssm.GetParameterInput{
			Name:           aws.String(secretName),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			log.Fatalf("Failed to fetch parameter %s: %v", secretName, err)
		}
		lineChannelSecret := *param.Parameter.Value

		// LINE_ACCESS_TOKENの読み込み
		secretName = "/app/line_access_token"
		param, err = svc.GetParameter(&ssm.GetParameterInput{
			Name:           aws.String(secretName),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			log.Fatalf("Failed to fetch parameter %s: %v", secretName, err)
		}
		lineAccessToken := *param.Parameter.Value

		// 本番環境の場合、parameter storeから環境変数を読み込む
		Bot, err = linebot.New(
			lineChannelSecret,
			lineAccessToken,
		)
		if err != nil {
			log.Fatalf("Failed to create LINE Bot client: %v", err)
		}
	} else {
		// ローカル環境の場合、.envファイルから環境変数を読み込む
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		// Initialize the LINE Bot client using credentials from environment variables.
		Bot, err = linebot.New(
			os.Getenv("LINE_CHANNEL_SECRET"),
			os.Getenv("LINE_ACCESS_TOKEN"),
		)
		if err != nil {
			log.Fatalf("Failed to create LINE Bot client: %v", err)
		}
	}
}
