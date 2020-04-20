package main

import (
	"fmt"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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
	e := echo.New()
	c, err := ioutil.ReadFile("template.html")
	if err != nil {
		panic(err)
	}
	t := &Template{
		templates: template.Must(template.New("template").Parse(string(c))),
	}
	e.Renderer = t

	// 1. Plaintext
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	// 2. Query parameter retrieving.
	e.GET("/query", func(c echo.Context) error {
		return c.String(http.StatusOK, c.QueryParam("id"))
	})

	// 3. Template parsing.
	e.GET("/template", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template", map[string]interface{}{
			"list": templateList,
		})
	})

	e.Start(":3000")
}
