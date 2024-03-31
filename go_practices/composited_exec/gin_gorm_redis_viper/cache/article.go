package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go_practices/composited_exec/gin_gorm_redis_viper/global"
	"go_practices/composited_exec/gin_gorm_redis_viper/model"
	"strconv"
	"time"
)

// article的过期时长
const ArticleDuration = time.Minute * 5

//cache的名字
func getArticleCacheName(articleId uint64) string {
	return "article_" + strconv.FormatUint(articleId, 10)
}

var ctx = context.Background()

//从cache得到一篇文章
func GetOneArticleCache(articleId uint64) (*model.Article, error) {
	key := getArticleCacheName(articleId)
	val, err := global.RedisDb.Get(ctx, key).Result()

	if err == redis.Nil || err != nil {
		return nil, err
	} else {
		article := model.Article{}
		if err := json.Unmarshal([]byte(val), &article); err != nil {
			//t.Error(target)
			return nil, err
		}
		return &article, nil
	}
}

//向cache保存一篇文章
func SetOneArticleCache(articleId uint64, article *model.Article) error {
	key := getArticleCacheName(articleId)
	content, err := json.Marshal(article)
	if err != nil {
		return err
	}
	errSet := global.RedisDb.Set(ctx, key, content, ArticleDuration).Err()
	if errSet != nil {
		return errSet
	}
	return nil
}
