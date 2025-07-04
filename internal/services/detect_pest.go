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

type DetectPestImageServiceInterface interface {
	DetectPestImage(formFile upload.UploadFileInterface, formKey string) (map[string][]ResponseApiPython, error)
}

type DetectPestImageService struct {
	BucketClient  bucket.BucketClientInterface
	ImageValidate bucket.ImageValidateInterface
	KafkaClient   kafka.Messaging
}

func NewDetectPestImageService(bucketClient bucket.BucketClientInterface, imageValidate bucket.ImageValidateInterface, kafkaClient kafka.Messaging) DetectPestImageServiceInterface {
	return &DetectPestImageService{
		BucketClient:  bucketClient,
		ImageValidate: imageValidate,
		KafkaClient:   kafkaClient,
	}
}

func (dp *DetectPestImageService) DetectPestImage(formFile upload.UploadFileInterface, formKey string) (map[string][]ResponseApiPython, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	file, header, err := upload.UploadFile(formFile, formKey)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	if err := dp.ImageValidate.ValidateImageSizeAndType(header); err != nil {
		return nil, err
	}

	configsBucket := bucket.BucketConfig{
		Name:          os.Getenv("BUCKET_NAME"),
		ObjectLocking: false,
		Region:        "",
	}

	endpoint := os.Getenv("ENDPOINT")

	if err := dp.BucketClient.PutObject(ctx, configsBucket, header, file); err != nil {
		return nil, err
	}

	urlImage := fmt.Sprintf("http://%s/%s/%s", endpoint, configsBucket.Name, header.Filename)

	typeDetect := "pest"

	message, err := dp.KafkaClient.SendAndReceiverService(ctx, urlImage, typeDetect)
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
