package controllers

import (
	"bee_project/controllers/base"
	"bee_project/logics"
	"bee_project/models"
)

type UserController struct {
	base.BaseController
}

// UserRegister @Title 注册
// @Description 用户注册接口
// @Param name1 formData string false "姓"
// @Param name2 formData string false "名"
// @Param email formData string false "邮箱"
// @Param phone formData string false "手机号码"
// @Param password formData string false "密码"
// @Success 200 {object} models.Json
// @Failure 500 {object} models.Json
// @router /register [post]
func (c *UserController) UserRegister() {
	name1 := c.GetString("name1")
	name2 := c.GetString("name2")
	name := name1 + name2
	user := models.User{
		Username: name,
		Email:    c.GetString("email"),
		Phone:    c.GetString("phone"),
		Password: c.GetString("pwd"),
	}
	ok, resp := logics.CreatNewUser(user)
	c.Ctx.Output.JSON(resp, true, true)
	if !ok {
		c.ServeResponse(base.ErrDatabase)
		return
	}
	c.ServeResponse(base.ErrOK)
}

// UserLogin @Title 登录
// @Description 用户登录接口
// @Param username formData string false "用户名"
// @Param password formData string false "密码"
// @Success 200 {object} models.Json
// @Failure 500 {object} models.Json
// @router /login [post]
func (c *UserController) UserLogin() {
	user := models.User{Username: c.GetString("username"), Password: c.GetString("password")}
	ok, resp := logics.UserLogin(user)
	c.Ctx.Output.JSON(resp, true, true)
	if !ok {
		c.ServeResponse(base.ErrDatabase)
		return
	}
	c.SetSession("user", user.Username)
	c.GetSession("user")
	c.ServeResponse(base.ErrOK)
}

// UserInfo @Title 获取用户信息
// @Description 获取用户信息接口
// @Success 200 {object} models.Json
// @Failure 500 {object} models.Json
// @router /info [get]
func (c *UserController) UserInfo() {
	username := c.GetSession("user")
	resp := logics.Info(username)
	c.Ctx.Output.JSON(resp, true, true)
}
