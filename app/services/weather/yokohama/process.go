package yokohama

import (
	weather "github.com/kenya6565/weather-scraping-line-bot/app/models/weather"
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

func (y *YokohamaWeatherProcessor) FilterAreas(weatherReport weather.WeatherInfo) ([]weather.AreaInfoInterface, []weather.TimeSeriesInfo) {
	var areas []weather.AreaInfoInterface
	var timeSeriesInfos []weather.TimeSeriesInfo

	for _, timeSeries := range weatherReport.TimeSeries {
		for _, area := range timeSeries.Areas {
			if area.GetCode() == y.AreaCode && area.GetPops() != nil {
				areas = append(areas, area)
				timeSeriesInfos = append(timeSeriesInfos, timeSeries)
			}
		}
	}

	return areas, timeSeriesInfos
}

func (y *YokohamaWeatherProcessor) ProcessAreaInfos(areas []weather.AreaInfoInterface, timeSeriesInfos []weather.TimeSeriesInfo) []string {
	var messages []string
	for i, area := range areas {
		messages = append(messages, GeneratePrecipProbMessage(area, timeSeriesInfos[i])...)
	}
	return messages
}
