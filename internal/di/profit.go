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

	repository := repositories.NewProfitRepository(db.DB)
	service := services.NewProfitService(repository)
	controller := controllers.NewProfitController(service)

	return controller
}
