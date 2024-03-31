package dao

import (
	"fmt"
	"go_practices/composited_exec/gin_gorm_jwt/global"
	"go_practices/composited_exec/gin_gorm_jwt/model"
)

//select一条记录
func SelectOneUser(userName string) (*model.User, error) {
	fields := []string{"userId", "username", "password"}
	userOne := &model.User{}
	err := global.DBLink.Select(fields).Where("username=?", userName).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		fmt.Println(userOne)
		return userOne, nil
	}
}
