package a

import "github.com/kenya6565/weather-scraping-line-bot/app/server"

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
