package controller

import (
	"github.com/gin-gonic/gin"
	"go_practices/composited_exec/gin_gorm_viper_apicode/global"
	"go_practices/composited_exec/gin_gorm_viper_apicode/pkg/page"
	"go_practices/composited_exec/gin_gorm_viper_apicode/pkg/result"
	"go_practices/composited_exec/gin_gorm_viper_apicode/pkg/validCheck"
	"go_practices/composited_exec/gin_gorm_viper_apicode/request"
	"go_practices/composited_exec/gin_gorm_viper_apicode/service"
)

type ArticleController struct{}

func NewArticleController() ArticleController {
	return ArticleController{}
}

//得到一篇文章的详情
func (a *ArticleController) GetOne(c *gin.Context) {
	result := result.NewResult(c)
	param := request.ArticleRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	valid, _ := validCheck.BindAndValid(c, &param)
	if !valid {
		result.Error(global.ErrParam)
		return
	}

	articleOne, err := service.GetOneArticle(param.ID)
	if err != nil {
		result.Error(global.ErrArticleNot)
	} else {
		result.Success(&articleOne)
	}
	return
}

//得到多篇文章，按分页返回
func (a *ArticleController) GetList(c *gin.Context) {
	result := result.NewResult(c)
	pageInt := 0
	//is exist?
	curPage := c.Query("page")
	//if curPage not exist
	if len(curPage) == 0 {
		pageInt = 1
	} else {
		param := request.ArticleListRequest{Page: validCheck.StrTo(c.Param("page")).MustInt()}
		valid, _ := validCheck.BindAndValid(c, &param)
		if !valid {
			result.Error(global.ErrParam)
			return
		}
		pageInt = param.Page
	}

	pageSize := 2
	pageOffset := (pageInt - 1) * pageSize

	articles, err := service.GetArticleList(pageOffset, pageSize)
	if err != nil {
		result.Error(global.ErrArticleS)
	} else {
		sum, _ := service.GetArticleSum()
		pageInfo, _ := page.GetPageInfo(pageInt, pageSize, sum)
		result.Success(gin.H{"list": &articles, "pageinfo": pageInfo})
	}
	return
}
