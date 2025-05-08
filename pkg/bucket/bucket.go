package bucket

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func BuckerIsExists(ctx context.Context, miniClient *minio.Client, bucketName string) (bool, error) {

	exist, errBucketExist := miniClient.BucketExists(ctx, bucketName)
	if errBucketExist != nil {
		return false, fmt.Errorf("erro ao verificar existência do bucket com o nome %s: %w", bucketName, errBucketExist)
	}

	return exist, nil
}

func CreateBucket(ctx context.Context, minioClient *minio.Client, bucketName string, region string, objectLookin bool) (bool, error) {

	exits, err := BuckerIsExists(ctx, minioClient, bucketName)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}

	if !exits {
		err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: region, ObjectLocking: objectLookin})
		return true, err
	}

	return true, nil

}
