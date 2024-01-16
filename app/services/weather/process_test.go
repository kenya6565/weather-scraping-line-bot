package weather

import (
	"testing"

	"github.com/kenya6565/weather-scraping-line-bot/app/domain/weather"
	"github.com/stretchr/testify/assert"
)

func createSampleWeatherInfo() []domain.WeatherInfo {
	return []domain.WeatherInfo{
		{
			TimeSeries: []domain.TimeSeriesInfo{
				{
					Areas: []domain.AreaInfo{
						{
							Area: struct {
								Name string `json:"name"`
								Code string `json:"code"`
							}{
								Name: "西部",
								Code: "140020",
							},
							Pops: []string{"10", "20", "30", "40", "50"},
						},
					},
					TimeDefines: []string{"2023-07-27T00:00:00+09:00", "2023-07-27T06:00:00+09:00", "2023-07-27T12:00:00+09:00", "2023-07-27T18:00:00+09:00", "2023-07-28T00:00:00+09:00"},
				},
			},
		},
	}
}

func TestTransformWeatherData(t *testing.T) {
	weatherInfo := createSampleWeatherInfo()
	cwc := CityWeatherConfig{
		AreaCode: "140020",
		AreaName: "西部",
	}

	// 正常ケース
	result := cwc.TransformWeatherData(weatherInfo)
	assert.NotEmpty(t, result, "結果は空であってはならない")
	assert.Equal(t, 4, len(result[0].TimeDefines), "先頭のTimeDefinesが除去されていること")
	assert.Equal(t, 4, len(result[0].Areas[0].Pops), "先頭のPopsが除去されていること")

	// 異常ケース
	weatherInfo[0].TimeSeries[0].TimeDefines = weatherInfo[0].TimeSeries[0].TimeDefines[:4] // TimeDefinesの要素を4つに減らす
	result = cwc.TransformWeatherData(weatherInfo)
	assert.Empty(t, result, "TimeDefinesの要素が5でない場合、結果は空でなければならない")
}
