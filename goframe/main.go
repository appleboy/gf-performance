package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	templateList = make(g.Slice, 10)
)

func init() {
	for i := 0; i < 10; i++ {
		templateList[i] = g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name-%d`, i),
		}
	}
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		// 1. Plaintext
		group.GET("/hello", func(r *ghttp.Request) {
			r.Response.Write("hello world")
		})
		// 2. Query parameter retrieving.
		group.GET("/query", func(r *ghttp.Request) {
			r.Response.Write(r.GetQuery("id"))
		})
		// 3. Template parsing.
		group.GET("/template", func(r *ghttp.Request) {
			r.Response.WriteTpl("template.html", g.Map{
				"list": templateList,
			})
		})
	})
	s.SetPort(3000)
	s.Run()
}
