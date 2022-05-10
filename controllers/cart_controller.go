package controllers

import (
	"bee_project/controllers/base"
	"bee_project/logics"
	"bee_project/messages"
)

type CartController struct {
	base.BaseController
}

// InsertGoods @Title 商品加入购物车
// @Description 添加购物车接口
// @Param id formData int true "商品id"
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /insert [post]
func (c *CartController) InsertGoods() {
	username := c.Itf(c.GetSession("user"))
	id, _ := c.GetInt("id")
	request := messages.InsertGoodReq{
		Username: username,
		GoodsId:  id,
	}
	success := logics.InsertGood(request)
	if !success {
		c.ServeResponse(base.ErrGet)
	}
	c.ServeResponse(base.ErrOK)
}

// SelectCart @Title 查看购物车
// @Description 查看购物车接口
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router / [get]
func (c *CartController) SelectCart() {
	username := c.Itf(c.GetSession("user"))
	login, cart := logics.SelectCart(username)
	if !login {
		c.ServeResponse(base.ErrGet)
		return
	}
	c.ServeResponse(base.ErrOK, base.GeneralResponse{Data: cart})
}

// DeleteCart @Title 删除购物车商品
// @Description 删除购物车商品接口
// @Param id query int true "id"
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /delete [delete]
func (c *CartController) DeleteCart() {
	username := c.Itf(c.GetSession("user"))
	id, err := c.GetInt("id")
	if err != nil {
		return
	}
	request := messages.DeleteGoodReq{
		Id:       id,
		Username: username,
	}
	login := logics.DeleteCart(request)
	if !login {
		c.ServeResponse(base.ErrGet)
		return
	}
	c.ServeResponse(base.ErrOK)
}
