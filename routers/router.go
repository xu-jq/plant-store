// Package routers @APIVersion 1.0.0
// @Title plant_store
package routers

import (
	"bee_project/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	ns :=
		beego.NewNamespace("",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSNamespace("/good",
				beego.NSInclude(
					&controllers.GoodsController{},
				),
			),
			beego.NSNamespace("/cart",
				beego.NSInclude(
					&controllers.CartController{},
				),
			),
		)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")

	//user
	beego.Router("/user/register", &controllers.UserController{}, "post:UserRegister")
	beego.Router("/user/login", &controllers.UserController{}, "post:UserLogin")
	beego.Router("/user/info", &controllers.UserController{}, "get:UserInfo")
	beego.Router("/user/change_info", &controllers.UserController{}, "put:ChangeInfo")

	//good
	beego.Router("/good", &controllers.GoodsController{}, "get:GetAllGoods")
	beego.Router("/good/select", &controllers.GoodsController{}, "get:SelectByName")
	beego.Router("/good/increase", &controllers.GoodsController{}, "get:IncreasingOrder")
	beego.Router("/good/decrease", &controllers.GoodsController{}, "get:DecreasingOrder")

	//cart
	beego.Router("/cart/insert", &controllers.CartController{}, "post:InsertGoods")
	beego.Router("/cart", &controllers.CartController{}, "get:SelectCart")
	beego.Router("/cart/delete", &controllers.CartController{}, "delete:DeleteCart")
}
