package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/jsonutil"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/kafka"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/upload"
	"github.com/gin-gonic/gin"
)

type ResponseApiPythonDisease struct {
	Disease string `json:"disease"`
}

type ResponseApiPythonError struct {
	Error string `json:"erro"`
}

func DiseaseDetect(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	minioClient, err := bucket.CreateMinioClient("minio:9000", "MJ4Ile4FCe3tpjgqgKVx", "4aXSs4tygx9A85GfOvQ32W9ls8BDXZRqFeBQxmFb", "", false)

	if err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "erro ao se conectar com o minio", c)
		return
	}

	file, header, err := upload.UploadFile(c.Request, "image")
	if err != nil {
		myerror.HttpErrors(http.StatusBadRequest, err.Error(), c)
		return
	}

	defer file.Close()

	if err := bucket.VerifyImageSize(header); err != nil {
		myerror.HttpErrors(http.StatusRequestEntityTooLarge, err.Error(), c)
		return
	}

	if err := bucket.VerifyImageType(header); err != nil {
		myerror.HttpErrors(http.StatusUnsupportedMediaType, err.Error(), c)
		return
	}

	bucketName := "images-analisys"
	region := ""
	objectLookin := false

	uploadIsSuccess, err := bucket.AsyncSendImageToBucket(ctx, minioClient, bucketName, file, *header, region, objectLookin)

	log.Println(uploadIsSuccess)

	if err != nil {

		if errors.Is(err, ctx.Err()) {
			myerror.HttpErrors(http.StatusRequestTimeout, "tempo excedido", c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	typeDetect := "disease"

	urlImage := fmt.Sprintf("http://%s/%s/%s", "minio:9000", bucketName, header.Filename)

	message, err := kafka.SendAndReceiver(ctx, urlImage, typeDetect)

	if err != nil {

		if errors.Is(err, ctx.Err()) {
			myerror.HttpErrors(http.StatusRequestTimeout, err.Error(), c)
			return
		}

		if errors.As(err, &sarama.ConsumerError{}) {
			log.Println(err.Error())
			myerror.HttpErrors(http.StatusInternalServerError, "erro ao consumir mensagem do kafka", c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro ao se comunicar com o servidor", c)
	}

	var responseApiPythonDisease ResponseApiPythonDisease

	err = jsonutil.ConvertStringToJson(message, &responseApiPythonDisease)

	if err != nil {
		log.Println(err)
		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, responseApiPythonDisease)
}
