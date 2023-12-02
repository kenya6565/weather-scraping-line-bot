package weather

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchDataFromJMA(t *testing.T) {
	// 気象庁のAPIと同じレスポンスを返すテストサーバーを準備
	// 事前にAPIのレスポンスを準備しておく方法もあるが、それだとAPIからデータを取得するプロセスをテストすることができない
	// 今回はAPIを叩いて適切にfetchできているか自体のテストを行うのでテストサーバーを用意すると完全にテストできる
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 指定したJSON文字列をHTTPレスポンスとして書き出す
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

	// エラーケース1: 無効なURL
	c.JmaApiEndpoint = "invalid_url"
	_, err = c.FetchDataFromJMA()
	assert.Error(t, err)

	// エラーケース2: レスポンスボディが読み取れない
	ts.Close() // テストサーバーを閉じてレスポンスボディが読み取れない状態を作る
	_, err = c.FetchDataFromJMA()
	assert.Error(t, err)

	// エラーケース3: JSONのアンマーシャルができない
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `invalid_json`) // 無効なJSONをレスポンスとして返す
	}))
	defer ts.Close()
	c.JmaApiEndpoint = ts.URL
	_, err = c.FetchDataFromJMA()
	assert.Error(t, err)

	// エラーケース4: レスポンスボディの読み取りエラー
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1e6（100万）個の "a" を連結した大きな文字列をレスポンスとして返す
		// これでレスポンスボディの読み取り時にエラーを発生させてerrorを返しているかの挙動を確認する
		var bigData string
		for i := 0; i < 1e6; i++ {
			bigData += "a"
		}
		fmt.Fprintln(w, bigData)
	}))
	defer ts.Close()
	c.JmaApiEndpoint = ts.URL
	_, err = c.FetchDataFromJMA()
	assert.Error(t, err)
}
