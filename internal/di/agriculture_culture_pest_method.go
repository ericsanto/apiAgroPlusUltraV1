package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type AgricultureCulturePestMethodBuilder struct{}

func NewAgricultureCulturePestMethodBuilder() *AgricultureCulturePestMethodBuilder {
	return &AgricultureCulturePestMethodBuilder{}
}

func (acpmb *AgricultureCulturePestMethodBuilder) Builder() controllers.AgricultureCulturePestMethodControllerInterface {

	agricultureCulturePestMethodRepository := repositories.NewAgricultureCulturePestMethodRepository(db.DB)
	agricultureCulturePestMethodService := services.NewAgricultureCulturePestMethodService(agricultureCulturePestMethodRepository)
	agricultureCulturePestMethodController := controllers.NewAgricultureCulturePestMethodController(agricultureCulturePestMethodService)

	return agricultureCulturePestMethodController
}
