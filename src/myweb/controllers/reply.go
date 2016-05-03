package controllers

import (
	"github.com/astaxie/beego"
	"myweb/models"
)

type ReplyController struct {
	beego.Controller
}

func (r *ReplyController) Add() {
	tid := r.Input().Get("tid")
	nikename := r.Input().Get("nikename")
	content := r.Input().Get("content")
	err := models.AddReply(tid, nikename, content)
	if err != nil {
		beego.Error(err)
	}
	r.Redirect("/topic/view/"+tid, 302)
}

func (r *ReplyController) Delete() {
	if !checkUser(r.Ctx) {
		return
	}
	tid := r.Input().Get("tid")
	rid := r.Input().Get("rid")
	err := models.DeleteReply(rid)
	if err != nil {
		beego.Error(err)
	}
	r.Redirect("/topic/view/"+tid, 302)
}
