package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	//check for valid URL
	url, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	//check for queries
	var normalizedURL string
	if url.RawQuery != "" {
		normalizedURL = fmt.Sprintf("%s%s?%s", url.Hostname(), url.Path, url.RawQuery)
	} else {
		normalizedURL = fmt.Sprintf("%s%s", url.Hostname(), url.Path)
	}

	//formatting
	normalizedURL = strings.ToLower(normalizedURL)
	normalizedURL = strings.TrimSuffix(normalizedURL, "/")

	return normalizedURL, nil

}
