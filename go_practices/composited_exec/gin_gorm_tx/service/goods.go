package service

import (
	"go_practices/composited_exec/gin_gorm_tx/dao"
	"go_practices/composited_exec/gin_gorm_tx/model"
)

// 得到一件商品的详情
func GetOneGoods(goodsId int64) (*model.Goods, error) {
	return dao.SelectOneGoods(goodsId)
}
