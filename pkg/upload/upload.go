package upload

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) (multipart.File, *multipart.FileHeader, error) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, nil, fmt.Errorf("erro: %w", err)
	}

	return file, header, nil
}
