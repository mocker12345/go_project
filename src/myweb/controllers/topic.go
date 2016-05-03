package controllers

import (
	"github.com/astaxie/beego"
	"myweb/models"
)

type TopicController struct {
	beego.Controller
}

func (t *TopicController) Get() {
	t.TplName = "topic.html"
	t.Data["IsLogin"] = checkUser(t.Ctx)
	t.Data["IsTopic"] = true
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
	tid := t.Input().Get("tid")
	category := t.Input().Get("category")
	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content, category)
	} else {
		err = models.TopicModify(tid, title, content, category)
	}
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
func (t *TopicController) View() {
	t.TplName = "topic_view.html"
	topic, err := models.GetTopic(t.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
		t.Redirect("/", 302)
		return
	}
	replise, err := models.GetAllReplies(t.Ctx.Input.Params()["0"])
	t.Data["Replies"] = replise
	t.Data["Topic"] = topic
	t.Data["IsLogin"] = checkUser(t.Ctx)
}
func (t *TopicController) Modify() {
	t.TplName = "topic_modify.html"
	tid := t.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		t.Redirect("/", 302)
		return
	}
	t.Data["Topic"] = topic
}

func (t *TopicController) Delete() {
	if !checkUser(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(t.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	t.Redirect("/", 302)
	return
}
