package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(context *gin.Context) {
		cookie, err := context.Cookie("cookie_name")
		if err != nil {
			//cookie = "noset"
			context.SetCookie("cookie_name",
				"value", 60, "/",
				"localhost", false, false)

		}
		fmt.Println("cookie: ", cookie)
	})
	r.Run()
}
