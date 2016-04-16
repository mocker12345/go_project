package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkUser(c.Ctx)
	fmt.Println(c.Data["IsLogin"])
}
