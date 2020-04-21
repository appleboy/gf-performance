package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	// 1. Plaintext
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	// 2. Query parameter retrieving.
	e.GET("/query", func(c echo.Context) error {
		return c.String(http.StatusOK, c.QueryParam("id"))
	})

	// 3. JSON response.
	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id":   10000,
			"name": "benchmarks",
		})
	})

	e.Start(":3000")
}
