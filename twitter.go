package main

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
	//"github.com/takp/timely/helpers"
	"time"
)

const (
	TwitterBaseURL = "https://twitter.com"
	example = "https://twitter.com/search?q=from%3Anaoya_ito%20since%3A2016-10-14"
)

func Twitter(args string) {
	fmt.Println("--- Twitter most shared links ---")
	// User List
	userIDs := []string{"mizuno_takaaki", "yukihiro_matz", "naoya_ito",
		"takoratta", "masuidrive", "mizchi", "Jxck_", "t_wada", "miyagawa"}

	links := make(map[string]int)

	// Fetch each account's tweets
	for _, userID := range userIDs {
		fmt.Println("Fetch Twitter ID:", userID)

		dateStr := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
		urlStr := []string{TwitterBaseURL, "/search?q=from%3A", userID, "%20since%3A", dateStr}
		url := strings.Join(urlStr, "")

		fmt.Println("Fetch URL:", url)
		doc, err := goquery.NewDocument(url)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Get link URL for each tweets
		doc.Find(".tweet").Each(func(i int, s *goquery.Selection) {
			link, exists := s.Find(".content .tweet-text a").Attr("href")
			if exists {
				if !strings.HasPrefix(link, "http") {
					link = TwitterBaseURL + link
				}
				links[link] = links[link] + 1
			}
		})
		time.Sleep(500 * time.Millisecond)
	}

	// Sortして表示
	for k, v := range links {
		fmt.Println(v, "times:", k)
	}

}
