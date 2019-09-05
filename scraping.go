package main

import (
	"fmt"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

func main () {
	doc, err := goquery.NewDocument("https://nijisanji.ichikara.co.jp/member/")
	if err != nil {
		fmt.Println(err)
	}
	u := url.URL{}
	u.Scheme = doc.Url.Scheme
	u.Host = doc.Url.Host

	// ライバーリスト取得
	charaList := doc.Find("div.roundcorner")

	// 掲載イベントURL一覧を取得
	charaList.Each(func(i int, s *goquery.Selection) {
		// 名前
		name, _ := s.Find("img").Attr("alt")
		fmt.Println(name)
		// src
		img, _ := s.Find("img").Attr("src")
		fmt.Println(img)
	})
}
