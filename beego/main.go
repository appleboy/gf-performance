package main

import (
	"fmt"
	"github.com/astaxie/beego"
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
	c.Ctx.Output.JSON(jsonList, false, true)
}

func main() {
	beego.Router("/hello", &Controller{}, "GET:Hello")
	beego.Router("/query", &Controller{}, "GET:Query")
	beego.Router("/json", &Controller{}, "GET:Template")
	beego.Run(":3000")
}
