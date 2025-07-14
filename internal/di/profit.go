package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type ProfitBuilder struct{}

func NewProfitBuilder() *ProfitBuilder {
	return &ProfitBuilder{}
}

func (pb *ProfitBuilder) Builder() controllers.ProfitControllerInterface {

	farmRepository := repositories.NewFarmRepository(db.DB)
	plantingRepository := repositories.NewPlantingRepository(db.DB, farmRepository)

	repository := repositories.NewProfitRepository(db.DB, plantingRepository)
	service := services.NewProfitService(repository)
	controller := controllers.NewProfitController(service)

	return controller
}
