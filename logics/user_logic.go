package logics

import (
	"github.com/beego/beego/v2/core/logs"

	"bee_project/models"
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
	//TODO...这个models.Json应该是我们自定义的响应结构体，这个其实用base下的GeneralResponse就行，因为是同一个东西，
	//TODO...这种响应是必定在controllers层返回的，因此放在controller比较合适，而且应该像base下的errors.go一样，把错误封装起来，比如用户名密码错误、未登录等等

	//TODO...此处存在问题，models层应该只负责查用户信息，具体的判断任务应该由logics层去完成
	login := models.UserLogin(user)
	if login {
		logs.Info("login success,username[%s]", user.Username)
		resp := models.Json{Flag: true, Message: "登录成功"}
		return true, resp
	}
	logs.Warn("login failed")
	//TODO...修改响应，将其封装
	resp := models.Json{Flag: false, Message: "登录失败"}
	return false, resp
}

//TODO...该函数命名为GetUserByName比较合适，而且传入参数应该使用string
func Info(username interface{}) models.Json {
	//TODO...此处models.GetInfo这个函数的命名和入参也应该改一下
	user, err := models.GetInfo(username)
	if err != nil || user.Id == 0 {
		logs.Warn("get user info failed")
		//TODO...修改响应，将其封装
		resp := models.Json{
			Flag:    false,
			Message: "请先登录",
			Data:    user,
		}
		return resp
	}
	logs.Info("get user info success,username[%s]", username)
	//TODO...修改响应，将其封装
	resp := models.Json{
		Flag:    true,
		Message: "get user info success",
		Data:    user,
	}
	return resp
}
