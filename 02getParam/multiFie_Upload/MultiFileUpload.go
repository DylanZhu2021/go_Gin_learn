package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusBadRequest, "get form failed")
		}
		files := form.File["files"]

		for _, file := range files {
			err := context.SaveUploadedFile(file, "upload/"+file.Filename)
			if err != nil {
				context.String(http.StatusBadRequest, "save file: %v failed", file.Filename)
				return
			}
		}
		context.String(http.StatusOK, "upload %v files", len(files))
	})
	r.Run(":8080")
}
