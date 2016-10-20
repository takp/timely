package main

import (
	"fmt"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"github.com/takp/timely/helpers"
	"github.com/takp/timely/services"
)

const (
	QiitaBaseURL = "http://qiita.com"
	QiitaPopularURL = "http://qiita.com/popular-items"
	OpeningMessage = "Opening..."
)

func Qiita(args string) {
	urls := []string{}

	// If args are present, Read csv file and go to file.
	// But if the csv data is nil, fetch new urls.
	if args != "" {
		fmt.Println(OpeningMessage)
		urls, _ = services.ReadCSV("qiita")
		if len(urls) == 0 {
			urls = fetchQiita(args)
		}
		services.WriteCSV(urls, "qiita")
		openQiitaPage(args, urls)
	// if args does not exist, fetch new data and save to CSV file.
	} else {
		fmt.Println("--- Qiita 人気の投稿 ---")
		urls = fetchQiita(args)
		services.WriteCSV(urls, "qiita")
	}
}

func fetchQiita(args string) []string {
	urls := []string{}
	doc, err := goquery.NewDocument(QiitaPopularURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	doc.Find(".popularItem").Each(func(i int, s *goquery.Selection){
		title := s.Find(".popularItem_articleTitle_text").Text()
		url, _ := s.Find(".popularItem_articleTitle_text").Attr("href")
		urls = append(urls, url)

		if args == "" {
			fmt.Println(i + 1, title)
		}
	})
	urls = helpers.AddBaseURL(urls, QiitaBaseURL)
	return urls
}

func openQiitaPage(args string, urls []string) {
	itemNo, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
	}

	if itemNo < 1 || itemNo > 20 {
		fmt.Println("Can not open. The number must be between 1 to 20.")
		return
	}

	url := urls[itemNo - 1]
	fmt.Println("URL:", url)
	helpers.OpenPage(url)
}
