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

func ProcessAreaInfos(areas []AreaInfo, timeSeriesInfos []TimeSeriesInfo) {
	for i, area := range areas {
		PrintPrecipProb(area, timeSeriesInfos[i])
	}
}

func PrintPrecipProb(area AreaInfo, timeSeries TimeSeriesInfo) {
	if len(*area.Pops) < 2 || len(timeSeries.TimeDefines) < 2 {
		// Skip if there is less than 2 values of arrays in pops and time defines
		return
	}
	for i, popStr := range (*area.Pops)[1:] { // Skip the first pop
		// converting string to int
		pop, err := strconv.Atoi(popStr)
		if err != nil {
			fmt.Println("Error converting pop to integer: ", err)
			return
		}
		if pop >= 50 {
			// Skip the first time define
			timeDefine := timeSeries.TimeDefines[i+1]
			// converting to time.Time type
			parsedTime, err := time.Parse(time.RFC3339, timeDefine)
			if err != nil {
				fmt.Println("Error parsing time: ", err)
				return
			}
			jst := time.FixedZone("Asia/Tokyo", 9*60*60)
			fmt.Printf("Time: %s, Precipitation Probability: %d\n", parsedTime.In(jst).Format("2006-01-02 15:04"), pop)
		}
	}
}
