package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type ProductionCostBuilder struct{}

func NewProductionCostBuilder() *ProductionCostBuilder {
	return &ProductionCostBuilder{}
}

func (pcb *ProductionCostBuilder) Builder() controllers.ProductionCostControllerInterface {

	farmRepository := repositories.NewFarmRepository(db.DB)
	farmService := services.NewFarmService(farmRepository)

	batchRepository := repositories.NewBatchRepository(db.DB, farmRepository)
	batchService := services.NewBatchService(batchRepository)

	plantingRepository := repositories.NewPlantingRepository(db.DB, farmRepository)
	plantingService := services.NewPlantingService(plantingRepository)

	productionCostRepository := repositories.NewProductionCostRepository(db.DB)
	productionCostService := services.NewProductionCostService(productionCostRepository, plantingService, batchService, farmService)

	productionCostController := controllers.NewProductionCostController(productionCostService)

	return productionCostController
}
