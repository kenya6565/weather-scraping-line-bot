package weather

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
	Pops []string `json:"pops"` // nilを許容しないため、ポインタではなくスライスとして定義
}
