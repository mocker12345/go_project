package controllers

import (
	"github.com/astaxie/beego"
	"myweb/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
	}
	c.TplName = "category.html"
	c.Data["IsCategory"] = true
	c.Data["IsLogin"] = checkUser(c.Ctx)

	var err error
	c.Data["Categorys"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}

}
