package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Goods struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Color string `json:"color"`
	Stock int    `json:"stock"`
}

type Cart struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	GoodsId  int    `json:"goods_id"`
	Amount   int    `json:"amount"`
	Price    int    `json:"price"`
}

func init() {
	orm.RegisterModel(
		new(User), new(Goods), new(Cart),
	)
}
