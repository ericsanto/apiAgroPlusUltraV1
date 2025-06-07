package bucket

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/minio/minio-go/v7"
)

const (
	maxSizeFile = 1024 * 1024 * 20
)

func SyncSendImageToBucket(ctx context.Context, minioClient *minio.Client, bucketName string, image multipart.File, header multipart.FileHeader, region string, objectLookin bool) (bool, error) {

	IsCreatedBucket, err := CreateBucket(ctx, minioClient, bucketName, region, objectLookin)
	if err != nil {
		return false, fmt.Errorf("erro ao criar bucket: %w", err)
	}

	if !IsCreatedBucket {
		return false, fmt.Errorf("bucket não foi criado: %w", err)
	}

	info, err := minioClient.PutObject(ctx, bucketName, header.Filename, image, header.Size, minio.PutObjectOptions{ContentType: header.Header.Get("Content-Type")})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", header.Filename, info.Size)

	return true, nil
}

func AsyncSendImageToBucket(ctx context.Context, minioClient *minio.Client, bucketName string, image multipart.File, header multipart.FileHeader, region string, objectLookin bool) (bool, error) {

	type SendedImage struct {
		Success bool
		Err     error
	}

	resultSendedImage := make(chan SendedImage)

	go func() {

		_, err := SyncSendImageToBucket(ctx, minioClient, bucketName, image, header, region, objectLookin)

		resultSendedImage <- SendedImage{Success: err == nil, Err: err}
	}()

	select {
	case <-ctx.Done():
		return false, ctx.Err()

	case result := <-resultSendedImage:
		return result.Success, result.Err
	}

}

func VerifyImageSize(header *multipart.FileHeader) error {

	if header.Size > maxSizeFile {
		return myerror.ErrImageSizeToLarge
	}

	return nil

}

func VerifyImageType(header *multipart.FileHeader) error {

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

func ValidateImageSizeAndTypeImage(header *multipart.FileHeader) error {

	if err := VerifyImageSize(header); err != nil {
		return err
	}

	if err := VerifyImageType(header); err != nil {
		return err
	}

	return nil
}

func UploadImageToBucketService(ctx context.Context, minioClient *minio.Client, bucketName string, file multipart.File, header *multipart.FileHeader, region string, objectLookin bool) error {

	_, err := AsyncSendImageToBucket(ctx, minioClient, bucketName, file, *header, region, objectLookin)

	if err != nil {

		if errors.Is(err, ctx.Err()) {
			return fmt.Errorf("tempo excedido")
		}

		return err
	}

	return nil
}
