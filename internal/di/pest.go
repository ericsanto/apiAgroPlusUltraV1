package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type PestBuilder struct{}

func NewPestBuilder() *PestBuilder {
	return &PestBuilder{}
}

func (pb *PestBuilder) Builder() controllers.PestControllerInterface {

	pestRepository := repositories.NewPestRepository(db.DB)
	pestService := services.NewPestService(pestRepository)
	pestController := controllers.NewPestController(pestService)

	return pestController
}
