package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type IrrigationTypeBuilder struct{}

func NewIrrigationTypeBuilder() *IrrigationTypeBuilder {
	return &IrrigationTypeBuilder{}
}

func (itb *IrrigationTypeBuilder) Builder() controllers.IrrigationTypeControllerInterface {

	irrigationTypeRepository := repositories.NewIrrigationTypeRepository(db.DB)
	irrigationTypeService := services.NewIrrigationTypeService(irrigationTypeRepository)
	irrigatinTypeController := controllers.NewIrrigationTypeController(irrigationTypeService)

	return irrigatinTypeController
}
