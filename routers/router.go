// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"firstAPI/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// StudentController
	namespace := beego.NewNamespace("/school",
		beego.NSNamespace("/student",
			beego.NSRouter("/",&controllers.StudentController{}, "get:GetAll"),
			beego.NSRouter("/id", &controllers.StudentController{},"get:GetById"),
			beego.NSRouter("/register",&controllers.StudentController{}, "post:Post"),
			beego.NSRouter("/login",&controllers.StudentController{},"post:Login"),
			beego.NSRouter("/update",&controllers.StudentController{},"put:Update"),
			beego.NSRouter("/delete",&controllers.StudentController{},"delete:Delete"),
			),
		)
	beego.AddNamespace(namespace)
}
