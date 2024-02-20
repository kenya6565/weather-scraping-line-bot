package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/db"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/notification"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/server"
	"github.com/kenya6565/weather-scraping-line-bot/app/presentation/notifications"
	n "github.com/kenya6565/weather-scraping-line-bot/app/services/notification"
)

func HandleAPIGatewayRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// イベント実行(api gateway)
	if request.Headers != nil {
		log.Println("Received event: ", request.Body)
		notifications.HandleLineEvent(request.Body)
		return events.APIGatewayProxyResponse{StatusCode: 200}, nil
	}

	// 定期実行(event bridge)
	n.NotifyWeatherToAllUsers()
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

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
		lambda.Start(HandleAPIGatewayRequest)
	} else {
		server.ActivateLocalServer()
	}
}
