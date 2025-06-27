package bucket

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
)

type ProccessImageToBucket struct {
	MinioClient BucketClientInterface
}

func NewProccessImageToBucket(minioClient BucketClientInterface) *ProccessImageToBucket {
	return &ProccessImageToBucket{MinioClient: minioClient}
}

func (pib *ProccessImageToBucket) SyncSendImageToBucket(ctx context.Context, configsBucket BucketConfig, image multipart.File, header *multipart.FileHeader) (bool, error) {

	IsCreatedBucket, err := pib.MinioClient.CreateBucket(ctx, configsBucket)
	if err != nil {
		return false, fmt.Errorf("erro ao criar bucket: %w", err)
	}

	if !IsCreatedBucket {
		return false, fmt.Errorf("bucket não foi criado: %w", err)
	}

	if err := pib.MinioClient.PutObject(ctx, configsBucket, header, image); err != nil {
		return false, err
	}

	log.Printf("Successfully uploaded %s of size \n", header.Filename)

	return true, nil
}

func (pib *ProccessImageToBucket) AsyncSendImageToBucket(ctx context.Context, configsBucket BucketConfig, image multipart.File, header *multipart.FileHeader) (bool, error) {

	type SendedImage struct {
		Success bool
		Err     error
	}

	resultSendedImage := make(chan SendedImage)

	go func() {

		_, err := pib.SyncSendImageToBucket(ctx, configsBucket, image, header)

		resultSendedImage <- SendedImage{Success: err == nil, Err: err}
	}()

	select {
	case <-ctx.Done():
		return false, ctx.Err()

	case result := <-resultSendedImage:
		return result.Success, result.Err
	}

}

func (pib *ProccessImageToBucket) UploadImageToBucket(ctx context.Context, configsBucket BucketConfig, image multipart.File, header *multipart.FileHeader) error {

	_, err := pib.AsyncSendImageToBucket(ctx, configsBucket, image, header)

	if err != nil {

		if errors.Is(err, ctx.Err()) {
			return fmt.Errorf("tempo excedido")
		}

		return err
	}

	return nil
}
