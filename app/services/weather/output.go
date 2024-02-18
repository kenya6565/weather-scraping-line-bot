package weather

import (
	"fmt"
	"strconv"
	"time"

	domain "github.com/kenya6565/weather-scraping-line-bot/app/domain/weather"
)

const PRECIPITATION_PROBABILITY = 50

// GenerateRainMessages generates messages based on precipitation probabilities.
func GenerateRainMessages(timeSeriesInfos []domain.TimeSeriesInfo) []string {
	var messages []string
	shouldNotify := false

	for _, series := range timeSeriesInfos {
		for _, popStr := range series.Areas[0].Pops {
			pop, err := strconv.Atoi(popStr)
			if err == nil && pop >= PRECIPITATION_PROBABILITY {
				shouldNotify = true
				break
			}
		}
		if shouldNotify {
			break
		}
	}

	// if all pops are less than PRECIPITATION_PROBABILITY
	if !shouldNotify {
		return messages
	}

	for _, series := range timeSeriesInfos {
		for i, popStr := range series.Areas[0].Pops {
			pop, err := strconv.Atoi(popStr)
			if err != nil {
				continue
			}
			startTime, endTime, err := getTimeRange(series.TimeDefines[i])
			if err != nil {
				continue
			}
			message := fmt.Sprintf("%s~%s:降水確率%d%%にゃ🐾", startTime[11:], endTime, pop)
			messages = append(messages, message)
		}
	}

	return messages
}

// APIから取得した時間表記を日本時間に修正する
func getTimeRange(timeDefine string) (string, string, error) {
	parsedTime, err := time.Parse(time.RFC3339, timeDefine)
	if err != nil {
		return "", "", err
	}
	// JSTのタイムゾーンを設定する
	// Setting the timezone to JST
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	// 開始時間を"2006-01-02 15:04"の形式でフォーマットする
	startTime := parsedTime.In(jst).Format("2006-01-02 15:04")
	// 終了時間を"15:04"の形式でフォーマットする
	endTime := parsedTime.In(jst).Add(5*time.Hour + 59*time.Minute).Format("15:04")
	return startTime, endTime, nil
}
