package main

import (
	"fmt"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"github.com/takp/timely/helpers"
)

const (
	QiitaBaseURL = "http://qiita.com"
	QiitaPopularURL = "http://qiita.com/popular-items"
)

func Qiita(args string) {
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

	if args != "" {
		openQiitaPage(args, urls)
	}
}

func openQiitaPage(args string, urls []string) {
	itemNo, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(itemNo)
	if itemNo < 1 || itemNo > 20 {
		fmt.Println("Can not open. The number must be between 1 to 20.")
	}

	url := QiitaBaseURL + urls[itemNo - 1]
	fmt.Println("Open URL:", url)
	helpers.OpenPage(url)
}
