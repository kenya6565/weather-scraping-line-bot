package yokohama

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	model "github.com/kenya6565/weather-scraping-line-bot/app/model"
)

type YokohamaWeatherProcessor struct {
	JmaApiEndpoint string
	AreaCode       string
}

// 構造体を作るメソッドをyokohama package内で定義しておく
// そうすることでweather/factory.go内でyokohama packageをimportせずに構造体の定義ができエラーを回避できる。
func NewYokohamaWeatherProcessor() *YokohamaWeatherProcessor {
	return &YokohamaWeatherProcessor{
		JmaApiEndpoint: "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json",
		AreaCode:       "140020",
	}
}

// FetchDataFromJma makes a GET request to the JMA API, reads the response body, and returns it.
func (y *YokohamaWeatherProcessor) FetchDataFromJMA() ([]model.WeatherInfo, error) {
	resp, err := http.Get(y.JmaApiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// get API data as slice
	var weatherReport []model.WeatherInfo
	err = json.Unmarshal(body, &weatherReport)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return weatherReport, nil
}
