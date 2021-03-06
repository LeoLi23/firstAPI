package main

import (
	"firstAPI/models"
	_ "firstAPI/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {// classified into models/index.go ==> initialize mysql database(remote)
	orm.Debug = true

	if err1 := orm.RegisterDriver("mysql", orm.DRMySQL); err1 != nil {
		logs.Error(err1.Error())
	}

	orm.RegisterModel(new(models.User))

	if err2 := orm.RegisterDataBase("default","mysql","root:12345678@tcp(127.0.0.1:3306)/test");err2 != nil {
		logs.Error(err2.Error())
		panic(err2.Error())
	}
	fmt.Println("Connected to the database")
	_ = orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
