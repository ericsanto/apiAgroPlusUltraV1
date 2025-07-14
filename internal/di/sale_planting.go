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

	farmRepository := repositories.NewFarmRepository(db.DB)
	farmService := services.NewFarmService(farmRepository)

	batchRepository := repositories.NewBatchRepository(db.DB, farmRepository)
	batchService := services.NewBatchService(batchRepository)

	plantingRepository := repositories.NewPlantingRepository(db.DB, farmRepository)
	plantingService := services.NewPlantingService(plantingRepository)

	salePlantingRepository := repositories.NewSalePlantingRepository(db.DB)
	salePlantingService := services.NewSalePlantingService(salePlantingRepository, plantingService, batchService, farmService)
	salePlantingController := controllers.NewSalePlantingController(salePlantingService)

	return salePlantingController
}
