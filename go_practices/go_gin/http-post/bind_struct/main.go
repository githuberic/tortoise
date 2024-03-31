package main

import (
	"github.com/gin-gonic/gin"
	"go_practices/go_gin/http-post/json_time"
)

type Person struct {
	Name     string             `form:"name"`
	Age      int                `form:"age" binding:"required,gt=10"`
	Address  string             `form:"address"`
	Birthday json_time.JSONTime `form:"birthday" time_format:"2006-01-02"`
}

/**
curl -X POST -H "Content-Type:application/json" -d'{"name":"lgq","age":15,"address":"杭州","birthday":"2008-12-09"}' "http://127.0.0.1:8081/httpStruct"
*/
func main() {
	r := gin.Default()
	r.POST("/httpStruct", bindStruct)
	r.Run(":8081")
}

func bindStruct(c *gin.Context) {
	var p Person
	if err := c.ShouldBind(&p); err != nil {
		c.String(200, "person bind error : %v", err)
	} else {
		c.String(200, "%v", p)
	}
}
