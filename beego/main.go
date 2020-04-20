package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

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

// 3. Template parsing.
func (c *Controller) Template() {
	c.TplName = "template.html"
	c.Data["list"] = templateList
	c.Render()
}

func main() {
	beego.Router("/hello", &Controller{}, "GET:Hello")
	beego.Router("/query", &Controller{}, "GET:Query")
	beego.Router("/template", &Controller{}, "GET:Template")
	beego.Run(":3000")
}
