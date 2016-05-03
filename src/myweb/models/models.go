package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/myweb.db"
	_SQLITE3_DEIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Views           int64 `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:size(5000)`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Author           string
	RepalyTime       time.Time `orm:"index"`
	RepalyCount      int64
	ReplayLastUserId int64
	Category         string
}
type Commit struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:size(1000)`
	Created time.Time `orm:index`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic), new(Commit))

	orm.RegisterDriver(_SQLITE3_DEIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DEIVER, _DB_NAME, 10)

}
func AddReply(tid, nikename, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	reply := &Commit{
		Tid:     tidNum,
		Name:    nikename,
		Content: content,
		Created: time.Now(),
	}
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}
	return nil
}
func GetAllReplies(tid string) (replies []*Commit, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	qs := o.QueryTable("commit")
	replies = make([]*Commit, 0)
	_, err = qs.Filter("tid", tidNum).All(&replies)
	if err != nil {
		return nil, err
	}
	return replies, err

}
func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	reply := &Commit{Id: ridNum}
	_, err = o.Delete(reply)
	if err != nil {
		return err
	}
	return nil
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	if err == nil {
		return err
	}
	return nil
}
func DeleteCategory(id string) error {

	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	if err != nil {
		return err

	}
	return nil
}
func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func AddTopic(title string, content string, category string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:      title,
		Content:    content,
		Category:   category,
		Created:    time.Now(),
		Updated:    time.Now(),
		RepalyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}
func GetAllTopic(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func TopicModify(tid, title, content, category string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		topic.Category = category
		o.Update(topic)
	}
	return nil
}
