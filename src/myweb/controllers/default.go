package controllers

import (
	"github.com/astaxie/beego"
	"myweb/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkUser(c.Ctx)
	topics, err := models.GetAllTopic(true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
}
