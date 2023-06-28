package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(
		"your Channel Secret",
		"your Channel Access Token",
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", callbackHandler)

	// 非同期にスクレイピングを開始
	go scrape()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Callback処理はここに書く
}

func scrape() {
	c := colly.NewCollector()

	c.OnHTML("特定のHTML要素", func(e *colly.HTMLElement) {
		// ここで降水確率をパースする処理を書く
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// 神奈川県の天気予報ページ
	err := c.Visit("https://weathernews.jp/onebox/tenki/kanagawa/")
	if err != nil {
		fmt.Println(err)
	}

	// 東京都の天気予報ページ
	err = c.Visit("https://weathernews.jp/onebox/tenki/tokyo/")
	if err != nil {
		fmt.Println(err)
	}
}
