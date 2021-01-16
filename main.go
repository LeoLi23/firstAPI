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

func init() {
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
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
	//lr := &models.LoginRequest{
	//	Username : "someone",
	//	Password : "mypassword",
	//}
	//token, _ := models.GenerateToken(lr,2, 0)
	//j, _ := models.ValidateToken(token)
	//fmt.Println("Token: ", token)
	//fmt.Println("Issued at: ",j.IssuedAt)
	//fmt.Println("Expired at: ", j.ExpiresAt)
	//_,_ = new(controllers.UserController).CheckStatus("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjIsImV4cCI6MTYxMDc5MDc3MiwiaWF0IjoxNjEwNzkwNzEyLCJpc3MiOiJzaGl5aSJ9.uMv4lDhMPzldV4A7wnx1mIf54pGb9yOgTTjYBSXDwkA")

}
