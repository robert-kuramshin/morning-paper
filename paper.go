package main

import (
	"fmt"
	"strings"
	"github.com/mmcdole/gofeed"
)

var line_len int = 28
var line_brk string = "----------------------------\n"

var feed_urls = []string{"http://feeds.feedburner.com/ElectrekPodcast","https://hackaday.com/blog/feed","http://feeds.bbci.co.uk/news/world/rss.xml"}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func parse_feed(url string, item_limit int) string {
	res := ""

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)

	res += feed.Title + "\n"
	res += line_brk

	n_items := min(len(feed.Items),item_limit)

	for indx,item := range feed.Items[:n_items] {
		//fmt.Println(item.Title)
		line := ""
		words := strings.Fields(item.Title)
		for _,word := range words {
			if len(line) + len(word) + 1 > line_len {
				res += line + "\n"
				line = word
			} else {
				if(len(line) != 0) {
					line += " "
				}
				line += word
			}
		}
		if(indx < n_items - 1) {
			res += "\n"
		}
	}
	
	return res
}

func main() {
	for _,url := range feed_urls {
		fmt.Println(parse_feed(url,3))
	}
}
