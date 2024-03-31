package model

type Order struct {
	OrderId   int64  `gorm:"primaryKey;autoIncrement;column:orderId",json:"orderid"` // 自增
	UserId    int64  `gorm:"column:userId",json:"userid"`                            // 用户id
	SalePrice string `gorm:"column:salePrice",json:"saleprice"`                      // 售价
}

func (Order) TableName() string {
	return "m_order"
}
