package logics

import (
	"bee_project/messages"
	"bee_project/models"
)

func InsertGood(request messages.InsertGoodReq) bool {
	login := CheckLogin(request.Username)
	if login {
		_, good := models.SelectById(request.GoodsId)
		cart := models.SelectAmount(request.GoodsId, request.Username)
		if cart.Amount == 0 {
			newGood := models.Cart{
				Username: request.Username,
				GoodsId:  request.GoodsId,
				Amount:   1,
				Price:    good.Price,
			}
			err := models.InsertGood(newGood)
			if err != nil {
				return false
			}
		}
		newGood := models.Cart{
			Username: request.Username,
			GoodsId:  request.GoodsId,
			Amount:   cart.Amount + 1,
			Price:    cart.Price + good.Price,
		}
		models.UpdateGood(newGood)
		return true
	}
	return false
}

func SelectCart(username string) (bool, []models.Cart) {
	login := CheckLogin(username)
	if login {
		cart, _ := models.SelectCart(username)
		return true, cart
	}
	return false, nil
}

func DeleteCart(request messages.DeleteGoodReq) bool {
	login := CheckLogin(request.Username)
	if login {
		cart := models.Cart{Id: request.Id}
		models.DeleteCart(cart)
		return true
	}
	return false
}
