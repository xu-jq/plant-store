package routers

import (
	"github.com/gin-gonic/gin"
	"plant/controller"
	"plant/middlewares"
)

func User(r *gin.Engine) {
	user := r.Group("/user", middlewares.CORSMiddleware())
	{
		user.POST("/register", controller.Register)                      //注册(相同邮箱不能注册)
		user.POST("/login", controller.Login)                            //登录
		user.GET("/info", middlewares.AuthMiddleware(), controller.Info) //jwt info
	}
}
