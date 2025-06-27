package di

import (
	"fmt"
	"os"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/kafka"
)

type PestDetectImageBuilder struct {
}

func NewPestDetectImageBuilder() *PestDetectImageBuilder {
	return &PestDetectImageBuilder{}
}

func (p *PestDetectImageBuilder) Builder() (controllers.DetectPestImageControllerInterface, error) {
	configs := bucket.MinioConfig{
		AccessKey: os.Getenv("ACCESS_KEY_ID_MINIO"),
		SecretKey: os.Getenv("SECRET_KEY_MINIO"),
		Endpoint:  os.Getenv("ENDPOINT"),
		ID:        "",
		Secure:    false,
	}

	bucketClient, err := bucket.NewMinioClient(configs)

	if err != nil {
		return nil, fmt.Errorf("erro ao criar cliente bucket %w", err)
	}

	imageValidate := bucket.NewValidateImage()

	messagingService := kafka.NewKafka()

	serviceDetectPestImage := services.NewDetectPestImageService(bucketClient, imageValidate, messagingService)
	controllerDetectPestImage := controllers.NewDetectPestImageController(serviceDetectPestImage)

	return controllerDetectPestImage, nil
}
