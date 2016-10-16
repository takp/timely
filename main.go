package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Run all...")
		getAll()
		return
	}
	subcommand := os.Args[1]
	fmt.Println("Subcommand: ", subcommand)

	switch subcommand {
	case "q": qiita()
	case "qiita": qiita()
	case "g": github()
	case "github": github()
	}
}

func getAll() {
	qiita()
	github()
}

func qiita() {
	fmt.Println("Get qiita!")
}

func github() {
	fmt.Println("Get github!")
}
