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

	productionCostRepository := repositories.NewProductionCostRepository(db.DB)
	productionCostService := services.NewProductionCostService(productionCostRepository)
	productionCostController := controllers.NewProductionCostController(productionCostService)

	return productionCostController
}
