package weather

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchDataFromJMA(t *testing.T) {
	// テストサーバーを作成
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[{"timeSeries": [{"timeDefines": ["2023-07-27T18:00:00+09:00"], "areas": [{"area": {"name": "西部", "code": "140020"}, "pops": ["30"]}] }] }]`)
	}))
	defer ts.Close()

	c := &CityWeatherConfig{
		JmaApiEndpoint: ts.URL,
		AreaCode:       "140020",
		AreaName:       "西部",
	}

	// 正常ケース
	weatherReport, err := c.FetchDataFromJMA()
	assert.NoError(t, err)
	assert.NotNil(t, weatherReport)

	// エラーケース
	c.JmaApiEndpoint = "invalid_url"
	_, err = c.FetchDataFromJMA()
	assert.Error(t, err)
}
