package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const YokohamaWestAreaCode = "140020"

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

func fetchWeatherReport() ([]WeatherInfo, error) {
	resp, err := http.Get("https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json")
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

func filterAreas(weatherReport []WeatherInfo, code string) ([]AreaInfo, []TimeSeriesInfo) {
	var areas []AreaInfo
	var timeSeriesInfos []TimeSeriesInfo
	for _, info := range weatherReport {
		for _, timeSeries := range info.TimeSeries {
			for _, area := range timeSeries.Areas {
				if area.Area.Code == code && area.Pops != nil {
					areas = append(areas, area)
					timeSeriesInfos = append(timeSeriesInfos, timeSeries)
				}
			}
		}
	}
	return areas, timeSeriesInfos
}

func processAreaInfos(areas []AreaInfo, timeSeriesInfos []TimeSeriesInfo) {
	for i, area := range areas {
		printPrecipProb(area, timeSeriesInfos[i])
	}
}

func printPrecipProb(area AreaInfo, timeSeries TimeSeriesInfo) {
	if len(*area.Pops) < 2 || len(timeSeries.TimeDefines) < 2 {
		// Skip if there is less than 2 values of arrays in pops and time defines
		return
	}
	for i, popStr := range (*area.Pops)[1:] { // Skip the first pop
		// converting string to int
		pop, err := strconv.Atoi(popStr)
		if err != nil {
			fmt.Println("Error converting pop to integer: ", err)
			return
		}
		if pop >= 50 {
			// Skip the first time define
			timeDefine := timeSeries.TimeDefines[i+1]
			// converting to time.Time type
			parsedTime, err := time.Parse(time.RFC3339, timeDefine)
			if err != nil {
				fmt.Println("Error parsing time: ", err)
				return
			}
			jst := time.FixedZone("Asia/Tokyo", 9*60*60)
			fmt.Printf("Time: %s, Precipitation Probability: %d\n", parsedTime.In(jst).Format("2006-01-02 15:04"), pop)
		}
	}
}

func main() {
	weatherReport, err := fetchWeatherReport()
	if err != nil {
		fmt.Println(err)
		return
	}

	areas, timeSeriesInfos := filterAreas(weatherReport, YokohamaWestAreaCode)
	processAreaInfos(areas, timeSeriesInfos)
}
