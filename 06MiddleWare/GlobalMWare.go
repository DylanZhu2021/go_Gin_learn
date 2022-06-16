package main

/*
最后的get路由处理函数可以理解为最后的中间件，在不是调用c.Abort()的情况下，
所有的中间件都会被执行到。当某个中间件调用了c.Next(),则整个过程会产生嵌套关系。
如果某个中间件调用了c.Abort(),则此中间件结束后会直接返回，后面的中间件均不会调用。
*/
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MiddleWare1 中间件
func MiddleWare1() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("\n中间件1开始执行")
		context.Set("request", "MiddleW1")
		fmt.Println(context.Get("request"))
		status := context.Writer.Status()
		context.Next()
		fmt.Println("中间件1执行完成,", status)
		t2 := time.Since(t)
		fmt.Println("time", t2)
	}
}

func MiddleWare2() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("\n中间件2开始执行")
		context.Set("request", "MiddleW2")
		fmt.Println(context.Get("request"))
		status := context.Writer.Status()
		context.Next()
		//context.Abort()
		fmt.Println("中间件2执行完成,", status)
		t2 := time.Since(t)
		fmt.Println("time", t2)
	}
}

func MiddleWare3() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("\n中间件3开始执行")
		context.Set("request", "MiddleW3")
		status := context.Writer.Status()
		fmt.Println(context.Get("request"))
		//context.Abort()
		fmt.Println("中间件3执行完成,", status)
		t2 := time.Since(t)
		fmt.Println("time", t2)
	}
}
func main() {
	//1.创建路由
	//默认使用两个中间件:Logger,Recovery
	r := gin.Default()

	//2.使用中间件
	r.Use(MiddleWare1(), MiddleWare2(), MiddleWare3())

	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  "200",
		})
	})
	r.Run()
}
