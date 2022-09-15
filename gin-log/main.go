package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(MyLogger())
	r.Use(gin.Recovery())

	r.Any("/", func(c *gin.Context) {
		a := c.PostForm("aaa")
		c.JSON(200, gin.H{"data": a})
	})

	_ = r.Run()
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path

		body := "[Body]\r\n"

		c.Next()

		if c.Request.URL.RawQuery != "" {
			path += "?" + c.Request.URL.RawQuery
		}

		log.Println(path)

		if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
			log.Println("ParseMultipartForm err:", err)
		}

		for k, v := range c.Request.PostForm {
			body += k + ": " + fmt.Sprint(v) + "\r\n"
		}

		log.Println(body)
	}
}
