package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherInfo struct {
	TimeSeries []TimeSeriesInfo `json:"timeSeries"`
}

type TimeSeriesInfo struct {
	Areas       []AreaInfo `json:"areas"`
	TimeDefines []string   `json:"timeDefines"`
}

type AreaInfo struct {
	Area struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"area"`
	// can be nil
	Pops *[]string `json:"pops"`
}

func main() {
	resp, err := http.Get("https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var weatherReport []WeatherInfo
	if err := json.Unmarshal(body, &weatherReport); err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range weatherReport {
		for _, timeSeries := range info.TimeSeries {
			for _, area := range timeSeries.Areas {
				if area.Area.Code == "140020" && area.Pops != nil {
					fmt.Println("Area: ", area.Area.Name)
					fmt.Println("TimeDefines: ", timeSeries.TimeDefines)
					fmt.Println("Precipitation Probability: ", *area.Pops)
				}
			}
		}
	}
}
