package weather

import (
	weather "github.com/kenya6565/weather-scraping-line-bot/app/models/weather"
)

// WeatherProcessor defines the methods required for processing weather information for cities.
type WeatherProcessor interface {
	FetchDataFromJMA() ([]weather.WeatherInfo, error)
	FilterAreas([]weather.WeatherInfo) ([]weather.AreaInfo, []weather.TimeSeriesInfo)
	ProcessAreaInfos([]weather.AreaInfo, []weather.TimeSeriesInfo) []string
}
