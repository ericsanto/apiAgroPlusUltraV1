package bucket

import (
	"mime/multipart"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

const (
	maxSizeFile = 1024 * 1024 * 20
)

type ImageValidateInterface interface {
	VerifyImageSize(header *multipart.FileHeader) error
	VerifyImageType(header *multipart.FileHeader) error
	ValidateImageSizeAndType(header *multipart.FileHeader) error
}

type ValidateImage struct {
}

func NewValidateImage() ImageValidateInterface {
	return &ValidateImage{}
}

func (vi *ValidateImage) VerifyImageSize(header *multipart.FileHeader) error {

	if header.Size > maxSizeFile {
		return myerror.ErrImageSizeToLarge
	}

	return nil

}

func (vi *ValidateImage) VerifyImageType(header *multipart.FileHeader) error {

	allowedTypesImage := map[string]bool{
		"image/png":  true,
		"image/jpeg": true,
		"image/jpg":  true,
	}

	if !allowedTypesImage[header.Header.Get("Content-type")] {
		return myerror.ErrUnsupportedImageType
	}

	return nil
}

func (vi *ValidateImage) ValidateImageSizeAndType(header *multipart.FileHeader) error {

	if err := vi.VerifyImageSize(header); err != nil {
		return err
	}

	if err := vi.VerifyImageType(header); err != nil {
		return err
	}

	return nil
}
