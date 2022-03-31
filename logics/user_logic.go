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

func CheckLogin(username interface{}) (bool, models.Json) {
	user, err := models.GetInfo(username)
	if err != nil || user.Id == 0 {
		logs.Warn("user not login")
		resp := models.Json{
			Flag:    false,
			Message: "请先登录",
			Data:    user,
		}
		return false, resp
	}
	resp := models.Json{
		Flag:    true,
		Message: "get user info success",
		Data:    user,
	}
	return true, resp
}

func Info(username interface{}) models.Json {
	_, resp := CheckLogin(username)
	return resp
}

func ChangeInfo(session interface{}, user models.User) models.Json {
	login, resp := CheckLogin(session)
	if login {
		oldUser, _ := models.GetInfo(session)
		newUser := models.User{}
		if user.Phone == "" {
			newUser.Phone = oldUser.Phone
		} else {
			newUser.Phone = user.Phone
		}
		if user.Email == "" {
			newUser.Email = oldUser.Email
		} else {
			newUser.Email = user.Email
		}
		models.UpdateInfo(newUser, session)
		return models.Json{Flag: true, Message: "更新数据成功"}
	}
	return resp
}
