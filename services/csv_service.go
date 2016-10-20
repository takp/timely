package services

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(filename string) ([]string, error) {
	f, err := os.Open("csv/" + filename + ".csv")
	if err != nil {
		return nil, err
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	//fmt.Println(lines)

	urls := []string{}
	for _, line := range lines {
		urls = append(urls, line[0])
	}

	return urls, nil
}

func WriteCSV(urls []string, filename string) {
	csvFile := "csv/" + filename + ".csv"
	// Delete file at first
	os.Remove(csvFile)

	// Create file
	file, err := os.Create(csvFile)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	// Prepare data for writing
	urls_data := prepare_urls_data(urls)

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

func prepare_urls_data(urls []string) [][]string {
	urls_data := [][]string{}
	for _, url := range urls {
		urls_data = append(urls_data, []string{url})
	}
	return urls_data
}
