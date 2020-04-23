package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()
	app.Get("/hello", func(ctx iris.Context) {
		ctx.WriteString("hello world")
	})
	app.Get("/query", func(ctx iris.Context) {
		ctx.WriteString(ctx.URLParam("id"))
	})
	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"id":   10000,
			"name": "benchmarks",
		})
	})
	app.Run(iris.Addr(":3000"))
}