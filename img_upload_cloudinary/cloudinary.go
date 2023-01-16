package img_upload_cloudinary

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type Upload_Image struct {
	Filename string
	Path     string
}

// Code Example
// https://cloudinary.com/documentation/go_integration#configuration

// Google image Upload (s3?)
// https://articles.wesionary.team/golang-image-upload-with-google-cloud-storage-and-gin-part-2-99f4a642e06a

// Gin Single or multiple Image upload (body form)
// https://gin-gonic.com/ko-kr/docs/examples/upload-file/single-file/
// https://gin-gonic.com/ko-kr/docs/examples/upload-file/multiple-file/

// cloudinary & file variable
// https://www.topcoder.com/thrive/articles/uploading-files-using-golang-gin-and-cloudinary

// Fiber example
// https://www.youtube.com/watch?v=_0HfwZam220&t=484s

var (
	env_cloudinary_name       string = "naithar01"
	env_cloudinary_api_key    string = "732573379585195"
	env_cloudinary_api_secret string = "4sEBOzYQtWpYkzl9y9SBIZwpyx4"
)

var Cloudinary_Module *cloudinary.Cloudinary

func Init_Cloudinary() {
	Cloudinary_Module, _ = cloudinary.NewFromParams(env_cloudinary_name, env_cloudinary_api_key, env_cloudinary_api_secret)
	Cloudinary_Module.Config.URL.Secure = true
}

func Upload_Cloudinary_Image_Files(c *gin.Context, files []*multipart.FileHeader) []Upload_Image {
	upload_Img_list := []Upload_Image{}

	for _, file := range files {
		ctx := context.Background()

		open_file, _ := file.Open()

		upload_img, err := Cloudinary_Module.Upload.Upload(ctx, open_file, uploader.UploadParams{PublicID: file.Filename})

		image_file_type := Upload_Image{Filename: file.Filename, Path: upload_img.SecureURL}

		upload_Img_list = append(upload_Img_list, image_file_type)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		defer open_file.Close()
	}

	return upload_Img_list

}
