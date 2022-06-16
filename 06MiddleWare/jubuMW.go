package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MWare() gin.HandlerFunc {

	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		context.Set("request", "中间件")
		// 执行函数
		context.Next()
		// 中间件执行完后续的一些事情
		status := context.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
func main() {
	r := gin.Default()

	r.GET("/hello", MWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  "200",
		})
	})
	r.Run()
}
