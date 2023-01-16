package main

import (
	"fmt"
	"net/http"

	"github.com/Naithar01/go_gin_image-upload/img_upload_cloudinary"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB // Image Uplaod Setting

	img_upload_cloudinary.Init_Cloudinary()

	router.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		form, err := c.MultipartForm()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		files := form.File["files"]

		image_file_info_list := img_upload_cloudinary.Upload_Cloudinary_Image_Files(c, files)

		fmt.Println(image_file_info_list)

		c.JSON(http.StatusOK, image_file_info_list)
	})

	router.Run()
}
