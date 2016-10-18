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
	OpeningMessage = "Opening..."
)

func Qiita(args string) {
	if args == "" {
		fmt.Println("--- Qiita 人気の投稿 ---")
	} else {
		fmt.Println(OpeningMessage)
	}

	doc, err := goquery.NewDocument(QiitaPopularURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	urls := []string{}

	doc.Find(".popularItem").Each(func(i int, s *goquery.Selection){
		title := s.Find(".popularItem_articleTitle_text").Text()
		url, _ := s.Find(".popularItem_articleTitle_text").Attr("href")
		urls = append(urls, url)

		if args == "" {
			fmt.Println(i + 1, title)
		}
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

	if itemNo < 1 || itemNo > 20 {
		fmt.Println("Can not open. The number must be between 1 to 20.")
	}

	url := QiitaBaseURL + urls[itemNo - 1]
	fmt.Println("URL:", url)
	helpers.OpenPage(url)
}
