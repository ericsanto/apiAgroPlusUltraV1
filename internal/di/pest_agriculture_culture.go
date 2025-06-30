package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type PestAgricultureCultureBuilder struct{}

func NewPestAgricultureCultureBuilder() *PestAgricultureCultureBuilder {
	return &PestAgricultureCultureBuilder{}
}

func (pacb *PestAgricultureCultureBuilder) Builder() controllers.PestAgricultureCultureControllerInterface {

	pestAgricultureCultureRepository := repositories.NewPestAgricultureCultureRepository(db.DB)
	pestAgricultureCultureService := services.NewPestAgricultureCultureService(pestAgricultureCultureRepository)
	pestAgricultureCultureController := controllers.NewPestAgricultureCultureController(pestAgricultureCultureService)

	return pestAgricultureCultureController
}
