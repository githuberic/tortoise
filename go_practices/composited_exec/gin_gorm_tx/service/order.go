package service

import (
	"go_practices/composited_exec/gin_gorm_tx/dao"
	"go_practices/composited_exec/gin_gorm_tx/model"
)

// 新加一个订单
func AddOneOrder(goodsId int64, buyNum int) (*model.Order, error) {
	return dao.AddOneOrder(goodsId, buyNum)
}

// 新加一个订单,使用tx
func AddOneOrderTx(goodsId int64, buyNum int) (*model.Order, error) {
	return dao.AddOneOrderTx(goodsId, buyNum)
}
