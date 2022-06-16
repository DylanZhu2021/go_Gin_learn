package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//理解同步与异步的差别：
//1.同步：当服务器接收到多个请求时，服务器对每个请求独立处理，即等一个请求处理完后在处理下一个
//2.异步：当服务器接收到多个请求时，服务器对请求同时处理
//以此来实现高并发！！！！！
func main() {
	r := gin.Default()
	r.GET("/async", func(context *gin.Context) {
		newCon := context.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done!path:" + newCon.Request.URL.Path)
		}()
	})

	r.GET("/sync", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.JSON(http.StatusOK, gin.H{
			"message": "this is sync test!",
			"status":  "200",
		})
		log.Println("Done!path:" + context.Request.URL.Path)

	})

	r.Run(":8080")
}
