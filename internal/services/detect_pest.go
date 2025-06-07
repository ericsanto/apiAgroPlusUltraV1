package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/jsonutil"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/kafka"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/upload"
)

type ResponseApiPython struct {
	Pest                  string  `json:"pest"`
	Confidence            float32 `json:"confidence"`
	HitPercentage         float32 `json:"hit_percentage"`
	HitPercentageFormated string  `json:"hit_percentage_formated"`
}

var (
	ENDPOINT         = os.Getenv("ENDPOINT")
	SECRET_KEY_MINIO = os.Getenv("SECRET_KEY_MINIO")
	ACCESS_KEY_MINIO = os.Getenv("ACCESS_KEY_ID_MINIO")
	bucketName       = os.Getenv("BUCKET_NAME")
	region           = ""
	objectLookin     = false
	typeDetect       = "pest"
)

func DetectPestImage(formFile upload.UploadFileInterface) (map[string][]ResponseApiPython, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	minioClient, err := bucket.CreateMinioClient(ENDPOINT, ACCESS_KEY_MINIO, SECRET_KEY_MINIO, "", false)

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("erro ao se conectar com minio")
	}

	file, header, err := upload.UploadFile(formFile, "image")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	if err := bucket.ValidateImageSizeAndType(header); err != nil {
		return nil, err
	}

	if err := bucket.UploadImageToBucket(ctx, minioClient, bucketName, file, header, region, objectLookin); err != nil {
		return nil, err
	}

	urlImage := fmt.Sprintf("http://%s/%s/%s", ENDPOINT, bucketName, header.Filename)

	message, err := kafka.SendAndReceiverKafkaService(ctx, urlImage, typeDetect)
	if err != nil {
		return nil, err
	}

	var responseApiPython map[string][]ResponseApiPython

	err = jsonutil.ConvertStringToJson(message, &responseApiPython)
	log.Println(err)

	if err != nil {
		return nil, err
	}

	return responseApiPython, nil
}
