package main

import (
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

// 1. Plaintext
func (c *Controller) Hello() {
	c.Ctx.WriteString("hello world")
}

// 2. Query parameter retrieving.
func (c *Controller) Query() {
	c.Ctx.WriteString(c.Ctx.Input.Query("id"))
}

// 3. JSON response.
func (c *Controller) Template() {
	c.Ctx.Output.JSON(map[string]interface{}{
		"id":   10000,
		"name": "benchmarks",
	}, false, true)
}

func main() {
	beego.Router("/hello", &Controller{}, "GET:Hello")
	beego.Router("/query", &Controller{}, "GET:Query")
	beego.Router("/json", &Controller{}, "GET:Template")
	beego.Run(":3000")
}
