package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

//各种数据格式的响应
func main() {

	r := gin.Default()
	//1.json响应
	r.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "this is json response",
			"status":  "200",
		})
	})

	//2.结构体响应
	r.GET("/struct", func(context *gin.Context) {
		var student struct {
			name string
			age  int
		}
		student.age = 20
		student.name = "dylan"
		context.JSON(http.StatusOK, student)
	})

	//3.XML响应
	r.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"message": "this is a XML response",
			"status":  "200",
		})
	})

	//4.YAML响应
	r.GET("/yaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{
			"message": "this is a YAML response",
			"status":  "200",
		})
	})

	//5.Protobuf格式，谷歌开发的高效存储读取的工具

	r.GET("/protobuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		//定义数据
		label := "label"
		//传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		context.ProtoBuf(200, data)
	})

	r.Run(":8080")
}
