package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("# usage ./crawler URL maxConcurrancey maxPages")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := args[0]

	maxConcurrancy, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid concurrancy", err)
		return
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid pages", err)
		return
	}
	cfg, err := configure(rawBaseURL, maxConcurrancy, maxPages)
	if err != nil {
		fmt.Printf("error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
	printReport(cfg.pages, cfg.baseURL.String())
}
