package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

type MemoLog struct {
	gin.LogFormatterParams
	operator string
}

func TaskMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		op := c.Query("operator")
		log.Printf("operator :|%s | %s |%s", op, c.Request.Method, c.Request.URL)
	}
}
