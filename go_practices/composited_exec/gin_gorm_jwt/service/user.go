package service

import (
	"go_practices/composited_exec/gin_gorm_jwt/dao"
	"go_practices/composited_exec/gin_gorm_jwt/model"
)

func GetOneUser(userName string) (*model.User, error) {
	return dao.SelectOneUser(userName)
}
