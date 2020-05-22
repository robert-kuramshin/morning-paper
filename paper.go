package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

var line_len int = 28

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	res := ""
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.bbci.co.uk/news/world/rss.xml")//https://hackaday.com/blog/feed/"https://w1.weather.gov/xml/current_obs/KATL.rss")//
	//fmt.Println(feed.Title)
	for _,item := range feed.Items[:3] {
		//fmt.Println(item.Title)
		i := 0
		title := item.Title
		tlen := len(title)
		for(i<tlen) {
			res += title[i:min(i+line_len,tlen)]
			res += "\n"
			i+=line_len
		}
		res += "\n"
	}

	fmt.Println(res)
}
