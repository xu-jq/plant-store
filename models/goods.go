package models

import "github.com/beego/beego/v2/client/orm"

func GetAllGoods() ([]Goods, error) {
	var goods []Goods
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM goods").QueryRows(&goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func SelectByName(name string) ([]Goods, error) {
	var goods []Goods
	o := orm.NewOrm()
	qs := o.QueryTable("goods")
	_, err := qs.Filter("name__icontains", name).All(&goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func IncreasingOrder() ([]Goods, error) {
	var goods []Goods
	o := orm.NewOrm()
	qs := o.QueryTable("goods")
	_, err := qs.OrderBy("price").All(&goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func DecreasingOrder() ([]Goods, error) {
	var goods []Goods
	o := orm.NewOrm()
	qs := o.QueryTable("goods")
	_, err := qs.OrderBy("-price").All(&goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func SelectById(id int) (error, Goods) {
	var good Goods
	o := orm.NewOrm()
	qs := o.QueryTable("goods")
	err := qs.Filter("id", id).One(&good)
	if err != nil {
		return err, good
	}
	return nil, good
}
