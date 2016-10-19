package services

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSV(urls []string, filename string, baseURL string) {
	// Create file
	file, err := os.Create("csv/" + filename + ".csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	// Prepare data for writing
	urls_data := prepare_urls_data(urls, baseURL)

	// Write CSV file
	w := csv.NewWriter(file)
	for _, url_data := range urls_data {
		if err := w.Write(url_data); err != nil {
			log.Fatalln("error writing urls to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func prepare_urls_data(urls []string, baseURL string) [][]string {
	urls_data := [][]string{
		{"url"},
	}
	for _, url := range urls {
		url = baseURL + url
		urls_data = append(urls_data, []string{url})
	}
	return urls_data
}
