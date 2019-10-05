package handler

import (
	"EscapeAD/api/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewsfeedGet Request represents ping command.
func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
