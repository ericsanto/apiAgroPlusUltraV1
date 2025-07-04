package upload

import (
	"fmt"
	"mime/multipart"
)

type UploadFileInterface interface {
	FormFile(string) (multipart.File, *multipart.FileHeader, error)
}

type UploadFileSInterface interface {
	UploadFile(c UploadFileInterface) (multipart.File, *multipart.FileHeader, error)
}

type UploadFileS struct {
	key string
}

func NewUploadFileS(key string) UploadFileSInterface {
	return &UploadFileS{key: key}
}

func (uf *UploadFileS) UploadFile(c UploadFileInterface) (multipart.File, *multipart.FileHeader, error) {

	file, header, err := c.FormFile(uf.key)
	if err != nil {
		return nil, nil, fmt.Errorf("erro: %w", err)
	}

	return file, header, nil
}
