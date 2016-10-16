package main

import (
	"fmt"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const (
	GithubTrendingURL = "https://github.com/trending"
)

func Github(args string) {
	fmt.Println("--- Github Trending Repositories ---")
	doc, err := goquery.NewDocument(GithubTrendingURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	urls := []string{}

	doc.Find(".repo-list-item").Each(func(i int, s *goquery.Selection){
		name := s.Find("h3.repo-list-name").Text()
		url, _ := s.Find(".repo-list-name a").Attr("href")
		desc := s.Find(".repo-list-description").Text()

		name = strings.Trim(name, "\n")
		name = strings.Trim(name, "\n ")
		name = strings.TrimSpace(name)
		desc = strings.TrimSpace(desc)
		urls = append(urls, url)
		fmt.Println(i + 1, name, ":", desc)
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

	fmt.Println(itemNo)
	if itemNo < 1 || itemNo > 25 {
		fmt.Println("Can not open. The number must be between 1 to 25.")
	}

	url := "http://github.com" + urls[itemNo-1]
	fmt.Println("Opening URL:", url)
	openBrowser(url)
}