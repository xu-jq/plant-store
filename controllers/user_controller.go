package controllers

import (
	"bee_project/controllers/base"
	"bee_project/logics"
	"bee_project/messages"
	"fmt"
)

type UserController struct {
	base.BaseController
}

// UserRegister @Title 注册
// @Description 用户注册接口
// @Param username formData string false "用户名"
// @Param password formData string false "密码"
// @Param email formData string false "邮箱"
// @Param phone formData string false "手机号"
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /register [post]
func (c *UserController) UserRegister() {
	request := messages.UserRegisterReq{}
	if errResp := c.ParseForm(&request); errResp != nil {
		c.ServeResponse(base.ErrParse)
		return
	}
	fmt.Println("username:", request.Username)
	ok := logics.CreatNewUser(request)
	if !ok {
		c.ServeResponse(base.ErrUser)
		return
	}
	c.ServeResponse(base.ErrOK)
}

// UserLogin @Title 登录
// @Description 用户登录接口
// @Param username formData string false "用户名"
// @Param password formData string false "密码"
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /login [post]
func (c *UserController) UserLogin() {
	request := messages.UserLoginReq{}
	err := c.ParseForm(&request)
	if err != nil {
		c.ServeResponse(base.ErrParse)
		return
	}

	ok := logics.UserLogin(request)
	if !ok {
		c.ServeResponse(base.ErrLogin)
		return
	}
	err = c.SetSession("user", request.Username)
	if err != nil {
		return
	}
	c.ServeResponse(base.ErrOK)
}

// UserInfo @Title 获取用户信息
// @Description 获取用户信息接口
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /info [get]
func (c *UserController) UserInfo() {
	var itf interface{}
	itf = c.GetSession("user")
	if itf == nil {
		c.ServeResponse(base.ErrGet)
		return
	}
	username := itf.(string)
	user, ok := logics.GetUserByName(username)
	if !ok {
		c.ServeResponse(base.ErrGet)
		return
	}
	c.ServeResponse(base.ErrOK, base.GeneralResponse{Data: user})
}

// ChangeInfo @Title 修改用户信息
// @Description 修改用户信息接口
// @Param email formData string false "邮箱"
// @Param phone formData string false "电话"
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /change_info [put]
func (c *UserController) ChangeInfo() {
	request := messages.UserUpdate{}
	err := c.ParseForm(&request)
	if err != nil {
		c.ServeResponse(base.ErrParse)
		return
	}
	var itf interface{}
	itf = c.GetSession("user")
	if itf == nil {
		c.ServeResponse(base.ErrGet)
		return
	}
	username := itf.(string)
	logics.ChangeInfo(username, request)
	c.ServeResponse(base.ErrOK)
}
