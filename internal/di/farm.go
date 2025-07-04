package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type FarmBuilder struct{}

func NewFarmBuilder() *FarmBuilder {
	return &FarmBuilder{}
}

func (fb *FarmBuilder) Builder() controllers.FarmControllerInterface {

	farmRepository := repositories.NewFarmRepository(db.DB)
	farmService := services.NewFarmService(farmRepository)
	farmController := controllers.NewFarmController(farmService)

	return farmController
}
