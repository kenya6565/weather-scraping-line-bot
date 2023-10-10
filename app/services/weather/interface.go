package weather

import (
	weather "github.com/kenya6565/weather-scraping-line-bot/app/models/weather"
)

// WeatherProcessor defines the methods required for processing weather information for cities.
type WeatherProcessor interface {
	FetchDataFromJMA() ([]weather.WeatherInfo, error)
	FilterAreas([]weather.WeatherInfo) ([]weather.AreaInfoInterface, []weather.TimeSeriesInfo)
	ProcessAreaInfos([]weather.AreaInfoInterface, []weather.TimeSeriesInfo) []string
}
