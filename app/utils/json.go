package utils

import (
	"encoding/json"
	"fmt"

	weather "github.com/kenya6565/weather-scraping-line-bot/app/models/weather"
	yokohama "github.com/kenya6565/weather-scraping-line-bot/app/models/weather/yokohama"
)

func UnmarshalJSON(b []byte, weatherReports *[]weather.WeatherInfo) error {
	var rawWeather struct {
		TimeSeries []struct {
			AreasData   []json.RawMessage `json:"areas"`
			TimeDefines []string          `json:"timeDefines"`
		} `json:"timeSeries"`
	}

	if err := json.Unmarshal(b, &rawWeather); err != nil {
		return err
	}

	for _, rawTS := range rawWeather.TimeSeries {
		var tsi weather.TimeSeriesInfo
		for _, rawArea := range rawTS.AreasData {
			var areaMap map[string]interface{}
			if err := json.Unmarshal(rawArea, &areaMap); err != nil {
				return err
			}
			areaCode, ok := areaMap["area"].(map[string]interface{})["code"].(string)
			if !ok {
				return fmt.Errorf("failed to assert area code as string")
			}

			switch areaCode {
			case "140020": // Yokohama
				var yokohamaArea yokohama.YokohamaAreaInfo
				if err := json.Unmarshal(rawArea, &yokohamaArea); err != nil {
					return err
				}
				tsi.Areas = append(tsi.Areas, &yokohamaArea)

			// TODO: Add other cities as needed
			default:
				return fmt.Errorf("unsupported area code: %s", areaCode)
			}
		}
		tsi.TimeDefines = rawTS.TimeDefines

		// Create a new WeatherInfo and append the tsi to its TimeSeries field
		wi := weather.WeatherInfo{
			TimeSeries: []weather.TimeSeriesInfo{tsi},
		}
		*weatherReports = append(*weatherReports, wi)
	}

	return nil
}
