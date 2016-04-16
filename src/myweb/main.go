package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myweb/models"
	_ "myweb/routers"
)

func init() {
	models.RegisterDB()
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
