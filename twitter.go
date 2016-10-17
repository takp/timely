package main

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"time"
	"github.com/takp/timely/helpers"
	"strconv"
	"github.com/BurntSushi/toml"
)

const (
	TwitterBaseURL = "https://twitter.com"
)

func Twitter(args string) {
	fmt.Println("--- Twitter most shared links from the engineer accounts ---")
	// Get twitter accounts list from config.toml
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	userIDs := conf.TwitterAccounts

	links := make(map[string]int)
	// Fetch each account's tweets
	for _, userID := range userIDs {
		fmt.Println("Fetch Tweets:", userID)
		dateStr := time.Now().AddDate(0, 0, -5).Format("2006-01-02")
		urlStr := []string{TwitterBaseURL, "/search?q=from%3A", userID, "%20since%3A", dateStr}
		url := strings.Join(urlStr, "")
		//fmt.Println("Fetch URL:", url)
		doc, err := goquery.NewDocument(url)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Get link URL for each tweets
		doc.Find(".tweet").Each(func(i int, s *goquery.Selection) {
			link, exists := s.Find(".content .tweet-text a.twitter-timeline-link").Attr("href")
			if exists {
				// Reply tweetの時は除外すべきか？？
				if !strings.HasPrefix(link, "http") {
					link = TwitterBaseURL + link
				}
				links[link] = links[link] + 1
			}
		})
		time.Sleep(500 * time.Millisecond)
	}

	// Sort and display
	ranking := helpers.SortMap(links)
	for i, item := range ranking {
		fmt.Println(i + 1, item.Key, ":",item.Value, "times")
		if i > 8 {
			break
		}
	}

	if args != "" {
		openTwitterSharedLinks(args, ranking)
	}

}

func openTwitterSharedLinks(args string, ranking helpers.PairList) {
	itemNo, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
	}

	if itemNo < 1 || itemNo > 10 {
		fmt.Println("Can not open. The number must be between 1 to 10.")
	}

	item := ranking[itemNo - 1]
	url := item.Key
	fmt.Println("Opening URL:", url)
	helpers.OpenPage(url)
}
