package weather

import (
	"testing"

	domain "github.com/kenya6565/weather-scraping-line-bot/app/domain/weather"
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
	result, err := cwc.TransformWeatherData(weatherInfo)
	assert.NoError(t, err, "エラーが発生しないこと")
	assert.NotEmpty(t, result, "結果は空であってはならない")
	assert.Equal(t, 4, len(result[0].TimeDefines), "先頭のTimeDefinesが除去されていること")
	assert.Equal(t, 4, len(result[0].Areas[0].Pops), "先頭のPopsが除去されていること")

	// 異常ケース
	weatherInfo[0].TimeSeries[0].TimeDefines = weatherInfo[0].TimeSeries[0].TimeDefines[:4] // TimeDefinesの要素を4つに減らす
	result, _ = cwc.TransformWeatherData(weatherInfo)
	assert.Empty(t, result, "結果は空でなければならない")

	// 異常ケース: 対応するエリアが見つからない
	weatherInfo = createSampleWeatherInfo()
	cwc = CityWeatherConfig{
		AreaCode: "999999",
		AreaName: "存在しないエリア",
	}
	result, err = cwc.TransformWeatherData(weatherInfo)
	assert.Error(t, err, "対応するエリアが見つからない場合、エラーが発生する")
	assert.Nil(t, result, "結果はnilであるべき") // "ni" を "nil" に修正
}
