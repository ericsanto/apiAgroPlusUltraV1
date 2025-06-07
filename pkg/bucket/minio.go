package bucket

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func CreateMinioClient(endpoint string, accessKey, secretKe, id string, secure bool) (*minio.Client, error) {

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKe, id),
		Secure: secure,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
