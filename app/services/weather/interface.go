package weather

import (
	domain "github.com/kenya6565/weather-scraping-line-bot/app/domain/weather"
)

// WeatherProcessor defines the methods required for processing weather information for cities.
type WeatherProcessor interface {
	FetchDataFromJMA() ([]domain.WeatherInfo, error)
	TransformWeatherData([]domain.WeatherInfo) []domain.TimeSeriesInfo
	GetCityName() string
	GetJmaApiEndpoint() string
	GetAreaCode() string
	GetAreaName() string
}
