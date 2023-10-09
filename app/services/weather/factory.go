package weather

import (
	"errors"

	"github.com/kenya6565/weather-scraping-line-bot/app/services/weather/yokohama"
)

// 構造体の定義はfactory.goでは行わない。
func GetWeatherProcessorForCity(city string) (WeatherProcessor, error) {
	switch city {
	case "yokohama":
		// Create a new instance of YokohamaWeatherProcessor
		return yokohama.NewYokohamaWeatherProcessor(), nil
	// 他の都市も必要に応じて追加する
	default:
		return nil, errors.New("unknown city")
	}
}
