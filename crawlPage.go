package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedCurrentURL, err := normalizeURL(currentURL.String())
	if err != nil {
		return
	}

	if _, exists := pages[normalizedCurrentURL]; exists {
		pages[normalizedCurrentURL]++
		return
	}

	pages[normalizedCurrentURL] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("htmlBody err: %v\n", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("getURLsFromHTML err: %v\n", err)
		return
	}

	for _, nextURL := range nextURLs {
		crawlPage(rawBaseURL, nextURL, pages)
	}

}
