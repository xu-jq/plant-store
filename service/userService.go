package service

import (
	"plant/dao"
	"plant/model"
)

// CheckEmail 保证邮箱唯一（同一邮箱不能注册）
func CheckEmail(user model.User) model.Json {
	a := dao.SelectEmail(user)
	if a {
		json := dao.AddUser(user)
		return json
	} else {
		json := model.Json{
			Code: 200,
			Flag: true,
			Msg:  "该邮箱已注册",
		}
		return json
	}
}

func Login(user model.User) model.Json {

	res := dao.Login(user)
	if res {
		json := model.Json{
			Code: 200,
			Flag: true,
			Msg:  "登录成功",
		}
		return json
	} else {
		json := model.Json{
			Code: 200,
			Flag: false,
			Msg:  "登录失败，请重试",
		}
		return json
	}
}
