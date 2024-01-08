package main

import (
	"os"

	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/db"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/notification"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/server"
)

func main() {
	db.InitFirestoreClient()
	notification.InitLineBot()

	// prd env
	if os.Getenv("AWS_EXECUTION_ENV") == "AWS_Lambda" {
		// lambda.Start()
		// local env
	} else {
		server.ActivateLocalServer()
	}
}
