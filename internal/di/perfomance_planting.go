package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type PerformancePlantingBuilder struct{}

func NewPerfomancePlantingBuilder() *PerformancePlantingBuilder {
	return &PerformancePlantingBuilder{}
}

func (ppb *PerformancePlantingBuilder) Builder() controllers.PerformancePlantingControllerInterface {
	repositoryPerformanceCulture := repositories.NewPerformanceCultureRepository(db.DB)
	servicePerformanceCulture := services.NewPerformancePlantingService(repositoryPerformanceCulture)
	controllerPerformanceCulture := controllers.NewPerformancePlantingController(servicePerformanceCulture)

	return controllerPerformanceCulture
}
