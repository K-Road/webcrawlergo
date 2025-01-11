package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	sortedPages := sortPages(pages, "desc")
	for _, page := range sortedPages {
		url := page.URL
		count := page.Count
		fmt.Printf("Found %d internal links to %s\n", count, url)
	}
}

type Page struct {
	URL   string
	Count int
}

func sortPages(pages map[string]int, order string) []Page {
	pageSlice := []Page{}
	for url, count := range pages {
		pageSlice = append(pageSlice, Page{URL: url, Count: count})
	}

	if order == "desc" {
		sort.Slice(pageSlice, func(i, j int) bool {
			if pageSlice[i].Count == pageSlice[j].Count {
				return pageSlice[i].URL < pageSlice[j].URL
			}
			return pageSlice[i].Count > pageSlice[j].Count
		})
	} else {
		sort.Slice(pageSlice, func(i, j int) bool {
			if pageSlice[i].Count == pageSlice[j].Count {
				return pageSlice[i].URL > pageSlice[j].URL
			}
			return pageSlice[i].Count < pageSlice[j].Count
		})
	}

	return pageSlice
}
