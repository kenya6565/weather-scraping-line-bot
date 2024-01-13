package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	domain "github.com/kenya6565/weather-scraping-line-bot/app/domain/weather"
)

func (c *CityWeatherConfig) FetchDataFromJMA() ([]domain.WeatherInfo, error) {
	resp, err := http.Get(c.JmaApiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// get API data as slice
	var weatherReport []domain.WeatherInfo
	err = json.Unmarshal(body, &weatherReport)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return weatherReport, nil
}
