package weather

// WeatherProcessor defines the methods required for processing weather information for cities.
type WeatherProcessor interface {
	FetchDataFromJMA() ([]WeatherInfo, error)
	FilterAreas([]WeatherInfo) ([]AreaInfo, []TimeSeriesInfo)
}


// WeatherData represents raw weather data for a city.
type WeatherData struct {
	// Define fields for raw weather data here.
}

// ProcessedWeatherInfo represents processed weather information for a city.
type ProcessedWeatherInfo struct {
	// Define fields for processed weather information here.
}
