package weather

import (
	"errors"
)

type CityWeatherConfig struct {
	CityName       string
	JmaApiEndpoint string
	AreaCode       string
	AreaName       string
}

func (c *CityWeatherConfig) GetCityName() string {
	return c.CityName
}

func (c *CityWeatherConfig) GetJmaApiEndpoint() string {
	return c.JmaApiEndpoint
}

func (c *CityWeatherConfig) GetAreaCode() string {
	return c.AreaCode
}

func (c *CityWeatherConfig) GetAreaName() string {
	return c.AreaName
}

func GetWeatherProcessorForCity(city string) (WeatherProcessor, error) {
	switch city {
	case "横浜":
		return &CityWeatherConfig{
			CityName:       "横浜",
			JmaApiEndpoint: "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json",
			AreaCode:       "140020",
			AreaName:       "西部",
		}, nil
	case "東京":
		return &CityWeatherConfig{
			CityName:       "東京",
			JmaApiEndpoint: "https://www.jma.go.jp/bosai/forecast/data/forecast/130000.json",
			AreaCode:       "130010",
			AreaName:       "東京地方",
		}, nil
	case "大阪":
		return &CityWeatherConfig{
			CityName:       "大阪",
			JmaApiEndpoint: "https://www.jma.go.jp/bosai/forecast/data/forecast/270000.json",
			AreaCode:       "270000",
			AreaName:       "大阪府",
		}, nil
		// TODO: 他の都市も必要に応じて追加
	default:
		return nil, errors.New("unknown city")
	}
}
