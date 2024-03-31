package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"go_practices/composited_exec/gin_gorm_tx/global"
	"go_practices/composited_exec/gin_gorm_tx/model"
)

//添加一个订单
func AddOneOrder(goodsId int64, buyNum int) (*model.Order, error) {
	// 添加order
	order := model.Order{UserId: 1, SalePrice: "20.00"}
	resultCO := global.DBLink.Create(&order) // 通过数据的指针来创建
	if resultCO.Error != nil {
		return nil, resultCO.Error
	}

	// 减库存
	result := global.DBLink.Debug().Table("m_goods").Where("goodsId = ? and stock >= ?", goodsId, buyNum).Update("stock", gorm.Expr("stock - ?", buyNum))
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected <= 0 {
		return nil, errors.New("减库存失败")
	}

	// 添加订单商品
	orderId := order.OrderId
	orderGoods := model.OrderGoods{OrderId: orderId, GoodsId: goodsId, BuyNum: buyNum}
	resultCOG := global.DBLink.Create(&orderGoods) // 通过数据的指针来创建
	if resultCOG.Error != nil {
		return nil, resultCOG.Error
	}

	//return
	return &order, nil
}

// 添加一个订单
func AddOneOrderTx(goodsId int64, buyNum int) (*model.Order, error) {
	// 事务开始后，需要使用 tx 处理数据
	tx := global.DBLink.Begin()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("this is in defer recover")
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	// 添加order
	order := model.Order{UserId: 1, SalePrice: "20.00"}
	resultCO := tx.Debug().Create(&order) // 通过数据的指针来创建
	if (resultCO.Error != nil) {
		tx.Rollback()
		return nil, resultCO.Error
	}

	/*
		var z int = 0
		var i int = 100 / z
		fmt.Println("i:%i",i)
	*/

	// 减库存
	result := tx.Debug().Table("m_goods").Where("goodsId = ? and stock >= ?", goodsId, buyNum).Update("stock", gorm.Expr("stock - ?", buyNum))
	if (result.Error != nil) {
		tx.Rollback()
		return nil, result.Error
	}
	if (result.RowsAffected <= 0) {
		tx.Rollback()
		return nil, errors.New("减库存失败")
	}

	// 添加订单商品
	orderId := order.OrderId;
	orderGoods := model.OrderGoods{OrderId: orderId, GoodsId: goodsId, BuyNum: buyNum}
	resultCOG := tx.Debug().Create(&orderGoods) // 通过数据的指针来创建
	if (resultCOG.Error != nil) {
		tx.Rollback()
		return nil, resultCOG.Error
	}

	// Commit tx
	errCM := tx.Commit().Error
	if (errCM != nil) {
		tx.Rollback()
		return nil, errCM
	} else {
		return &order, nil
	}
}
