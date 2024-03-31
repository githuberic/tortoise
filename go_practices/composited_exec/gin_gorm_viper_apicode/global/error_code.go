package global

import "go_practices/composited_exec/gin_gorm_viper_apicode/pkg/result"

var (
	// OK
	OK = result.NewError(0, "OK")

	//参数模块
	ErrParam = result.NewError(400, "参数不合法")

	//文章模块报错
	ErrArticleNot = result.NewError(10001, "文章不存在")
	ErrArticleS   = result.NewError(10002, "文章查询出错")
)
