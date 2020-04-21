package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	jsonList = make([]interface{}, 10)
)

func init() {
	for i := 0; i < 10; i++ {
		jsonList[i] = map[string]interface{}{
			"id":   i,
			"name": fmt.Sprintf(`name-%d`, i),
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// 1. Plaintext
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	// 2. Query parameter retrieving.
	r.GET("/query", func(c *gin.Context) {
		c.String(http.StatusOK, c.Query("id"))
	})
	// 3. JSON response.
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, jsonList)
	})
	r.Run(":3000")
}
