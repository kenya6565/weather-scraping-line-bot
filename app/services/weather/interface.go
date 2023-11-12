package weather

import (
	model "github.com/kenya6565/weather-scraping-line-bot/app/model"
)

// WeatherProcessor defines the methods required for processing weather information for cities.
type WeatherProcessor interface {
	FetchDataFromJMA() ([]model.WeatherInfo, error)
	FilterAreas([]model.WeatherInfo) ([]model.AreaInfo, []model.TimeSeriesInfo)
}
