package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type TypePestBuilder struct{}

func NewTypePestBuilder() *TypePestBuilder {
	return &TypePestBuilder{}
}

func (tpb *TypePestBuilder) Builder() controllers.TypePestControllerInterface {

	typePestRepository := repositories.NewTypePestRepository(config.DB)
	typePestService := services.NewTypePestService(typePestRepository)
	typeServiceController := controllers.NewTypePestController(typePestService)

	return typeServiceController
}
