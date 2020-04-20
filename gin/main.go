package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	templateList = make([]interface{}, 10)
)

func init() {
	for i := 0; i < 10; i++ {
		templateList[i] = map[string]interface{}{
			"id":   i,
			"name": fmt.Sprintf(`name-%d`, i),
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.LoadHTMLFiles("template.html")
	// 1. Plaintext
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	// 2. Query parameter retrieving.
	r.GET("/query", func(c *gin.Context) {
		c.String(http.StatusOK, c.Query("id"))
	})
	// 3. Template parsing.
	r.GET("/template", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template.html", gin.H{
			"list": templateList,
		})
	})
	r.Run(":3000")
}
