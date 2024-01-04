package weather

import (
	model "github.com/kenya6565/weather-scraping-line-bot/app/model"
)

// WeatherProcessor defines the methods required for processing weather information for cities.
type WeatherProcessor interface {
	FetchDataFromJMA() ([]model.WeatherInfo, error)
	TransformWeatherData([]model.WeatherInfo) []model.TimeSeriesInfo
	GetCityName() string
	GetJmaApiEndpoint() string
	GetAreaCode() string
	GetAreaName() string
}
