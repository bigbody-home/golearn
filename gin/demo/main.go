package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Job struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func main() {
	g := gin.Default()
	g.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong"})
	})
	g.GET("/reader", func(c *gin.Context) {
		res, err := http.Get("http://www.amb-chine.fr/chn/zgzfg/zgds/dsjy/")
		if err != nil {
			panic(err)
		}
		reader := res.Body
		defer res.Body.Close()
		ext := make(map[string]string)
		ext["Content-Disposition"] = `attachment; filename="gopher.png"`
		c.DataFromReader(http.StatusOK, res.ContentLength, res.Header.Get("Content-Type"), reader, ext)
	})
	g.POST("/send", func(c *gin.Context) {

		var ct Job
		err := c.ShouldBindWith(&ct, binding.JSON)
		if err == nil {
			fmt.Println(ct)
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": ct})
		} else {
			fmt.Println(ct)
			c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "message": err.Error()})
		}

	})
	g.Run()
}
