package main

import (
	"fmt"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/takp/timely/helpers"
)

const (
	GithubBaseURL = "https://github.com"
	GithubTrendingURL = "https://github.com/trending"
)

func Github(args string) {
	if args == "" {
		fmt.Println("--- Github Trending Repositories ---")
	} else {
		fmt.Println(OpeningMessage)
	}

	doc, err := goquery.NewDocument(GithubTrendingURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	urls := []string{}

	doc.Find(".repo-list-item").Each(func(i int, s *goquery.Selection){
		//name := s.Find("h3.repo-list-name").Text()
		url, _ := s.Find(".repo-list-name a").Attr("href")
		desc := s.Find(".repo-list-description").Text()
		name := url
		desc = strings.TrimSpace(desc)
		urls = append(urls, url)
		if args == "" {
			fmt.Println(i + 1, name, ":", desc)
		}
	})

	if args != "" {
		openGithubPage(args, urls)
	}
}

func openGithubPage(args string, urls []string) {
	itemNo, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
	}

	if itemNo < 1 || itemNo > 25 {
		fmt.Println("Can not open. The number must be between 1 to 25.")
	}

	url := GithubBaseURL + urls[itemNo - 1]
	fmt.Println("Item:", itemNo, "URL:", url)
	helpers.OpenPage(url)
}
