package model

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
	// can be nil
	// TODO: API見たら0も入ってくるのでnilにならないかもなので検証
	Pops *[]string `json:"pops"`
}