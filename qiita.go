package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const (
	QiitaPopularURL = "http://qiita.com/popular-items"
)

func Qiita() {
	fmt.Println("--- Qiita 人気の投稿 ---")
	doc, err := goquery.NewDocument(QiitaPopularURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	titles := []string{}
	urls := []string{}

	doc.Find(".popularItem").Each(func(i int, s *goquery.Selection){
		title := s.Find(".popularItem_articleTitle_text").Text()
		titles = append(titles, title)

		url, _ := s.Find(".popularItem_articleTitle_text").Attr("href")
		urls = append(urls, url)

		fmt.Println(i + 1, " ", title)
	})
	fmt.Println(titles)
	fmt.Println(urls)

}
