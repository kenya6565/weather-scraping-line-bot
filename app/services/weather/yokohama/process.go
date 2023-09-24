package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"
	"github.com/kenya6565/weather-scraping-line-bot/app/services/weather"
)

func FilterAreas(weatherReport []weather.WeatherInfo, code string) ([]weather.AreaInfo, []weather.TimeSeriesInfo) {
	var areas []weather.AreaInfo
	var timeSeriesInfos []weather.TimeSeriesInfo
	for _, info := range weatherReport {
		for _, timeSeries := range info.TimeSeries {
			for _, area := range timeSeries.Areas {
				if area.Area.Code == code && area.Pops != nil {
					areas = append(areas, area)
					timeSeriesInfos = append(timeSeriesInfos, timeSeries)
				}
			}
		}
	}
	return areas, timeSeriesInfos
}

func ProcessAreaInfos(areas []weather.AreaInfo, timeSeriesInfos []weather.TimeSeriesInfo) []string {
	var messages []string
	for i, area := range areas {
		messages = append(messages, GeneratePrecipProbMessage(area, timeSeriesInfos[i])...)
	}
	return messages
}

func GeneratePrecipProbMessage(area weather.AreaInfo, timeSeries weather.TimeSeriesInfo) []string {
	var messages []string

	if len(*area.Pops) < 2 || len(timeSeries.TimeDefines) < 2 {
		return messages
	}

	for i, popStr := range (*area.Pops)[1:] {
		pop, err := strconv.Atoi(popStr)
		if err != nil {
			fmt.Println("Error converting pop to integer: ", err)
			continue
		}
		if pop >= 20 {
			timeDefine := timeSeries.TimeDefines[i+1]
			parsedTime, err := time.Parse(time.RFC3339, timeDefine)
			if err != nil {
				fmt.Println("Error parsing time: ", err)
				continue
			}

			jst := time.FixedZone("Asia/Tokyo", 9*60*60)
			startTime := parsedTime.In(jst)
			var endTime time.Time

			switch startTime.Hour() {
			case 0:
				endTime = startTime.Add(time.Hour * 5).Add(time.Minute * 59)
			case 6:
				endTime = startTime.Add(time.Hour * 5).Add(time.Minute * 59)
			case 12:
				endTime = startTime.Add(time.Hour * 5).Add(time.Minute * 59)
			case 18:
				endTime = startTime.Add(time.Hour * 5).Add(time.Minute * 59)
			default:
				// If the hour does not match the above cases, we skip it.
				continue
			}

			message := fmt.Sprintf("時間: %s ~ %s, 降水確率: %d%%",
				startTime.Format("2006-01-02 15:04"),
				endTime.Format("15:04"),
				pop)
			messages = append(messages, message)
		}
	}
	return messages
}

// ProcessWeatherResponse processes the HTTP response body from the JMA API.
func ProcessWeatherResponse(body io.Reader) ([]weather.WeatherInfo, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var weatherReport []weather.WeatherInfo
	err = json.Unmarshal(data, &weatherReport)
	return weatherReport, err
}
