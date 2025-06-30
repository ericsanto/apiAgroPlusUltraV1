package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type AgricultureCultureBuilder struct{}

func NewAgricultureCultureBuiler() *AgricultureCultureBuilder {
	return &AgricultureCultureBuilder{}
}

func (acb *AgricultureCultureBuilder) Builder() controllers.AgricultureCultureControllerInterface {

	agricultureCultureRepository := repositories.NewAgricultureCultureRepository(db.DB)
	agricultureCultureService := services.NewAgricultureCultureService(agricultureCultureRepository)
	agricultureCultureHandler := controllers.NewAgricultureController(agricultureCultureService)

	return agricultureCultureHandler

}
