package weather

import (
	"fmt"
	"strconv"
	"time"

	model "github.com/kenya6565/weather-scraping-line-bot/app/model"
)

// GenerateRainMessages generates messages based on precipitation probabilities.
func GenerateRainMessages(timeSeriesInfos []model.TimeSeriesInfo) []string {
	var messages []string

	for _, series := range timeSeriesInfos {
		for i, popStr := range series.Areas[0].Pops {
			pop, err := strconv.Atoi(popStr)
			// if err != nil || pop <= 0 {
			// 	continue // Skip if conversion fails or pop is below 20
			// }

			startTime, endTime, err := getTimeRange(series.TimeDefines[i])
			if err != nil {
				continue // Skip if error in getting time range
			}

			message := fmt.Sprintf("時間: %s ~ %s, 降水確率: %d%%", startTime, endTime, pop)
			messages = append(messages, message)
		}
	}

	return messages
}

// APIから取得した時間表記を日本時間に修正する
func getTimeRange(timeDefine string) (string, string, error) {
	parsedTime, err := time.Parse(time.RFC3339, timeDefine)
	if err != nil {
		return "", "", err
	}
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	startTime := parsedTime.In(jst).Format("2006-01-02 15:04")
	endTime := parsedTime.In(jst).Add(5*time.Hour + 59*time.Minute).Format("15:04")
	return startTime, endTime, nil
}
