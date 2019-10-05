package handler

import (
	"EscapeAD/api/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewsfeedPostRequest mimic NewsfeedPost
type NewsfeedPostRequest struct {
	Title string `json:"title`
	Post  string `json:"post`
}

// NewsfeedPost Request represents Post feed command.
func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewsfeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
