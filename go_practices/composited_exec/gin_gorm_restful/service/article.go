package service

import (
	"go_practices/composited_exec/gin_gorm_restful/dao"
	"go_practices/composited_exec/gin_gorm_restful/model"
)

/***
得到一篇文章的详情
*/
func GetOneArticle(articleId int64) (*model.Article, error) {
	return dao.SelectOneArticle(articleId)
}

//得到多篇文章，按分页返回
func GetArticleList(page int, pageSize int) ([]*model.Article, error) {
	articles, err := dao.SelectAllArticle(page, pageSize)
	if err != nil {
		return nil, err
	} else {
		return articles, nil
	}
}
