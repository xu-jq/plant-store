package messages

type InsertGoodReq struct {
	Username string `json:"username"`
	GoodsId  int    `json:"goods_id"`
}

type DeleteGoodReq struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
