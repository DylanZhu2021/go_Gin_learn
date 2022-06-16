package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//this block:how to get arguments from post-form
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./html/*")
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/index", func(context *gin.Context) {
		types := context.DefaultPostForm("type", "post")
		username := context.PostForm("username")
		password := context.PostForm("password")

		context.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
			"types":    types,
		})
	})
	r.Run(":8080")
}
