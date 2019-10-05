package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingGet Request represents ping command.
func PingGet() *gin.HandlerFunc {
	func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"hello": "Found me"})
	}
}
