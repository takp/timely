package main

import (
	"fmt"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/takp/timely/helpers"
	"github.com/takp/timely/services"
)

const (
	GithubBaseURL = "https://github.com"
	GithubTrendingURL = "https://github.com/trending"
	GithubFileName = "github"
)

func Github(args string) {
	urls := []string{}
	if args != "" {
		fmt.Println(OpeningMessage)
		urls, _ = services.ReadCSV(GithubFileName)
		if len(urls) == 0 {
			urls = fetchGithub()
		}
		services.WriteCSV(urls, GithubFileName)
		openGithubPage(args, urls)
	} else {
		fmt.Println("--- Github Trending Repositories ---")
		urls = fetchGithub()
		services.WriteCSV(urls, GithubFileName)
	}
}

func fetchGithub() []string {
	urls := []string{}
	doc, err := goquery.NewDocument(GithubTrendingURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	doc.Find(".repo-list-item").Each(func(i int, s *goquery.Selection){
		//name := s.Find("h3.repo-list-name").Text()
		url, _ := s.Find(".repo-list-name a").Attr("href")
		desc := s.Find(".repo-list-description").Text()
		name := url
		desc = strings.TrimSpace(desc)
		urls = append(urls, url)
		fmt.Println(i + 1, name, ":", desc)
	})
	urls = helpers.AddBaseURL(urls, GithubBaseURL)
	return urls
}

func openGithubPage(args string, urls []string) {
	itemNo, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
	}

	if itemNo < 1 || itemNo > 25 {
		fmt.Println("Can not open. The number must be between 1 to 25.")
		return
	}

	url := urls[itemNo - 1]
	fmt.Println("Item:", itemNo, "URL:", url)
	helpers.OpenPage(url)
}
