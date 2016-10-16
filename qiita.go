package main

import (
	"fmt"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"os/exec"
	"runtime"
)

const (
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

	url := "http://qiita.com" + urls[itemNo-1]
	fmt.Println("Open URL:", url)
	openBrowser(url)
}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		fmt.Println(runtime.GOOS)
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
