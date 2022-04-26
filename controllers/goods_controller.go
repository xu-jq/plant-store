package controllers

import (
	"bee_project/controllers/base"
	"bee_project/models"
)

type GoodsController struct {
	base.BaseController
}

// GetAllGoods @Title 获取所有商品
// @Description 获取所有商品接口
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router / [get]
func (c *GoodsController) GetAllGoods() {
	goods, err := models.GetAllGoods()
	if err != nil {
		c.ServeResponse(base.ErrDatabase)
		return
	}
	c.ServeResponse(base.ErrOK, base.GeneralResponse{Data: goods})
}

// SelectByName @Title 模糊查询商品
// @Description 模糊查询商品接口
// @Param name query string false "商品名"
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /select [get]
func (c *GoodsController) SelectByName() {
	name := c.GetString("name")
	goods, err := models.SelectByName(name)
	if err != nil {
		c.ServeResponse(base.ErrDatabase)
		return
	}
	c.ServeResponse(base.ErrOK, base.GeneralResponse{Data: goods})
}

// IncreasingOrder @Title 商品按价格升序排序
// @Description 商品按价格升序排序接口
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /increase [get]
func (c *GoodsController) IncreasingOrder() {
	goods, err := models.IncreasingOrder()
	if err != nil {
		c.ServeResponse(base.ErrDatabase)
		return
	}
	c.ServeResponse(base.ErrOK, base.GeneralResponse{Data: goods})
}

// DecreasingOrder @Title 商品按价格降序排序
// @Description 商品按价格降序排序接口
// @Success 200 {object} base.GeneralResponse
// @Failure 500 {object} base.GeneralResponse
// @router /decrease [get]
func (c *GoodsController) DecreasingOrder() {
	goods, err := models.DecreasingOrder()
	if err != nil {
		c.ServeResponse(base.ErrDatabase)
		return
	}
	c.ServeResponse(base.ErrOK, base.GeneralResponse{Data: goods})
}

func (c *GoodsController) GetGoodsByPage() {
	pageSize, _ := c.GetInt("pageSize")
	pageNum, _ := c.GetInt("pageNum")
	goods, _, err := models.GetGoodsByPage(pageSize, pageNum)
	if err != nil {
		c.ServeResponse(base.ErrDatabase)
		return
	}
	c.ServeResponse(base.ErrOK, base.GeneralResponse{Data: goods})
}
