package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type AgricultureCultureIrrigationBuilder struct{}

func NewAgricultureCultureIrrigationBuilder() *AgricultureCultureIrrigationBuilder {
	return &AgricultureCultureIrrigationBuilder{}
}

func (acib *AgricultureCultureIrrigationBuilder) Builder() controllers.AgricultureCultureIrrigationControllerInterface {

	agricultureCultureIrrigationRepository := repositories.NewAgricultureCultureIrrigationRepository(db.DB)
	agricultureCultureIrrigationService := services.NewAgricultureCultureIrrigationService(agricultureCultureIrrigationRepository)
	agricultureCultureIrrigationController := controllers.NewAgricultureCultureIrrigationController(agricultureCultureIrrigationService)

	return agricultureCultureIrrigationController
}
