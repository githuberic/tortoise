package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// curl -X POST "http://127.0.0.1:8080/httpost01" -d '{"name":"lgq"}'
func main() {
	r := gin.Default()
	r.POST("/httpost01", func(c *gin.Context) {
		binBodys, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.String(http.StatusOK, string(binBodys))
	})
	r.Run(":8081")
}
