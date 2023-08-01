package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
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
					for i, pop := range *area.Pops {
						popInt, err := strconv.Atoi(pop)
						if err != nil {
							fmt.Println("Error converting pop to integer: ", err)
							continue
						}
						if popInt >= 50 {
							parsedTime, _ := time.Parse(time.RFC3339, timeSeries.TimeDefines[i])
							jst := time.FixedZone("Asia/Tokyo", 9*60*60)
							fmt.Printf("Time: %v, Precipitation Probability: %v\n", parsedTime.In(jst).Format("2006-01-02 15:04"), pop)
						}
					}
				}
			}
		}
	}
}
