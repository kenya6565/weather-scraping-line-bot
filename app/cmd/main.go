package main

import (
	"os"

	"github.com/kenya6565/weather-scraping-line-bot/app/server"
)

func main() {
	// prd env
	if os.Getenv("AWS_EXECUTION_ENV") == "AWS_Lambda" {
		// lambda.Start()
		// local env
	} else {
		server.ActivateLocalServer()
	}
}
