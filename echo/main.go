package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
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
		return c.JSON(http.StatusOK, jsonList)
	})

	e.Start(":3000")
}
