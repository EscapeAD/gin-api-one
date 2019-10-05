package main

import (
	"EscapeAD/api/httpd/handler"
	"EscapeAD/api/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	feed := newsfeed.New()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run()

}
