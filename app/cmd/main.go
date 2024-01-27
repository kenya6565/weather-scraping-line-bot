package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/db"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/notification"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/server"
	n "github.com/kenya6565/weather-scraping-line-bot/app/services/notification"
)

func main() {
	// output log when something goes wrong
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()
	db.InitFirestoreClient()
	notification.InitLineBot()

	// prd env
	if os.Getenv("AWS_EXECUTION_ENV") == "AWS_Lambda" {
		log.Println("Running on AWS Lambda")
		lambda.Start(n.NotifyWeatherToAllUsers)
	} else {
		server.ActivateLocalServer()
	}
}
