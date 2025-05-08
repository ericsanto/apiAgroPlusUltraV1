package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/kafka"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/upload"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func DetectPestImage(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	endpoint := "172.18.0.2:9000"
	acessaKeyId := "Drc6MYn9RSGvrhxGkzzd"
	secretKey := "QLSkP6aBlS3OF90ipLda4OreXfP5p6XWvkISyxB7"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(acessaKeyId, secretKey, ""),
		Secure: false,
	})

	if err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "erro ao se conectar com o minio", c)
		return
	}

	file, header, err := upload.UploadFile(c)
	if err != nil {
		myerror.HttpErrors(http.StatusBadRequest, err.Error(), c)
		return
	}

	defer file.Close()

	bucketName := "images-analisys"
	region := ""
	objectLookin := false

	resultChan := make(chan struct {
		success bool
		err     error
	})

	go func() {
		upload, err := bucket.SendImageToBucket(ctx, minioClient, bucketName, file, *header, region, objectLookin)

		resultChan <- struct {
			success bool
			err     error
		}{success: upload, err: err}
	}()

	select {
	case <-ctx.Done():
		myerror.HttpErrors(http.StatusRequestTimeout, "A requisição excedeu o limite de tempo", c)
		return

	case result := <-resultChan:
		if !result.success {
			myerror.HttpErrors(http.StatusBadRequest, result.err.Error(), c)
			return
		}
	}

	urlImageBucket := fmt.Sprintf("%s/%s/%s", endpoint, bucketName, header.Filename)

	if err := kafka.SendMessageKafka(urlImageBucket); err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	c.Status(http.StatusOK)
}
