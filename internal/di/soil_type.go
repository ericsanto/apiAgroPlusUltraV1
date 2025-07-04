package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type SoilTypeBuilder struct{}

func NewSoilTypeBuilder() *SoilTypeBuilder {
	return &SoilTypeBuilder{}
}

func (stb *SoilTypeBuilder) Builder() controllers.SoilTypeControollerInterface {

	typeSoilRepo := repositories.NewSoilRepository(db.DB)
	typeSoilService := services.NewSoilTypeService(typeSoilRepo)
	typeSoilController := controllers.NewSoilTypeController(typeSoilService)

	return typeSoilController
}
