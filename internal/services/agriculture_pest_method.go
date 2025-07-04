package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type AgricultureCulturePestMethodServiceInterface interface {
	PostAgricultureCulturePestMethod(requestAgricultureCulturePestMethod requests.AgricultureCulturePestMethodRequest) error
	GetAllAgricultureCulturePestMethod() ([]responses.AgricultureCulturePestMethodResponse, error)
	GetAllAgricultureCulturePestMethodByParam(cultureName, pestName, methodName interface{}) ([]responses.AgricultureCulturePestMethodResponse, error)
}

type AgricultureCulturePestMethodService struct {
	agricultureCulturePestMethdRepository repositories.AgricultureCulturePestMethodRepositoryInterface
}

func NewAgricultureCulturePestMethodService(agricultureCulturePestMethdRepository repositories.AgricultureCulturePestMethodRepositoryInterface) AgricultureCulturePestMethodServiceInterface {
	return &AgricultureCulturePestMethodService{agricultureCulturePestMethdRepository: agricultureCulturePestMethdRepository}
}

func (a *AgricultureCulturePestMethodService) PostAgricultureCulturePestMethod(requestAgricultureCulturePestMethod requests.AgricultureCulturePestMethodRequest) error {

	entityAgricultureCulturePestMethod := entities.AgricultureCulturePestMethodEntity{
		AgricultureCultureId:     requestAgricultureCulturePestMethod.AgricultureCultureId,
		PestId:                   requestAgricultureCulturePestMethod.PestId,
		SustainablePestControlId: requestAgricultureCulturePestMethod.SustainablePestControlId,
		Description:              requestAgricultureCulturePestMethod.Description,
	}

	if err := a.agricultureCulturePestMethdRepository.CreateAgricultureCulturePestMethod(entityAgricultureCulturePestMethod); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}

func (a *AgricultureCulturePestMethodService) GetAllAgricultureCulturePestMethod() ([]responses.AgricultureCulturePestMethodResponse, error) {

	result, err := a.agricultureCulturePestMethdRepository.FindAllAgricultureCulturePestMethod()
	if err != nil {
		return result, fmt.Errorf("erro: %v", err)
	}

	return result, nil

}

func (a *AgricultureCulturePestMethodService) GetAllAgricultureCulturePestMethodByParam(cultureName, pestName, methodName interface{}) ([]responses.AgricultureCulturePestMethodResponse, error) {

	responseAgricultureCulturePestMethod, err := a.agricultureCulturePestMethdRepository.FindByQueryParamAgricultureCulturePestMethod(cultureName, pestName, methodName)
	if err != nil {
		return responseAgricultureCulturePestMethod, fmt.Errorf("erro: %v", err)
	}

	return responseAgricultureCulturePestMethod, nil
}
