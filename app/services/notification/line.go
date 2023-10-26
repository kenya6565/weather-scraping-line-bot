package line

import (
	"log"
	"strings"

	"github.com/kenya6565/weather-scraping-line-bot/app/db"
	weather "github.com/kenya6565/weather-scraping-line-bot/app/services/weather"
	config "github.com/kenya6565/weather-scraping-line-bot/app/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

// HandleEvent handles incoming Line bot events. It performs actions based on the event type.
// Currently, it supports 'follow' events where a user starts following the bot.
func HandleEvent(event *linebot.Event) {
	switch event.Type {
	case linebot.EventTypeFollow:
		// Store the user's ID to Firestore when the user follows the bot.
		err := db.StoreUserID(event.Source.UserID)
		if err != nil {
			log.Println("Failed to save user ID to Firestore:", err)
		}

		// Send a weather update to the user based on their location.
		// NotifyWeatherToUser(event.Source.UserID, city)

	case linebot.EventTypeMessage:
		if message, ok := event.Message.(*linebot.TextMessage); ok {
			log.Printf("Received message from user %s: %s", event.Source.UserID, message.Text)
		}
	}
}

// // このサンプルでは、緯度と経度をハードコーディングしている都市のリストを使用します。
// // 実際の実装では、都市のデータベースやAPIからこの情報を取得することができます。
// var cities = []CityLocation{
// 	{Name: "yokohama", Latitude: 35.4437, Longitude: 139.6380},
// 	// TODO: 他の都市の緯度と経度を追加
// }

// // determineNearestCity determines the nearest city based on given latitude and longitude.
// func determineNearestCity(lat, long float64) (string, error) {
// 	var nearestCity string
// 	smallestDistance := math.MaxFloat64 // 初期値として非常に大きな値を設定

// 	for _, city := range cities {
// 		distance := calculateDistance(lat, long, city.Latitude, city.Longitude)
// 		if distance < smallestDistance {
// 			smallestDistance = distance
// 			nearestCity = city.Name
// 		}
// 	}

// 	if nearestCity == "" {
// 		return "", fmt.Errorf("couldn't determine the nearest city for the given location")
// 	}

// 	return nearestCity, nil
// }

// // calculateDistance calculates the distance between two geographic coordinates using the Haversine formula.
// func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
// 	const R = 6371.0 // Earth's radius in kilometers

// 	dLat := (lat2 - lat1) * (math.Pi / 180.0)
// 	dLon := (lon2 - lon1) * (math.Pi / 180.0)

// 	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1*(math.Pi/180.0))*math.Cos(lat2*(math.Pi/180.0))*math.Sin(dLon/2)*math.Sin(dLon/2)
// 	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

// 	distance := R * c
// 	return distance
// }

// NotifyWeatherToUser sends a weather report to a specified user.
// It fetches the latest weather report, processes it to generate a user-friendly message,
// and then sends this message to the provided user ID.
func NotifyWeatherToUser(userId, city string) {
	processor, err := weather.GetWeatherProcessorForCity(city)
	if err != nil {
		log.Println("Failed to get weather processor for city:", err)
		return
	}

	weatherReport, err := processor.FetchDataFromJMA()
	if err != nil {
		log.Println("Failed to fetch weather report:", err)
		return
	}
	log.Println("weatherReport:", weatherReport)

	areas, timeSeriesInfos := processor.FilterAreas(weatherReport)
	messages := processor.ProcessAreaInfos(areas, timeSeriesInfos)

	if len(messages) == 0 {
		log.Println("All precipitation probabilities are less than 50%. No notification sent.")
		return
	}

	combinedMessage := strings.Join(messages, "\n")
	if _, err := config.Bot.PushMessage(userId, linebot.NewTextMessage(combinedMessage)).Do(); err != nil {
		log.Println("Failed to send weather notification:", err)
	}
}
