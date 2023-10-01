package weather

import (
	"errors"

	"github.com/kenya6565/weather-scraping-line-bot/app/services/weather/yokohama"
)

func GetWeatherProcessorForCity(city string) (WeatherProcessor, error) {
	switch city {
	case "yokohama":
		return &yokohama.YokohamaWeatherProcessor{
			JmaApiEndpoint: "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json",
			AreaCode:       "140020",
		}, nil
	// 他の都市も必要に応じて追加する
	default:
		return nil, errors.New("unknown city")
	}
}
