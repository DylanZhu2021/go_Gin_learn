package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//模拟实现权限验证中间件
//有2个路由，login和home
//login用于设置cookie
//home是访问查看信息的请求
//在请求home之前，先跑中间件代码，检验是否存在cookie
//访问home，会显示错误，因为权限校验未通过

func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Cookie("cookie_name")
		if err == nil {
			if cookie == "123456" {

				context.Next()
				return
			}
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "error",
			})
			context.Abort()
		}

	}
}
func main() {
	r := gin.Default()

	r.GET("/login", func(context *gin.Context) {
		//设置cookie
		context.SetCookie("cookie_name", "123456",
			60, "/", "localhost",
			false, false)
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  "200",
		})
	})

	r.GET("/home", MiddleWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    "home",
		})
	})
	r.Run()
}
