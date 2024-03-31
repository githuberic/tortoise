package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_practices/composited_exec/gin_gorm_tx/global"
	"go_practices/composited_exec/gin_gorm_tx/service"
	"strconv"
)

type GoodsController struct{}

func NewGoodsController() *GoodsController {
	return &GoodsController{}
}

//得到一件商品的详情
func (g GoodsController) GetOne(c *gin.Context) {
	result := global.NewResult(c)

	id := c.Params.ByName("id")
	fmt.Println("id:" + id)

	goodsId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		result.Error(400, "参数错误")
		fmt.Println(err.Error())
		return
	}

	goodsOne, err := service.GetOneGoods(goodsId)
	if err != nil {
		result.Error(404, "数据查询错误")
	} else {
		result.Success(&goodsOne)
	}
	return
}
