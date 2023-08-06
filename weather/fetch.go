package weather

import (
	"encoding/json"
	"io"
	"net/http"
)

const JmaApiEndpoint = "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json"

func FetchWeatherReport() ([]WeatherInfo, error) {
	resp, err := http.Get(JmaApiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherReport []WeatherInfo
	err = json.Unmarshal(body, &weatherReport)
	return weatherReport, err
}
