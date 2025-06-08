package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/jsonutil"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/kafka"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/upload"
)

func DiseaseDetect(formFile upload.UploadFileInterface, formKey string) (map[string]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	minioClient, err := bucket.CreateMinioClient(ENDPOINT, ACCESS_KEY_MINIO, SECRET_KEY_MINIO, "", false)

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("erro ao se conectar com minio")
	}

	file, header, err := upload.UploadFile(formFile, formKey)
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

	typeDetect := "disease"

	message, err := kafka.SendAndReceiverKafkaService(ctx, urlImage, typeDetect)
	if err != nil {
		return nil, err
	}

	var responseApiPython map[string]string

	fmt.Println(message)

	err = jsonutil.ConvertStringToJson(message, &responseApiPython)
	log.Println(err)

	if err != nil {
		return nil, err
	}

	return responseApiPython, nil
}
