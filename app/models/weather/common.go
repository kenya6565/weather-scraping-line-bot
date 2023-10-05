package models

type WeatherInfo struct {
	TimeSeries []TimeSeriesInfo `json:"timeSeries"`
}

type TimeSeriesInfo struct {
	Areas       []AreaInfoInterface `json:"areas"`
	TimeDefines []string            `json:"timeDefines"`
}

type AreaInfoInterface interface {
	GetCode() string
	GetPops() *[]string
}
