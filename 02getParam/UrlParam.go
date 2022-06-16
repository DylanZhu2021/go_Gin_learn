package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//this code block: how to get arguments from url
func main() {
	r := gin.Default()

	//method1:设置了一个默认值，如果没有这个参数那么就会使用默认值代替这个参数
	r.GET("user1", func(context *gin.Context) {
		name := context.DefaultQuery("name", "tom(default)")
		context.String(http.StatusOK, "you name is %v", name)
	})

	//method2:尝试获取参数,获取不到就是nil
	r.GET("user2", func(context *gin.Context) {
		name := context.Query("name")
		context.String(http.StatusOK, "you name is %v", name)
	})

	//method3:尝试获取参数，获取不到会返回一个false，然后另外做处理即可！
	r.GET("user3", func(context *gin.Context) {
		name, ok := context.GetQuery("user")
		if !ok {
			context.String(http.StatusOK, "please give your name!")
		} else {
			context.String(http.StatusOK, "you name is %v", name)
		}
	})
	r.Run(":8080")
}
