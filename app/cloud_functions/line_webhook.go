package cloud_functions

// import (
// 	"fmt"
// 	"net/http"
// 	notification "github.com/kenya6565/weather-scraping-line-bot/app/notification"

// )

// func LineWebhookFunction(w http.ResponseWriter, r *http.Request) {
// 	events, _ := notification.Bot.ParseRequest(r)
// 	for _, event := range events {
// 		notification.HandleEvent(event)
// 	}
// 	fmt.Fprint(w, "Line Webhook function executed!")
// }
