package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("username")
	pwd := this.Input().Get("password")
	autologin := this.Input().Get("autologin") == "on"
	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autologin {
			maxAge = 1<<32 - 1
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}

func checkUser(ctx *context.Context) bool {
	uname := ctx.Input.Cookie("uname")
	// if err != nil {
	// 	return false
	// }
	// uname := ck.Value
	pwd := ctx.Input.Cookie("pwd")
	// if err != nil {
	// 	return false
	// }
	// pwd := ck.Value

	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd

}
