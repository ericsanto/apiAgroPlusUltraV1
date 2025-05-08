package bucket

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func SendImageToBucket(ctx context.Context, minioClient *minio.Client, bucketName string, image multipart.File, header multipart.FileHeader, region string, objectLookin bool) (bool, error) {

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
