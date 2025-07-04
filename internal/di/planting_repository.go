package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type PlantingBuilder struct{}

func NewPlantingRepository() *PlantingBuilder {
	return &PlantingBuilder{}
}

func (prb *PlantingBuilder) Builder() controllers.PlantingControllerInterface {

	plantingRepository := repositories.NewPlantingRepository(db.DB)
	plantingService := services.NewPlantingService(plantingRepository)
	plantingController := controllers.NewPlantingController(plantingService)

	return plantingController
}
