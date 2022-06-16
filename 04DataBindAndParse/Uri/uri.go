package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Userinfo struct {
	Username string `form:"username" uri:"username" json:"username" binding:"required"`
	Password string `form:"password" uri:"password" json:"password" binding:"required"`
}

//URI数据解析和绑定
func main() {
	r := gin.Default()
	r.GET("/login/:username/:password", func(context *gin.Context) {
		var uri Userinfo
		err := context.ShouldBindUri(&uri)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
		if uri.Username == "root" && uri.Password == "admin" {
			context.JSON(http.StatusOK, gin.H{
				"status":   "welcome",
				"username": uri.Username,
				"password": uri.Password,
			})
		}
	})
	r.Run(":8080")
}
