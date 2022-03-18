package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"plant/common"
	"plant/model"
	"plant/service"
)

// Register 用户注册
// @Summary 用户注册接口
// @Description 传入用户信息进行注册
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Json
// @Router /user/register [post]
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

// Login 用户登录
// @Summary 用户登录接口
// @Description 传入用户名和密码进行登录
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param string formData model.User false "用户名和密码"
// @Success 200 {object} model.Json
// @Router /user/login [post]
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
// @Summary 解码token获取信息
// @Description 解码token获取信息
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param string query string false "token串"
// @Success 200 {object} model.Json
// @Router /user/info [get]
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"Code": 200, "Data": user})
}
