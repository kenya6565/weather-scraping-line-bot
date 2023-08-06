package main

import (
	"fmt"
	"weather-scraping-line-bot/weather"
)

const YokohamaWestAreaCode = "140020"

func main() {
	weatherReport, err := weather.FetchWeatherReport()
	if err != nil {
		fmt.Println(err)
		return
	}

	areas, timeSeriesInfos := weather.FilterAreas(weatherReport, YokohamaWestAreaCode)
	weather.ProcessAreaInfos(areas, timeSeriesInfos)
}
