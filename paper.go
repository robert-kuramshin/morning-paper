package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://w1.weather.gov/xml/current_obs/KATL.rss")//"http://feeds.bbci.co.uk/news/world/rss.xml")//https://hackaday.com/blog/feed/
	fmt.Println(feed.Title)
	for _,item := range feed.Items {
		fmt.Println(item.Title)
		if(item.Image != nil){
			fmt.Println(item.Image.URL)
		}
	}
}
