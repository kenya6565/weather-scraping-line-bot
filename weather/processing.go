package weather

import (
	"fmt"
	"strconv"
	"time"
)

func FilterAreas(weatherReport []WeatherInfo, code string) ([]AreaInfo, []TimeSeriesInfo) {
	var areas []AreaInfo
	var timeSeriesInfos []TimeSeriesInfo
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

func ProcessAreaInfos(areas []AreaInfo, timeSeriesInfos []TimeSeriesInfo) []string {
	var messages []string
	for i, area := range areas {
		messages = append(messages, GeneratePrecipProbMessage(area, timeSeriesInfos[i])...)
	}
	return messages
}

func GeneratePrecipProbMessage(area AreaInfo, timeSeries TimeSeriesInfo) []string {
	var messages []string

	// Skip if there is less than 2 values of arrays in pops and time defines
	if len(*area.Pops) < 2 || len(timeSeries.TimeDefines) < 2 {
		return messages
	}

	for i, popStr := range (*area.Pops)[1:] { // Skip the first pop
		// converting string to int
		pop, err := strconv.Atoi(popStr)
		if err != nil {
			fmt.Println("Error converting pop to integer: ", err)
			continue
		}
		if pop >= 20 {
			// Skip the first time define
			timeDefine := timeSeries.TimeDefines[i+1]
			// converting to time.Time type
			parsedTime, err := time.Parse(time.RFC3339, timeDefine)
			if err != nil {
				fmt.Println("Error parsing time: ", err)
				continue
			}
			jst := time.FixedZone("Asia/Tokyo", 9*60*60)
			message := fmt.Sprintf("Time: %s, Precipitation Probability: %d", parsedTime.In(jst).Format("2006-01-02 15:04"), pop)
			messages = append(messages, message)
		}
	}
	return messages
}
