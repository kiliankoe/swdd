package speiseplan

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

// FeedURL is the URL of the Menu RSS Feed
const FeedURL = "http://www.studentenwerk-dresden.de/feeds/speiseplan.rss"

// GetCurrent returns current menu data
func GetCurrent() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(FeedURL)

	for _, meal := range feed.Items {
		fmt.Println(meal.Title)
	}
}
