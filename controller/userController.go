package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plant/common"
	"plant/model"
	"plant/service"
)

// Register 注册
func Register(c *gin.Context) {
	name1 := c.PostForm("name1")
	name2 := c.PostForm("name2")
	name := name1 + name2
	user := model.User{
		Username: name,
		Email:    c.PostForm("email"),
		Phone:    c.PostForm("phone"),
		Password: c.PostForm("pwd"),
	}
	json := service.CheckEmail(user)
	c.JSON(http.StatusOK, json)
}

// Login 登录
func Login(c *gin.Context) {
	user := model.User{Email: c.PostForm("username"), Password: c.PostForm("password")}
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Code": 500,
			"Msg":  "系统异常",
		})
	} else {
		json := service.Login(user)
		json.Data = token
		c.JSON(http.StatusOK, json)
	}
	//json := service.Login(user)

}

// Info 解码token获取信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"Code": 200, "Data": user})
}
