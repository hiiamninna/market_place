package controller

import (
	"fmt"
	"market_place/collections"
	"market_place/library"
	"net/http"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const (
	maxFileSize = 1024 * 1024 * 2
	minFileSize = 1024 * 10
	JPG         = ".jpg"
	JPEG        = ".jpeg"
)

var allowedFormats = map[string]bool{
	JPG:  true,
	JPEG: true,
}

type Image struct {
	S3 library.S3
}

func NewImageController(s3 library.S3) Image {
	return Image{
		S3: s3,
	}
}

func (c Image) ImageUpload(ctx *fiber.Ctx) (int, string, interface{}, error) {

	var url string
	form, err := ctx.MultipartForm()
	if err != nil {
		return http.StatusInternalServerError, "upload image fail", nil, err
	}

	files := form.File["file"]

	for _, file := range files {
		// Check file size
		if file.Size < minFileSize {
			return http.StatusBadRequest, "file size must be at least 10 KB", nil, fmt.Errorf("file size is less 10 KB")
		}
		if file.Size > maxFileSize {
			return http.StatusBadRequest, "file size must be no more than 2 MB", nil, fmt.Errorf("file size is more than 2 MB")
		}

		// Check file format
		ext := filepath.Ext(file.Filename)
		if !allowedFormats[ext] {
			return http.StatusBadRequest, "file type must be JPG or JPEG", nil, fmt.Errorf("file type not JPG or JPEG")
		}

		name := generateUUID()

		uploadedFile, err := file.Open()
		if err != nil {
			return http.StatusInternalServerError, "upload image fail", nil, err
		}
		url, err = c.S3.UploadFile(uploadedFile, name)
		if err != nil {
			return http.StatusInternalServerError, "upload image fail", nil, err
		}
	}

	return http.StatusOK, "image uploaded successfully", collections.FileUpload{ImageUrl: url}, err
}
