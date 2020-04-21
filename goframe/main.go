package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

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
		// 3. JSON response.
		group.GET("/json", func(r *ghttp.Request) {
			r.Response.WriteJson(g.Map{
				"id":   10000,
				"name": "benchmarks",
			})
		})
	})
	s.SetPort(3000)
	s.Run()
}
