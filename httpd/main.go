package main

import (
	"EscapeAD/api/platform/newsfeed"
	"fmt"
)

func main() {
	// r := gin.Default()

	// r.GET("/ping", handler.PingGet())

	// r.Run()

	feed := newsfeed.New()
	fmt.Println(feed)
	feed.Add(newsfeed.Item{"Hello", "Feed 1"})
	fmt.Println(feed)
}
