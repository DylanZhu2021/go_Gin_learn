package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//this block:how to upload files to server.
func main() {
	r := gin.Default()

	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.String(http.StatusInternalServerError, "upload file failed")
		}
		//save files to you server in a path
		context.SaveUploadedFile(file, "uploadFile/"+file.Filename)
		context.String(http.StatusOK, "upload file success,your file name is %v", file.Filename)
	})

	r.Run(":8080")
}
