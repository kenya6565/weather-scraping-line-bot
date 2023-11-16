package weather

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

// FilterAreas filters the provided WeatherInfo for areas matching the CityWeatherConfig's AreaCode.
func (y *CityWeatherConfig) FilterAreas(weatherReport []model.WeatherInfo) ([]model.AreaInfo, []model.TimeSeriesInfo) {
	var matchedAreas []model.AreaInfo
	var matchedTimeSeriesInfos []model.TimeSeriesInfo

	for _, report := range weatherReport {
		for _, series := range report.TimeSeries {
			for _, area := range series.Areas {
				if area.Area.Code == y.AreaCode && area.Pops != nil && len(*area.Pops) > 0 {
					matchedAreas = append(matchedAreas, area)
					matchedTimeSeriesInfos = append(matchedTimeSeriesInfos, series)
				}
			}
		}
	}
	return matchedAreas, matchedTimeSeriesInfos
}
// APIから取得したデータをAreaInfoとTimeSeriesInfoのデータに分ける
	// TODO: 普通にfactory.goで作った構造体のフィールドと比較すればいいだけじゃね？
	// APIから取得したローデータから必要な情報のみだけを抽出するメソッドにしたい
	// popsだけ抽出する
	// areaNameとareaCodeに一致している。かつareasオブジェクトの中にpopsというオブジェクトがある
	// output.goはそのデータをただ利用するだけ
