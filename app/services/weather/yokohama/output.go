package yokohama

import (
	"fmt"
	"strconv"
	"time"

	weather "github.com/kenya6565/weather-scraping-line-bot/app/models/weather"
)

func GeneratePrecipProbMessage(area weather.AreaInfoInterface, timeSeries weather.TimeSeriesInfo) []string {
	var messages []string

	if len(*area.GetPops()) < 2 || len(timeSeries.TimeDefines) < 2 {
		return messages
	}

	for i, popStr := range (*area.GetPops())[1:] {
		pop, err := strconv.Atoi(popStr)
		if err != nil {
			fmt.Printf("Error converting pop to integer: %v\n", err)
			continue
		}
		if pop >= 20 {
			timeDefine := timeSeries.TimeDefines[i+1]
			startTime, endTime, err := getTimeRange(timeDefine)
			if err != nil {
				fmt.Printf("Error getting time range: %v\n", err)
				continue
			}

			message := fmt.Sprintf("時間: %s ~ %s, 降水確率: %d%%", startTime, endTime, pop)
			messages = append(messages, message)
		}
	}
	return messages
}

func getTimeRange(timeDefine string) (string, string, error) {
	parsedTime, err := time.Parse(time.RFC3339, timeDefine)
	if err != nil {
		return "", "", err
	}
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	startTime := parsedTime.In(jst)
	var endTime time.Time
	switch startTime.Hour() {
	case 0, 6, 12, 18:
		endTime = startTime.Add(time.Hour*5 + time.Minute*59)
	default:
		return "", "", fmt.Errorf("unexpected hour: %d", startTime.Hour())
	}
	return startTime.Format("2006-01-02 15:04"), endTime.Format("15:04"), nil
}
