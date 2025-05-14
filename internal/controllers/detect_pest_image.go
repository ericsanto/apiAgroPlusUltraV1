package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/kafka"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/upload"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	SECRET_KEY_MINIO = os.Getenv("SECRET_KEY_MINIO")
	ACCESS_KEY_MINIO = os.Getenv("ACESS_KEY_ID_MINIO")
	ENDPOINT         = "minio:9000"
)

func DetectPestImage(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	minioClient, err := minio.New(ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(ACCESS_KEY_MINIO, SECRET_KEY_MINIO, ""),
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

	upload, err := bucket.AsyncSendImageToBucket(ctx, minioClient, bucketName, file, *header, region, objectLookin)

	log.Println(upload)

	if err != nil {

		if errors.Is(err, ctx.Err()) {
			myerror.HttpErrors(http.StatusRequestTimeout, "tempo excedido", c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	urlImage := fmt.Sprintf("http://%s/%s/%s", ENDPOINT, bucketName, header.Filename)

	successSendedChannelMessage, err := kafka.KafkaChannelMessage(ctx, urlImage)

	if !successSendedChannelMessage {
		myerror.HttpErrors(http.StatusRequestTimeout, "tempo excedido", c)
		log.Print(err)
		return
	}

	c.Status(http.StatusOK)
}
