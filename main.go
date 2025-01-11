package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := args[0]
	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	// htmlBody, err := getHTML(rawBaseURL)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Print(htmlBody)

	pages := make(map[string]int)
	crawlPage(rawBaseURL, rawBaseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
