package models

import (
	"bee_project/libs/tools"
	"github.com/beego/beego/v2/client/orm"
)

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

func GetGoodsByPage(pageSize int, pageNum int) ([]Goods, int64, error) {
	o := orm.NewOrm()
	var goods = make([]Goods, 0)
	count, err := o.Raw("select * from goods").QueryRows(&goods)
	qs := o.QueryTable("goods") //1获取数据总数
	/*	qs:=o.QueryTable("Oauth2Mids") //2获取数据总数
		count,err := qs.Count()
		fmt.Println("count1111111",count)*/
	//fmt.Println("num111111111111", count) //数据库返回该表数据总条数
	//	总页数
	pageNum = tools.Page(pageSize, pageNum, count)
	//获取分页数据
	_, err = qs.Limit(pageSize, pageSize*(pageNum-1)).All(&goods)
	if err != nil {
		return goods, count, err
	}
	return goods, count, err
}
