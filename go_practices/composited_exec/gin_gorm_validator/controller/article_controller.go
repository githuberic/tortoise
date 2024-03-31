package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_practices/composited_exec/gin_gorm_validator/global"
	"go_practices/composited_exec/gin_gorm_validator/pkg/page"
	"go_practices/composited_exec/gin_gorm_validator/pkg/valid_check"
	"go_practices/composited_exec/gin_gorm_validator/request"
	"go_practices/composited_exec/gin_gorm_validator/service"
)

type ArticleController struct{}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

/**
得到一篇文章的详情
curl -X GET 'http://localhost:8080/article/getone/3'
*/
func (a *ArticleController) GetOne(c *gin.Context) {
	result := global.NewResult(c)
	param := request.ArticleRequest{ID: valid_check.StrTo(c.Param("id")).MustUInt64()}
	valid, errs := valid_check.BindAndValid(c, &param)
	if !valid {
		result.Error(400, errs.Error())
		return
	}

	articleOne, err := service.GetOneArticle(param.ID)
	if err != nil {
		result.Error(404, "数据查询错误")
	} else {
		result.Success(&articleOne)
	}
	return
}

//得到多篇文章，按分页返回
/**
curl -X GET 'http://localhost:8080/article/list?page=1'
*/
func (a *ArticleController) GetList(c *gin.Context) {
	result := global.NewResult(c)
	pageInt := 0
	//is exist?
	curPage := c.Query("page")
	//if curPage not exist
	if len(curPage) == 0 {
		pageInt = 1
	} else {
		param := request.ArticleListRequest{Page: valid_check.StrTo(c.Param("page")).MustInt()}
		valid, errs := valid_check.BindAndValid(c, &param)
		if !valid {
			result.Error(400, errs.Error())
			return
		}
		pageInt = param.Page
	}

	pageSize := 2
	pageOffset := (pageInt - 1) * pageSize

	articles, err := service.GetArticleList(pageOffset, pageSize)
	if err != nil {
		result.Error(404, "数据查询错误")
		fmt.Println(err.Error())
	} else {
		sum, _ := service.GetArticleSum()
		pageInfo, _ := page.GetPageInfo(pageInt, pageSize, sum)
		result.Success(gin.H{"list": &articles, "pageinfo": pageInfo})
	}
	return
}
