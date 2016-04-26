package controllers

import (
	"github.com/astaxie/beego"
	"myweb/models"
)

type TopicController struct {
	beego.Controller
}

func (t *TopicController) Get() {
	t.Data["IsLogin"] = checkUser(t.Ctx)
	t.Data["IsTopic"] = true
	t.TplName = "topic.html"
	topics, err := models.GetAllTopic(false)
	if err != nil {
		beego.Error(err)
	}
	t.Data["Topics"] = topics
}
func (t *TopicController) Post() {
	if !checkUser(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	title := t.Input().Get("title")
	content := t.Input().Get("content")

	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	t.Redirect("/topic", 302)
	return

}

func (t *TopicController) Add() {
	t.Data["IsTopic"] = true
	t.TplName = "topic_add.html"
}
