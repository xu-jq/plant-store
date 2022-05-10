package logics

import (
	"bee_project/messages"
	"github.com/beego/beego/v2/core/logs"

	"bee_project/models"
)

func CreatNewUser(user messages.UserRegisterReq) bool {
	newUser := models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	find := models.FindUserByName(newUser)
	if find {
		_, err := models.InsertUser(newUser)
		if err != nil {
			logs.Warn("insert user failed, err: %s", err.Error())
			return false
		}
		logs.Info("create user,username[%s]", newUser.Username)
		return true
	}
	logs.Warn("用户名重复，注册失败")
	return false
}

func UserLogin(user messages.UserLoginReq) bool {
	loginUser := models.User{
		Username: user.Username,
		Password: user.Password,
	}
	login := models.UserLogin(loginUser)
	if login {
		logs.Info("login success,username[%s]", loginUser.Username)
		return true
	}
	logs.Warn("login failed")
	return false
}

func CheckLogin(username string) bool {
	user, err := models.GetUserByName(username)
	if err != nil || user.Id == 0 || user.Username == "" {
		return false
	}
	return true
}

func GetUserByName(username string) (user models.User, bool bool) {
	user, err := models.GetUserByName(username)
	if err != nil || user.Id == 0 {
		logs.Warn("get user info failed")
		return user, false
	}
	logs.Info("get user info success,username[%s]", username)
	return user, true
}

func ChangeInfo(username string, request messages.UserUpdate) bool {
	login := CheckLogin(username)
	if !login {
		return false
	}
	oldUser, _ := models.GetUserByName(username)
	newUser := models.User{}
	if request.Phone == "" {
		newUser.Phone = oldUser.Phone
	} else {
		newUser.Phone = request.Phone
	}
	if request.Email == "" {
		newUser.Email = oldUser.Email
	} else {
		newUser.Email = request.Email
	}
	newUser.Username = username
	models.UpdateInfo(newUser)
	return true
}
