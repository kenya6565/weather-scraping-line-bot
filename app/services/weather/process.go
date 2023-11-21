package weather

import (
	model "github.com/kenya6565/weather-scraping-line-bot/app/model"
)

//  ### HERE IS JSON RESPONSE FROM API
// popsは0%でも必ず値が入ってくる
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

// process row data of API and return TimeSeriesInfo data
func (y *CityWeatherConfig) TransformWeatherData(weatherReport []model.WeatherInfo) []model.TimeSeriesInfo {
	var matchedTimeSeries []model.TimeSeriesInfo

	for _, report := range weatherReport {
		for _, series := range report.TimeSeries {
			// 5つの要素を持つTimeDefinesのみ取得する
			if len(series.TimeDefines) != 5 {
				continue
			}

			var matchedAreas []model.AreaInfo
			for _, area := range series.Areas {
				// areaCodeとareaNameが構造体と一致しているものを取得
				if area.Area.Code == y.AreaCode && area.Area.Name == y.AreaName {
					matchedAreas = append(matchedAreas, area)
				}
			}

			if len(matchedAreas) > 0 {
				matchedSeries := model.TimeSeriesInfo{
					Areas:       matchedAreas,
					TimeDefines: series.TimeDefines[1:], // 先頭のTimeDefinesは不必要なので除去
				}
				// 先頭のpopsは不必要なので除去
				for i := range matchedSeries.Areas {
					matchedSeries.Areas[i].Pops = matchedSeries.Areas[i].Pops[1:]
				}

				matchedTimeSeries = append(matchedTimeSeries, matchedSeries)
			}
		}
	}
	return matchedTimeSeries
}

// APIから取得したデータをAreaInfoとTimeSeriesInfoのデータに分ける
// TODO: 普通にfactory.goで作った構造体のフィールドと比較すればいいだけじゃね？
// APIから取得したローデータから必要な情報のみだけを抽出するメソッドにしたい
// areaNameとareaCodeに一致している。かつareasオブジェクトの中にpopsというオブジェクトがある
// popsフィールドがなければcontinue
