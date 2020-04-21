package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":   10000,
			"name": "benchmarks",
		})
	})
	r.Run(":3000")
}
