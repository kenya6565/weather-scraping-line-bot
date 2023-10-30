package weather

import (
	"errors"
)

type CityWeatherConfig struct {
	JmaApiEndpoint string
	AreaCode       string
}

func GetWeatherProcessorForCity(city string) (WeatherProcessor, error) {
	switch city {
	case "横浜":
		return &CityWeatherConfig{
			JmaApiEndpoint: "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json",
			AreaCode:       "140020",
		}, nil
	// TODO: 他の都市も必要に応じて追加する
	default:
		return nil, errors.New("unknown city")
	}
}
