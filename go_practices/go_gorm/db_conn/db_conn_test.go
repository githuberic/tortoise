package db_conn

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"testing"
)

/**
Verified
*/
func TestDBConnection(t *testing.T) {
	const (
		conn = "mysql://root:root_mysql@tcp(127.0.0.1:3307)/data_center?autocommit=true&charset=utf8"
	)
	DBEngine := strings.Replace(conn, "mysql://", "", -1)
	db, err := gorm.Open("mysql", DBEngine)
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Printf("successed in connecting databasen")
	}
	defer db.Close()
}
