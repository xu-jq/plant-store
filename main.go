package main

import (
	"github.com/gin-gonic/gin"

	"plant/routers"
)

// @title 前后端分离商城
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	routers.User(router)

	router.Run()
}
