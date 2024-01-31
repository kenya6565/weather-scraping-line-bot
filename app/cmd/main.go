package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/db"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/notification"
	"github.com/kenya6565/weather-scraping-line-bot/app/infrastructure/server"
	"github.com/kenya6565/weather-scraping-line-bot/app/presentation/notifications"
	n "github.com/kenya6565/weather-scraping-line-bot/app/services/notification"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// イベント実行(api gateway)
	if request.Headers != nil {
		// Create a new http.Request
		httpRequest, err := http.NewRequest("POST", "/", bytes.NewBufferString(request.Body))
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 400}, err
		}
		httpRequest.Header.Set("X-Line-Signature", request.Headers["X-Line-Signature"])

		// Parse the incoming event to a Line event
		lineEvents, err := notification.Bot.ParseRequest(httpRequest)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 400}, err
		}

		// Process the Line event
		for _, event := range lineEvents {
			notifications.HandleEvent(event)
		}

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
		// lambda.Start(n.NotifyWeatherToAllUsers)
		lambda.Start(handler)
	} else {
		server.ActivateLocalServer()
	}
}
