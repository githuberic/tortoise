package global

import (
	"github.com/jinzhu/gorm"
	"time"
)

var (
	DBLink *gorm.DB
)

func SetupDBLink() error {
	var err error
	DBLink, err = gorm.Open("mysql", "root:root_mysql@tcp(127.0.0.1:3307)/gorm_exec_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	// 全局禁用表名复数
	DBLink.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	//打开sql日志
	DBLink.LogMode(true)

	sqlDB := DBLink.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(30)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
