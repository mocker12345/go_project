package routers

import (
	"github.com/astaxie/beego"
	"myweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	//beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.AutoRouter(&controllers.ReplyController{})
}
