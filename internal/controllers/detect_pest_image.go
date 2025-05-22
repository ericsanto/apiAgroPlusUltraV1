package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/jsonutil"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/kafka"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/upload"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	ENDPOINT = "minio:9000"
)

type ResponseApiPython struct {
	Pest                  string  `json:"pest"`
	Confidence            float32 `json:"confidence"`
	HitPercentage         float32 `json:"hit_percentage"`
	HitPercentageFormated string  `json:"hit_percentage_formated"`
}

func DetectPestImage(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	minioClient, err := minio.New(ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
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

	_, err = bucket.AsyncSendImageToBucket(ctx, minioClient, bucketName, file, *header, region, objectLookin)

	if err != nil {

		if errors.Is(err, ctx.Err()) {
			myerror.HttpErrors(http.StatusRequestTimeout, "tempo excedido", c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	urlImage := fmt.Sprintf("http://%s/%s/%s", ENDPOINT, bucketName, header.Filename)

	successSendedChannelMessage, messageKey, err := kafka.KafkaChannelMessage(ctx, urlImage)

	if !successSendedChannelMessage {
		myerror.HttpErrors(http.StatusRequestTimeout, "tempo excedido", c)
		log.Print(err)
		return
	}

	message, err := kafka.ConsumerMessageKafka(messageKey)
	if err != nil {
		myerror.HttpErrors(http.StatusServiceUnavailable, myerror.ErrStatusServiceUnavailable.Error(), c)
		return
	}

	var responseApiPython map[string][]ResponseApiPython

	err = jsonutil.ConvertStringToJson(message, &responseApiPython)

	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, responseApiPython)
}
