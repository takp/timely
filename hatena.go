package main

import (
	"fmt"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"github.com/takp/timely/helpers"
)

const (
	HatenaTechBlogURL = "http://b.hatena.ne.jp/hotentry/it/%E6%8A%80%E8%A1%93%E3%83%96%E3%83%AD%E3%82%B0"
)

func Hatena(args string) {
	if args == "" {
		fmt.Println("--- はてなブックマーク 技術ブログ ホットエントリー ---")
	} else {
		fmt.Println(OpeningMessage)
	}
	doc, err := goquery.NewDocument(HatenaTechBlogURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	urls := []string{}

	doc.Find(".top li.entry-unit").Each(func(i int, s *goquery.Selection){
		title := s.Find("h3.hb-entry-link-container").Text()
		url, _ := s.Find("h3.hb-entry-link-container a").Attr("href")
		usersNo := s.Find("ul.users li a span").Text()
		urls = append(urls, url)
		if args == "" {
			fmt.Println(i + 1, title, ":", usersNo, "users")
		}
	})

	if args != "" {
		openHatenaPage(args, urls)
	}
}

func openHatenaPage(args string, urls []string) {
	itemNo, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
	}

	if itemNo < 1 || itemNo > 20 {
		fmt.Println("Can not open. The number must be between 1 to 25.")
	}

	url := urls[itemNo - 1]
	fmt.Println("Item:", itemNo, "URL:", url)
	helpers.OpenPage(url)
}
