package dao

import (
	"plant/common"
	"plant/model"
)

// AddUser 添加用户
func AddUser(user model.User) model.Json {
	db := common.InitDB()
	db.Save(&user)
	return model.Json{
		Code: 200,
		Flag: true,
		Msg:  "添加成功",
	}
}

// SelectEmail 查找相同邮箱
func SelectEmail(user model.User) bool {
	db := common.InitDB()
	res := db.Where("email = ?", user.Email).First(&user)
	affected := res.RowsAffected
	if affected == 0 {
		return true
	} else {
		return false
	}
}

// Login 校验账号密码(邮箱登录)
func Login(user model.User) bool {
	db := common.InitDB()
	res := db.Where("email = ? AND password = ?", user.Email, user.Password).Find(&user)
	affected := res.RowsAffected
	if affected != 0 {
		return true
	}
	return false
}
