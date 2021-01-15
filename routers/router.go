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
	beego.Router("/",&controllers.StudentController{}, "get:GetAll")
	beego.Router("/id", &controllers.StudentController{},"get:GetById")
	beego.Router("/register",&controllers.StudentController{}, "post:Post")
	beego.Router("/update",&controllers.StudentController{},"put:Update")
	beego.Router("/delete",&controllers.StudentController{},"delete:Delete")
}
