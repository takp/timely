package helpers

import (
	"strings"
)

func AddBaseURL(urls []string, baseURL string) []string {
	new_urls := []string{}
	for _, url := range urls {
		if !strings.HasPrefix(url, "http") {
			url = baseURL + url
		}
		new_urls = append(new_urls, url)
	}
	return new_urls
}
