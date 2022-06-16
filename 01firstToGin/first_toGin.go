package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则执行的函数
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello GIn")
	})
	//3.监听端口
	r.Run()
}
