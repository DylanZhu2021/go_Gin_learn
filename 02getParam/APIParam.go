package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	name   string
	action string
)

//this code block:get API arguments from function:context.Param()
//usage: /user/username(you input)/action(you input)
func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name = context.Param("name")
		action = context.Param("action")

		fmt.Println(name, " ", action)

		action = strings.Trim(action, "/")
		context.String(http.StatusOK, getFormat(name, action))
	})
	r.Run(":8080")
}

//return format string to function:context.String()
func getFormat(name string, action string) string {
	return "your name is " + name + "\n" + "you action is " + action
}
