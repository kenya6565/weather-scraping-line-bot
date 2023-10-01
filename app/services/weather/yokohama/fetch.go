package yokohama

import (
	"encoding/json"
	"io"
	"net/http"

)

type YokohamaWeatherProcessor struct {
	JmaApiEndpoint string
	AreaCode       string
}

// FetchDataFromJma makes a GET request to the JMA API, reads the response body, and returns it.
func (y *YokohamaWeatherProcessor) FetchDataFromJMA() ([]WeatherInfo, error) {

	resp, err := http.Get(y.JmaApiEndpoint)
	if err != nil {
		return nil, err
	}
	// Close the response body once all operations on it are done.
	// This is essential to release resources and avoid potential memory leaks.
	defer resp.Body.Close()
	// Read the entire response body and return its contents.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherReport []WeatherInfo
	err = json.Unmarshal(body, &weatherReport)
	return weatherReport, err
}
