package routers

import (
	"github.com/astaxie/beego"
	"myweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
