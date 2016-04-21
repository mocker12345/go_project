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
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))

	orm.RegisterDriver(_SQLITE3_DEIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DEIVER, _DB_NAME, 10)

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
func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
