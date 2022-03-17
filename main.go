package main

import (
	"github.com/gin-gonic/gin"

	"plant/routers"
)

func main() {
	router := gin.Default()

	routers.User(router)

	router.Run()
}
