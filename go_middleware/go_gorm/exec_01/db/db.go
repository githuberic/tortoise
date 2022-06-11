package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//参数含义:数据库用户名、密码、主机ip、连接的数据库、端口号
func dbConn(User, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		return nil
	}
	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(true)       //打印sql语句

	//开启连接池
	db.DB().SetMaxIdleConns(100)   //最大空闲连接
	db.DB().SetMaxOpenConns(10000) //最大连接数
	db.DB().SetConnMaxLifetime(30) //最大生存时间(s)

	return db
}

func GetDb() (conn *gorm.DB) {
	for {
		conn = dbConn("root", "root_mysql", "127.0.0.1", "data_center", 3307)
		if conn != nil {
			fmt.Printf("db error:%v", nil)
			fmt.Println("本次未获取到mysql连接")
		} else {
			return conn
		}
	}
}
