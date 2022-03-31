package logics

import (
	"bee_project/models"
	"github.com/beego/beego/v2/core/logs"
)

func CreatNewUser(user models.User) (bool, models.Json) {
	find := models.FindUserByName(user)
	if find {
		_, err := models.InsertUser(user)
		if err != nil {
			logs.Warn("insert user failed, err: %s", err.Error())
			resp := models.Json{Flag: false, Message: "insert user failed"}
			return false, resp
		}
		logs.Info("create user,username[%s]", user.Username)
		resp := models.Json{Flag: true, Message: "insert user success"}
		return true, resp
	}
	logs.Warn("用户名重复，注册失败")
	resp := models.Json{Flag: false, Message: "用户名重复，注册失败"}
	return false, resp
}

func UserLogin(user models.User) (bool, models.Json) {
	login := models.UserLogin(user)
	if login {
		logs.Info("login success,username[%s]", user.Username)
		resp := models.Json{Flag: true, Message: "登录成功"}
		return true, resp
	}
	logs.Warn("login failed")
	resp := models.Json{Flag: false, Message: "登录失败"}
	return false, resp
}

func Info(username interface{}) models.Json {
	user, err := models.GetInfo(username)
	if err != nil || user.Id == 0 {
		logs.Warn("get user info failed")
		resp := models.Json{
			Flag:    false,
			Message: "请先登录",
			Data:    user,
		}
		return resp
	}
	logs.Info("get user info success,username[%s]", username)
	resp := models.Json{
		Flag:    true,
		Message: "get user info success",
		Data:    user,
	}
	return resp
}
