package models

import "github.com/beego/beego/v2/client/orm"

type Json struct {
	Flag    bool        `json:"flag"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func init() {
	orm.RegisterModel(
		new(User),
	)
}
