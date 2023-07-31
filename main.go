package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TimeSeries struct {
	TimeDefines []string `json:"timeDefines"`
	Areas       []Area   `json:"areas"`
}

type Area struct {
	Name     map[string]string `json:"area"`
	Weathers []string          `json:"weathers,omitempty"`
	Winds    []string          `json:"winds,omitempty"`
	Waves    []string          `json:"waves,omitempty"`
	Pops     []string          `json:"pops,omitempty"`
	Temps    []string          `json:"temps,omitempty"`
}

type WeatherReport struct {
	PublishingOffice string       `json:"publishingOffice"`
	ReportDatetime   string       `json:"reportDatetime"`
	TimeSeries       []TimeSeries `json:"timeSeries"`
}

func main() {
	resp, err := http.Get("https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var reports []WeatherReport
	if err := json.Unmarshal(body, &reports); err != nil {
		panic(err)
	}

	// 横浜のデータを探します
	for _, report := range reports {
		for _, ts := range report.TimeSeries {
			for _, area := range ts.Areas {
				if area.Name["name"] == "横浜" {
					// データを表示します
					fmt.Printf("Weathers: %v\n", area.Weathers)
					fmt.Printf("Winds: %v\n", area.Winds)
					fmt.Printf("Waves: %v\n", area.Waves)
					fmt.Printf("Pops: %v\n", area.Pops)
					fmt.Printf("Temps: %v\n", area.Temps)
				}
			}
		}
	}
}
