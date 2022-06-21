package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//Person ..
type Person struct {
	//不能为空并且大于10
	Age  int    `form:"age" binding:"required,gt=10,lt=100"`
	Name string `form:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/verify", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})
	r.Run()
}
