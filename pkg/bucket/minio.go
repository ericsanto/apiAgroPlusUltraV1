package bucket

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket/config"
)

type BucketClientInterface interface {
	CreateBucket(ctx context.Context, configs config.BucketConfig) (bool, error)
	BucketIsExists(ctx context.Context, configs config.BucketConfig) (bool, error)
	PutObject(ctx context.Context, configsBucket config.BucketConfig, header *multipart.FileHeader, image multipart.File) error
}

type MinioConfig struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	ID        string
	Secure    bool
}

type MinioClient struct {
	client *minio.Client
}

func NewMinioClient(configs MinioConfig) (*MinioClient, error) {
	minioClient, err := minio.New(configs.Endpoint, &minio.Options{Creds: credentials.NewStaticV4(configs.AccessKey, configs.SecretKey, configs.ID),
		Secure: configs.Secure})

	if err != nil {
		return nil, err
	}

	return &MinioClient{client: minioClient}, nil
}

func (mc *MinioClient) BucketIsExists(ctx context.Context, configsBucket config.BucketConfig) (bool, error) {

	exist, errBucketExist := mc.client.BucketExists(ctx, configsBucket.Name)
	if errBucketExist != nil {
		return false, fmt.Errorf("erro ao verificar existência do bucket com o nome %s: %w", configsBucket.Name, errBucketExist)
	}

	return exist, nil
}

func (mc *MinioClient) CreateBucket(ctx context.Context, configsBucket config.BucketConfig) (bool, error) {

	exits, err := mc.BucketIsExists(ctx, configsBucket)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}

	if !exits {
		err := mc.client.MakeBucket(ctx, configsBucket.Name, minio.MakeBucketOptions{Region: configsBucket.Region,
			ObjectLocking: configsBucket.ObjectLocking})
		return true, err
	}

	return true, nil

}

func (mc *MinioClient) PutObject(ctx context.Context, configsBucket config.BucketConfig, header *multipart.FileHeader, image multipart.File) error {

	_, err := mc.client.PutObject(ctx, configsBucket.Name, header.Filename, image, header.Size, minio.PutObjectOptions{ContentType: header.Header.Get("Content-Type")})
	if err != nil {
		return fmt.Errorf("erro ao fazer put da imagem para o bucket %w", err)
	}

	return nil
}
