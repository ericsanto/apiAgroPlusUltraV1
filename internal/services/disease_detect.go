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

type DiseaseDetectServiceInterface interface {
	ServiceDiseaseDetect(formFile upload.UploadFileInterface, formKey string) (map[string]string, error)
}

type DiseaseDetectService struct {
	bucketClient  bucket.BucketClientInterface
	imageValidate bucket.ImageValidateInterface
	kafkaClient   kafka.Messaging
}

func NewDiseaseDetect(bucketClient bucket.BucketClientInterface, imageValidate bucket.ImageValidateInterface, kafkaClient kafka.Messaging) DiseaseDetectServiceInterface {
	return &DiseaseDetectService{
		bucketClient:  bucketClient,
		imageValidate: imageValidate,
		kafkaClient:   kafkaClient,
	}
}

func (dds *DiseaseDetectService) ServiceDiseaseDetect(formFile upload.UploadFileInterface, formKey string) (map[string]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	file, header, err := upload.UploadFile(formFile, formKey)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	if err := dds.imageValidate.ValidateImageSizeAndType(header); err != nil {
		return nil, err
	}

	configsBucket := bucket.BucketConfig{
		Name:          os.Getenv("BUCKET_NAME"),
		ObjectLocking: false,
		Region:        "",
	}

	if err := dds.bucketClient.PutObject(ctx, configsBucket, header, file); err != nil {
		return nil, err
	}

	endpoint := os.Getenv("ENDPOINT")

	urlImage := fmt.Sprintf("http://%s/%s/%s", endpoint, configsBucket.Name, header.Filename)

	typeDetect := "disease"

	message, err := dds.kafkaClient.SendAndReceiverService(ctx, urlImage, typeDetect)
	if err != nil {
		return nil, err
	}

	var responseApiPython map[string]string

	err = jsonutil.ConvertStringToJson(message, &responseApiPython)
	log.Println(err)

	if err != nil {
		return nil, err
	}

	return responseApiPython, nil
}
