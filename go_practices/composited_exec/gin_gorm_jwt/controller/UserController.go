package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_practices/composited_exec/gin_gorm_jwt/pkg/jwt"
	"go_practices/composited_exec/gin_gorm_jwt/pkg/result"
	"go_practices/composited_exec/gin_gorm_jwt/pkg/validCheck"
	"go_practices/composited_exec/gin_gorm_jwt/request"
	"go_practices/composited_exec/gin_gorm_jwt/service"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

//用户登录
func (u *UserController) Login(c *gin.Context) {
	result := result.NewResult(c)

	param := request.LoginRequest{
		UserName: c.Param("username"),
		PassWord: c.Param("password"),
	}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		result.Error(400, errs.Error())
		return
	}

	userOne, err := service.GetOneUser(param.UserName)
	if err != nil {
		result.Error(404, "数据查询错误")
	} else {
		//password is right?
		err := bcrypt.CompareHashAndPassword([]byte(userOne.Password), []byte(param.PassWord))
		// 没有错误则密码匹配
		if err != nil {
			result.Error(1001, "账号信息错误")
		} else {
			tokenString, _ := jwt.GenToken(param.UserName)
			m := map[string]string{
				"tokenStr": tokenString,
			}
			result.Success(m)
		}
	}
	return
}

//用户信息info
func (u *UserController) Info(c *gin.Context) {
	username := c.MustGet("username").(string)
	fmt.Println("user login begin")
	result := result.NewResult(c)

	m := map[string]string{
		"username": username,
	}
	result.Success(m)

	return
}

//生成用户pass
func (u *UserController) Pass(c *gin.Context) {
	oPwd := c.Query("password")

	result := result.NewResult(c)

	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(oPwd), bcrypt.DefaultCost)
	m := map[string]string{
		"origin-password": oPwd,
		"crypt-password":  string(hashPwd),
	}
	result.Success(m)

	return
}
