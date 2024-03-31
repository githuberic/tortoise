package model

type Goods struct {
	GoodsId   int64  `gorm:"column:goodsId",json:"goodsId"`     // 自增
	GoodsName string `gorm:"column:goodsName",json:"goodsName"` // 用户id
	Stock     int    `gorm:"column:stock",json:"stock"`         // 售价
}

func (Goods) TableName() string {
	return "m_goods"
}
