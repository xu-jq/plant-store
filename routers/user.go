package routers

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"plant/controller"
	_ "plant/docs"
	"plant/middlewares"
)

func User(r *gin.Engine) {
	user := r.Group("/user", middlewares.CORSMiddleware())
	{
		user.POST("/register", controller.Register)                      //注册(相同邮箱不能注册)
		user.POST("/login", controller.Login)                            //登录
		user.GET("/info", middlewares.AuthMiddleware(), controller.Info) //jwt info
	}

	// 通过 http://localhost:8080/swagger/index.html 查看api文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
}
