package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	jsonList = make(g.Slice, 10)
)

func init() {
	for i := 0; i < 10; i++ {
		jsonList[i] = g.Map{
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
		// 3. JSON response.
		group.GET("/json", func(r *ghttp.Request) {
			r.Response.WriteJson(jsonList)
		})
	})
	s.SetPort(3000)
	s.Run()
}
