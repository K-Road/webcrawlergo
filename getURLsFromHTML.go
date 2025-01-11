package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseUrl string) ([]string, error) {

	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base url: %v", err)
	}

	var urls []string

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return urls, fmt.Errorf("couldn't parse html: %v", err)
	}

	anchors := findAnchorTags(doc)

	for _, anchor := range anchors {
		url, err := url.Parse(anchor)
		if err != nil {
			return urls, err
		}

		resolvedURL := baseUrl.ResolveReference(url)
		finalurl := strings.TrimSuffix(resolvedURL.String(), "/")
		urls = append(urls, finalurl)
	}
	return urls, nil
}

func findAnchorTags(node *html.Node) []string {
	var anchors []string

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				anchors = append(anchors, attr.Val)
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		anchors = append(anchors, findAnchorTags(child)...)
	}

	return anchors
}
