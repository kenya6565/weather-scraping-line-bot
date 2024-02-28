package weather

import (
	domain "github.com/kenya6565/weather-scraping-line-bot/app/domain/weather"
)

type WeatherProcessor interface {
	FetchDataFromJMA() ([]domain.WeatherInfo, error)
	TransformWeatherData([]domain.WeatherInfo) ([]domain.TimeSeriesInfo, error)
	GetCityName() string
	GetJmaApiEndpoint() string
	GetAreaCode() string
	GetAreaName() string
}
