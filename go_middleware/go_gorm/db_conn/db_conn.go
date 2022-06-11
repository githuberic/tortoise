package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

func main() {
	const (
		conn = "mysql://root:root_mysql@tcp(127.0.0.1:3307)/db_user?autocommit=true&charset=utf8"
	)
	DBEngine := strings.Replace(conn, "mysql://", "", -1)
	db, err := gorm.Open("mysql", DBEngine)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	fmt.Printf("successed in connecting databasen")
}
