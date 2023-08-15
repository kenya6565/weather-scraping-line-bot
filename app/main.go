package main

import "weather-scraping-line-bot/server"

// server "weather-scraping-line-bot/server"

func main() {
	// weatherReport, err := weather.FetchWeatherReport()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// areas, timeSeriesInfos := weather.FilterAreas(weatherReport, line.YOKOHAMAWESTAREACODE)
	// weather.ProcessAreaInfos(areas, timeSeriesInfos)
	server.StartServer()
}
