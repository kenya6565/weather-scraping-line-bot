package main

import (
	"log"
	"os"

	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/db"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/notification"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/server"
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
		server.ActivateLambda()
	} else {
		server.ActivateLocalServer()
	}
}
