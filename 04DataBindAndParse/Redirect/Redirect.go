package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/test1")
	})

	r.GET("/test1", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		r.HandleContext(context)
	})

	r.GET("/test2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  "hello",
			"message": "this is redirect",
		})
	})

	r.Run(":8080")

}
