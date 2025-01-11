package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) { //}, rawCurrentURL string, pages map[string]int) {
	cfg.concurrentyControl <- struct{}{}
	defer func() {
		<-cfg.concurrentyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedCurrentURL, err := normalizeURL(currentURL.String())
	if err != nil {
		return
	}

	isFirst := cfg.addPageVisits(normalizedCurrentURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("htmlBody err: %v\n", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("getURLsFromHTML err: %v\n", err)
		return
	}

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}

}
