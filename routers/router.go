// Package routers @APIVersion 1.0.0
// @Title plant_store
package routers

import (
	"bee_project/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns :=
		beego.NewNamespace("",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
		)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")

	beego.Router("/user/register", &controllers.UserController{}, "post:UserRegister")
	beego.Router("/user/login", &controllers.UserController{}, "post:UserLogin")
	beego.Router("/user/info", &controllers.UserController{}, "get:UserInfo")
}
