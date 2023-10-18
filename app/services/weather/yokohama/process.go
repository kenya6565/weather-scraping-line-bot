package yokohama

import (
	model "github.com/kenya6565/weather-scraping-line-bot/app/model"
)

//  ### HERE IS JSON RESPONSE FROM API
// {
// 	"timeSeries": [
// 		# 1
// 		{
// 			"timeDefines": [
// 				"2023-07-27T18:00:00+09:00",
// 				"2023-07-28T00:00:00+09:00",
// 				"2023-07-28T06:00:00+09:00",
// 				"2023-07-28T12:00:00+09:00",
// 				"2023-07-28T18:00:00+09:00"
// 			],
// 			# 2
// 			"areas": [
// 				{
// 					# 3
// 					"area": {
// 						"name": "西部",
// 						# 4
// 						"code": "140020"
// 					},
// 					# 3
// 					"pops": [
// 						"30",
// 						"10",
// 						"10",
// 						"20",
// 						"20"
// 					]
// 				}
// 			]
// 		}
// 	]
// }

func (y *YokohamaWeatherProcessor) FilterAreas(weatherReport []model.WeatherInfo) ([]model.AreaInfo, []model.TimeSeriesInfo) {
	var areas []model.AreaInfo
	var timeSeriesInfos []model.TimeSeriesInfo

	for _, info := range weatherReport {
		for _, timeSeries := range info.TimeSeries {
			for _, area := range timeSeries.Areas {
				if area.Area.Code == y.AreaCode && area.Pops != nil {
					areas = append(areas, area)
					timeSeriesInfos = append(timeSeriesInfos, timeSeries)
				}
			}
		}
	}

	return areas, timeSeriesInfos
}

func (y *YokohamaWeatherProcessor) ProcessAreaInfos(areas []model.AreaInfo, timeSeriesInfos []model.TimeSeriesInfo) []string {
	var messages []string
	for i, area := range areas {
		messages = append(messages, GeneratePrecipProbMessage(area, timeSeriesInfos[i])...)
	}
	return messages
}
