package di

import (
	"os"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/bucket"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/kafka"
)

type DiseaseDetectBuilder struct{}

func NewDiseaseDetectBuilder() *DiseaseDetectBuilder {
	return &DiseaseDetectBuilder{}
}

func (ddb *DiseaseDetectBuilder) Builder() (controllers.DiseaseDetectControllerInterface, error) {

	imageValidade := bucket.NewValidateImage()
	kafkaClient := kafka.NewKafka()

	configs := bucket.MinioConfig{
		AccessKey: os.Getenv("ACCESS_KEY_ID_MINIO"),
		SecretKey: os.Getenv("SECRET_KEY_MINIO"),
		Endpoint:  os.Getenv("ENDPOINT"),
		ID:        "",
		Secure:    false,
	}

	bucketClient, err := bucket.NewMinioClient(configs)

	if err != nil {
		return nil, err
	}

	diseaseDetectService := services.NewDiseaseDetect(bucketClient, imageValidade, kafkaClient)

	diseaseDetectController := controllers.NewDiseaseDetectController(diseaseDetectService)

	return diseaseDetectController, nil
}
