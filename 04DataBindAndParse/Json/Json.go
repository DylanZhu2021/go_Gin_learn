package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Userinfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password"json:"password" binding:"required"`
}

//客户端传参，后端接收并解析到结构体，返回给客户端对应的输出
func main() {
	r := gin.Default()

	r.POST("/login", func(context *gin.Context) {
		var json Userinfo

		err := context.ShouldBindJSON(&json)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		if json.Username == "root" && json.Password == "admin" {
			context.JSON(http.StatusOK, gin.H{
				"status":   200,
				"username": json.Username,
				"password": json.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "username or password error",
			})
		}
	})

	//使用下面的curl语句
	//http://127.0.0.1:8080/login -H "Content-Type: application/json"
	//-d {\"username\":\"root\",\"password\":\"admin\"} -X POST
	r.Run(":8080")
}
