package upload

import (
	"fmt"
	"mime/multipart"
)

type UploadFileInterface interface {
	FormFile(string) (multipart.File, *multipart.FileHeader, error)
}

func UploadFile(c UploadFileInterface, key string) (multipart.File, *multipart.FileHeader, error) {

	file, header, err := c.FormFile(key)
	if err != nil {
		return nil, nil, fmt.Errorf("erro: %w", err)
	}

	return file, header, nil
}
