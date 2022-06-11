package main

import (
	"fmt"
	"go_practices/go_gorm/exec_01/db"
)

/*
type User struct {
	Id      int    `gorm:"primary_key" json:"id"`
	Name    string `gorm:"not null;"`
	Address string `gorm:"not null;"`
	Mobile  string `gorm:"mobile"`
}*/

type User struct {
	Id      int    `gorm:"primary_key" json:"id"`
	Name    string `json:"name"`
	Address string `json:"string"`
	Mobile  string `json:"mobile"` //1:男、2:女
}

//添加数据
func (user *User) Add() {
	conn := db.GetDb()
	defer conn.Close()

	err := conn.Create(user).Error
	if err != nil {
		fmt.Printf("创建失败 %v", err)
	}
}

//修改数据
func (user *User) Update() {
	conn := db.GetDb()
	defer conn.Close()

	err := conn.Model(user).Update(user).Error
	if err != nil {
		fmt.Println("修改失败")
	}
}

//删除数据
func (user *User) Del() {
	conn := db.GetDb()
	defer conn.Close()

	err := conn.Delete(user).Error
	if err != nil {
		fmt.Println("删除失败")
	}
}
