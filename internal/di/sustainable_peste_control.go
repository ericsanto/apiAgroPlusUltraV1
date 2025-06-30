package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type SustainablePestControlBuilder struct{}

func NewSustainablePestControl() *SustainablePestControlBuilder {
	return &SustainablePestControlBuilder{}
}

func (spcb *SustainablePestControlBuilder) Builder() controllers.SustainablePestControlControllerInterface {

	repositorySustainablePestControl := repositories.NewSustainablePestControlRepository(db.DB)
	serviceSustainablePestControl := services.NewSustainablePestControlService(repositorySustainablePestControl)
	controllerSustainablePestControl := controllers.NewSustainablePestControlController(serviceSustainablePestControl)

	return controllerSustainablePestControl

}
