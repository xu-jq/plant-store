package models

import (
	"github.com/beego/beego/v2/client/orm"
)

func InsertGood(newGood Cart) error {
	o := orm.NewOrm()
	_, err := o.Insert(&newGood)
	if err != nil {
		return err
	}
	return nil
}

func SelectAmount(gid int, username string) Cart {
	var cart Cart
	o := orm.NewOrm()
	o.Raw("SELECT * FROM cart WHERE username = ? AND goods_id = ?", username, gid).QueryRow(&cart)
	return cart
}

func UpdateGood(cart Cart) {
	o := orm.NewOrm()
	o.Raw("UPDATE cart SET amount = ? , price = ? WHERE username = ? AND goods_id = ?", cart.Amount, cart.Price, cart.Username, cart.GoodsId).QueryRow(&cart)
}

func SelectCart(username string) ([]Cart, error) {
	var cart []Cart
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM cart WHERE username = ?", username).QueryRows(&cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func DeleteCart(cart Cart) {
	o := orm.NewOrm()
	_, err := o.Delete(&cart)
	if err != nil {
		return
	}
}
