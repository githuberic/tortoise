package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_practices/composited_exec/gin_gorm_tx/global"
	"go_practices/composited_exec/gin_gorm_tx/service"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

//新加一个订单
func (o *OrderController) AddOne(c *gin.Context) {
	result := global.NewResult(c)

	goodsId := 1
	buyNum := 11

	orderOne, err := service.AddOneOrder(int64(goodsId), buyNum)
	if err != nil {
		result.Error(404, "数据处理错误")
	} else {
		result.Success(&orderOne)
	}
	return
}

//新加一个订单,使用tx
func (o *OrderController) AddOneTx(c *gin.Context) {
	result := global.NewResult(c)
	goodsId := 1
	buyNum := 11

	orderOne, err := service.AddOneOrderTx(int64(goodsId), buyNum)

	fmt.Println("orderOne:")
	fmt.Println(orderOne)
	fmt.Println(":end:")
	//if ()
	if orderOne == nil || err != nil {
		result.Error(20001, "数据处理错误")
	} else {
		result.Success(&orderOne)
	}
	return
}
