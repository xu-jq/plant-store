package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

func InsertUser(user User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FindUserByName(user User) bool {
	var users []User
	o := orm.NewOrm()
	name := user.Username
	num, _ := o.Raw("SELECT * FROM user WHERE username = ?", name).QueryRows(&users)
	if num == 0 {
		return true
	}
	return false
}

func UserLogin(user User) bool {
	var users []User
	o := orm.NewOrm()
	name := user.Username
	password := user.Password
	num, _ := o.Raw("SELECT * FROM user WHERE username = ? AND password = ?", name, password).QueryRows(&users)
	if num == 0 {
		return false
	}
	return true
}

func GetInfo(username interface{}) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM user WHERE username = ?", username).QueryRow(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateInfo(user User, username interface{}) error {
	o := orm.NewOrm()
	err := o.Raw("UPDATE user SET email = ? , phone = ? WHERE username = ?", user.Email, user.Phone, username).QueryRow(&user)
	if err != nil {
		return err
	}
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	return nil
}
