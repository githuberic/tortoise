package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model //嵌入常用字段
	Code       string
	Price      uint
}

func dbConn(User, Password, Host, Db string, Port int) string {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	return connArgs
}

func main() {
	conn := dbConn("root", "root_mysql", "127.0.0.1", "data_center", 3307)
	//db, err := gorm.Open("mysql", "root:@(localhost:3307)/root_mysql")
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		fmt.Printf("error %v", err)
		//panic("failed to connect database,%v", err)
	}
	//关闭数据库连接
	defer db.Close()

	//创建表
	db.AutoMigrate(&Product{})

	// 插入
	//db.Create(&Product{Code: "L1213", Price: 1200})

	// 读取
	var product Product
	db.First(&product, 1) // 查询id为1的product
	fmt.Printf("%#v\n", product)

	db.Find(&product, "code = ?", "L1213") // 查询code为l1212的product
	fmt.Printf("%#v\n", product)

	// 更新
	db.Model(&product).Update("price", 2200)

	//db.Create(&Product{Code: "L1213", Price: 1200})
	// 删除
	//db.Delete(&product)
}
