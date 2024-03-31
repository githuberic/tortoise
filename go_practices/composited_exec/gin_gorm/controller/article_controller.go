package controller

import (
	"github.com/gin-gonic/gin"
	"go_practices/composited_exec/gin_gorm/service"
	"net/http"
	"strconv"
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
	id := c.Params.ByName("id")
	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithStatus(400)
	}

	articleOne, err := service.GetOneArticle(articleId)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(http.StatusOK, &articleOne)
	}
	return
}

//得到多篇文章，按分页返回
/**
curl -X GET 'http://localhost:8080/article/list?page=1'
*/
func (a *ArticleController) GetList(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.AbortWithStatus(400)
	}

	pageSize := 2
	pageOffset := pageInt * pageSize
	articles, err := service.GetArticleList(pageOffset, pageSize)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(http.StatusOK, &articles)
	}
	return
}
