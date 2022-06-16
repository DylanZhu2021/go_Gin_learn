package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Userinfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password"json:"password" binding:"required"`
}

//表单数据解析和绑定
func main() {
	r := gin.Default()

	r.POST("/login", func(context *gin.Context) {
		var form Userinfo

		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		err := context.Bind(&form)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if form.Username == "root" && form.Password == "admin" {
			context.JSON(http.StatusOK, gin.H{
				"status":   "welcome",
				"username": form.Username,
				"password": form.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "username or password error",
			})
		}
	})
	r.Run(":8080")
}
