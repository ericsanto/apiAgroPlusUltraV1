package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type AgricultureCultureIrrigationServiceInterface interface {
	GetAgricultureCultureIrrigationFindByID(culture_id uint) ([]responses.AgricultureCultureIrrigationResponse, error)
	PostAgricultureCultureIrrigation(requestAgricultureCultureIrrigation requests.AgricultureCultureIrrigationRequest) error
}

type AgricultureCultureIrrigationService struct {
	agricultureCultureIrrigationRepository *repositories.AgricultureCultureIrrigationRepository
}

func NewAgricultureCultureIrrigationService(agricultureCultureIrrigationRepository *repositories.AgricultureCultureIrrigationRepository) *AgricultureCultureIrrigationService {
	return &AgricultureCultureIrrigationService{agricultureCultureIrrigationRepository: agricultureCultureIrrigationRepository}
}

func (a *AgricultureCultureIrrigationService) GetAgricultureCultureIrrigationFindByID(culture_id uint) ([]responses.AgricultureCultureIrrigationResponse, error) {

	agriculturesCulturesResponse, err := a.agricultureCultureIrrigationRepository.FindByIdAgricultureCultureIrrigation(culture_id)
	if err != nil {
		return nil, fmt.Errorf("erro: %v", err)
	}

	return agriculturesCulturesResponse, nil
}

func (a *AgricultureCultureIrrigationService) PostAgricultureCultureIrrigation(requestAgricultureCultureIrrigation requests.AgricultureCultureIrrigationRequest) error {

	entityAgricultureCultureIrrigation := entities.AgricultureCultureIrrigation{
		AgricultureCultureId:   requestAgricultureCultureIrrigation.AgricultureCultureID,
		IrrigationRecomendedId: requestAgricultureCultureIrrigation.IrrigationRecomendedID,
	}

	if err := a.agricultureCultureIrrigationRepository.CreateAgricultureCultureIrrigation(entityAgricultureCultureIrrigation); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}

func (a *AgricultureCultureIrrigationService) PutAgricultureCultureIrrigation(cultureId, irrigationId uint, requestAgricultureCultureIrrigation requests.AgricultureCultureIrrigationRequest) error {

	entityAgricultureCultureIrrigation := entities.AgricultureCultureIrrigation{
		AgricultureCultureId:   requestAgricultureCultureIrrigation.AgricultureCultureID,
		IrrigationRecomendedId: requestAgricultureCultureIrrigation.IrrigationRecomendedID,
	}

	if err := a.agricultureCultureIrrigationRepository.UpdateAgricultureCultureIrrigation(cultureId, irrigationId, entityAgricultureCultureIrrigation); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}

func (a *AgricultureCultureIrrigationService) DeleteAgricultureCulturueIrrigation(cultureId, irrigationId uint) error {

	if err := a.agricultureCultureIrrigationRepository.DeleteAgricultureCulturueIrrigation(cultureId, irrigationId); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}
