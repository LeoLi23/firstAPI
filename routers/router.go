package routers

import (
	"firstAPI/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// StudentController
	namespace := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSRouter("/display",&controllers.UserController{}, "get:GetAll"),
			beego.NSRouter("/name", &controllers.UserController{},"post:GetByName"),
			beego.NSRouter("/register",&controllers.UserController{}, "post:CreateUser"),
			beego.NSRouter("/login",&controllers.UserController{},"post:Login"),
			beego.NSRouter("/update",&controllers.UserController{},"put:Update"),
			beego.NSRouter("/delete",&controllers.UserController{},"delete:Delete"),
			),
		)
	beego.AddNamespace(namespace)
}
