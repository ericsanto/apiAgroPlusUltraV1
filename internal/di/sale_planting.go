package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type SalePlantingBuilder struct{}

func NewSalePlantingBuilder() *SalePlantingBuilder {
	return &SalePlantingBuilder{}
}

func (spb *SalePlantingBuilder) Builder() controllers.SalePlantingControllerInterface {
	salePlantingRepository := repositories.NewSalePlantingRepository(db.DB)
	salePlantingService := services.NewSalePlantingService(salePlantingRepository)
	salePlantingController := controllers.NewSalePlantingController(salePlantingService)

	return salePlantingController
}
