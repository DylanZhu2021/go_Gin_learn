package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("v2")
	{
		v2.GET("/login", login)
		v2.GET("/submit", submit)
	}
	r.Run(":8080")

}

func login(context *gin.Context) {
	name := context.DefaultQuery("name", "ABC_login")
	context.String(http.StatusOK, "welcome %v", name)
}
func submit(context *gin.Context) {
	name := context.DefaultQuery("name", "ABC_submit")
	context.String(http.StatusOK, "welcome %v", name)
}
