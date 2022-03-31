// Package routers
// @APIVersion 2.0.0
// @Title plant_store
// @Description plant_store APIs
// @Contact 604862834@qq.com
package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"bee_project/controllers"
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
	beego.Router("/user/change_info", &controllers.UserController{}, "get:ChangeInfo")

}
