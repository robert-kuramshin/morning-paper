package main

import (
	"strings"

	"github.com/mmcdole/gofeed"
)

var line_len int = 28
var line_brk string = "----------------------------\n"

var feed_urls = []string{"http://feeds.feedburner.com/ElectrekPodcast", "https://hackaday.com/blog/feed", "http://feeds.bbci.co.uk/news/world/rss.xml"}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parse_feed(url string, item_limit int) string {
	res := ""

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)

	if feed == nil {
		return ""
	}

	res += feed.Title + "\n"
	res += line_brk

	n_items := min(len(feed.Items), item_limit)

	for _, item := range feed.Items[:n_items] {
		//fmt.Println(item.Title)
		line := ""
		flushed := true
		words := strings.Fields(item.Title)
		for _, word := range words {
			if len(line)+len(word)+1 > line_len {
				res += line + "\n"
				line = word
				flushed = true
			} else {
				if len(line) != 0 {
					line += " "
				}
				line += word
				flushed = false
			}
		}
		if !flushed && line != "" {
			res += line + "\n"
		}

		res += "\n"
	}

	return res
}
