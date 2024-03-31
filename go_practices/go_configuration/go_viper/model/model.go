package model

import (
	"errors"
	"fmt"
	orm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"sync"
)

type MySqlPool struct{}

var instance *MySqlPool
var once sync.Once

var db *orm.DB
var err error

// 单例模式
func GetInstance() *MySqlPool {
	once.Do(func() {
		instance = &MySqlPool{}
	})

	return instance
}

func (pool *MySqlPool) InitPool() (isSuc bool) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.name"), viper.GetString("db.charset"))
	db, err = orm.Open("mysql", dsn)
	if err != nil {
		panic(errors.New("mysql连接失败"))
		return false
	}

	// 连接数配置也可以写入配置，在此读取
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(50)
	// db.LogMode(true)
	return true
}
