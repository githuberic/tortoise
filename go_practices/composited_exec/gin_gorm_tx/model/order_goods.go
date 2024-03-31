package model

type OrderGoods struct {
	OgId    int64 `gorm:"column:ogId",json:"ogid"`       // 自增
	OrderId int64 `gorm:"column:orderId",json:"orderid"` // 订单id
	GoodsId int64 `gorm:"column:goodsId",json:"goodsid"` // 用户id
	BuyNum  int   `gorm:"column:buyNum",json:"buynum"`   // 购买数量
}

func (OrderGoods) TableName() string {
	return "m_order_goods"
}
