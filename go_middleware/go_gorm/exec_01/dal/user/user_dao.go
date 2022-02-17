package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id      int    `gorm:"primary_key" json:"id"`
	Name    string `json:"name"`
	Address string `json:"string"`
	Mobile  string `json:"mobile"`
}

func dbConn(User, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		panic("failed to connect database,%v")
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
	dbConn := dbConn("root", "root_mysql", "127.0.0.1", "db_user", 3307)
	if dbConn == nil {
		fmt.Printf("db error:%v", nil)
		panic("Failed to connect database,%v")
	}
	return dbConn
}

//添加数据
func (user *User) Add() {
	db := GetDb()
	//关闭数据库连接
	defer db.Close()

	//user := &User{Name: "L1214", Address: "中国-浙江-杭州", Mobile: "13588829999"}
	err := db.Create(user).Error
	if err != nil {
		fmt.Printf("创建失败 %v", err)
	}
}

//修改数据
func (user *User) Update() {
	db := GetDb()
	//关闭数据库连接
	defer db.Close()

	err := db.Model(user).Update(user).Error
	if err != nil {
		fmt.Println("修改失败")
	}
}

func (user *User) Updates(colValues map[string]interface{}) {
	db := GetDb()
	//关闭数据库连接
	defer db.Close()

	err := db.Model(user).Updates(colValues).Error
	if err != nil {
		fmt.Println("修改失败")
	}
}

func main() {
	db := GetDb()
	//关闭数据库连接
	defer db.Close()

	//创建表
	db.AutoMigrate(&User{})

	var user User
	user = User{Name: "L1215", Address: "中国-浙江-杭州-西湖区", Mobile: "13588828888"}

	user.Add()
	// 插入
	// db.Create(&User{Name: "L1214", Address: "中国-浙江-杭州", Mobile: "13588829999"})

	// 读取
	db.First(&user, 103) // 查询id为1的user
	fmt.Println(user)

	db.First(&user, "name = ?", "L1215") // 查询code为l1212的user
	//user.Update()
	// 更新
	db.Model(&user).Update("Mobile", "1389342214")
	db.Model(&user).Updates(map[string]interface{}{"name": "L12155", "Address": "China-Zhejiang-HangZhou", "Mobile": "34512312345"})
}
