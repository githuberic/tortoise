package dao

import (
	"go_practices/composited_exec/gin_gorm_tx/global"
	"go_practices/composited_exec/gin_gorm_tx/model"
)

// 得到一件商品的详情
func SelectOneGoods(goodsId int64) (*model.Goods, error) {
	goodsOne := &model.Goods{}
	err := global.DBLink.Where("goodsId=?", goodsId).First(&goodsOne).Error
	if err != nil {
		return nil, err
	} else {
		return goodsOne, nil
	}
}
