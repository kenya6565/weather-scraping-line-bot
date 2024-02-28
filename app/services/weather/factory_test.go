package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeatherProcessorForCity(t *testing.T) {
	testCases := []struct {
		city           string
		expectedError  bool
		expectedConfig *CityWeatherConfig
	}{
		{"横浜", false, &CityWeatherConfig{"横浜", "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json", "140020", "西部"}},
		{"東京", false, &CityWeatherConfig{"東京", "https://www.jma.go.jp/bosai/forecast/data/forecast/130000.json", "130010", "東京地方"}},
		{"大阪", false, &CityWeatherConfig{"大阪", "https://www.jma.go.jp/bosai/forecast/data/forecast/270000.json", "270000", "大阪府"}},
		{"未知の都市", true, nil},
	}

	for _, tc := range testCases {
		processor, err := GetWeatherProcessorForCity(tc.city)
		if tc.expectedError {
			assert.Error(t, err, "未知の都市に対してエラーを返すべき")
		} else {
			assert.NoError(t, err, "エラーを返すべきではない")
			assert.Equal(t, tc.expectedConfig, processor, "返された設定が期待値と一致するべき")
		}
	}
}

func TestCityWeatherConfigMethods(t *testing.T) {
	c := CityWeatherConfig{
		CityName:       "横浜",
		JmaApiEndpoint: "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json",
		AreaCode:       "140020",
		AreaName:       "西部",
	}

	assert.Equal(t, "横浜", c.GetCityName(), "CityNameが正しく取得できること")
	assert.Equal(t, "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json", c.GetJmaApiEndpoint(), "JmaApiEndpointが正しく取得できること")
	assert.Equal(t, "140020", c.GetAreaCode(), "AreaCodeが正しく取得できること")
	assert.Equal(t, "西部", c.GetAreaName(), "AreaNameが正しく取得できること")
}
