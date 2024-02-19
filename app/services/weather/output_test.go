package weather

import (
	"testing"

	domain "github.com/kenya6565/weather-scraping-line-bot/app/domain/weather"
	"github.com/stretchr/testify/assert"
)

// create sample TimeSeriesInfo data
func createSampleTimeSeriesInfo() []domain.TimeSeriesInfo {
	return []domain.TimeSeriesInfo{
		{
			Areas: []domain.AreaInfo{
				{
					Pops: []string{"10", "30", "50", "70", "90"},
				},
			},
			TimeDefines: []string{"2023-07-27T00:00:00+09:00", "2023-07-27T06:00:00+09:00", "2023-07-27T12:00:00+09:00", "2023-07-27T18:00:00+09:00", "2023-07-28T00:00:00+09:00"},
		},
	}
}

func TestGenerateRainMessages(t *testing.T) {
	timeSeriesInfos := createSampleTimeSeriesInfo()

	// 正常ケース
	messages := GenerateRainMessages(timeSeriesInfos)
	assert.Equal(t, 5, len(messages), "メッセージの数はTimeSeriesInfoの数と一致する")

	timeSeriesInfos = createSampleTimeSeriesInfo()
	for i := range timeSeriesInfos[0].Areas[0].Pops {
		timeSeriesInfos[0].Areas[0].Pops[i] = "40"
	}
	messages = GenerateRainMessages(timeSeriesInfos)
	assert.Equal(t, 0, len(messages), "全ての降水確率が閾値未満の場合、メッセージは生成されない")

	// 異常ケース
	timeSeriesInfos = createSampleTimeSeriesInfo()
	timeSeriesInfos[0].TimeDefines[0] = "不正な日時"
	messages = GenerateRainMessages(timeSeriesInfos)
	assert.Equal(t, 4, len(messages), "不正な日時フォーマットの場合は対応するメッセージが生成されない")
	assert.Equal(t, 4, len(messages), "不正な日時フォーマットの場合は対応するメッセージが生成されない")

	timeSeriesInfos = createSampleTimeSeriesInfo()
	timeSeriesInfos[0].Areas[0].Pops[0] = "不正な数値"
	messages = GenerateRainMessages(timeSeriesInfos)
	assert.Equal(t, 4, len(messages), "数値に変換できない降水確率が含まれる場合、対応するメッセージが生成されない")
}
