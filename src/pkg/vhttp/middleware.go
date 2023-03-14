package vhttp

import (
	"github.com/gin-gonic/gin"
)

func WrapHandler(f func()) gin.HandlerFunc {
	return func(c *gin.Context) {
		f()
		c.Next()
	}
}
