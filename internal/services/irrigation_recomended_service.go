package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type IrrigatiionRecomendedServiceInterface interface {
	GetAllIrrigationRecomended() ([]responses.IrrigationRecomendedResponse, error)
	PostIrrigationRecomended(requestIrrigationRecomended requests.IrrigationRecomendedRequest) error
	GetByIdIrrigationRecomended(id uint) (responses.IrrigationRecomendedResponse, error)
	PutIrrigationRecomended(id uint, requestIrrigationRecomended requests.IrrigationRecomendedRequest) error
	DeleteIrrigationRecomended(id uint) error
}

type IrrigationRecomendedService struct {
	irrigationRecomendedRepository *repositories.IrrigationRecomendedRepository
}

func NewIrrigationRecomendedService(irrigationRecomendedRepostitory *repositories.IrrigationRecomendedRepository) *IrrigationRecomendedService {
	return &IrrigationRecomendedService{irrigationRecomendedRepository: irrigationRecomendedRepostitory}
}

func (i *IrrigationRecomendedService) GetAllIrrigationRecomended() ([]responses.IrrigationRecomendedResponse, error) {

	var responsesIrrigationsRecomendeds []responses.IrrigationRecomendedResponse

	entityIrrigationRecomended, err := i.irrigationRecomendedRepository.FindAllIrrigationRecomended()
	if err != nil {
		return nil, fmt.Errorf("erro: %v", err)
	}

	for _, v := range entityIrrigationRecomended {
		responseIrrigationRecomended := responses.IrrigationRecomendedResponse{
			Id:                v.Id,
			IrrigationMax:     v.IrrigationMax,
			IrrigationMin:     v.IrrigationMin,
			PhenologicalPhase: v.PhenologicalPhase,
			PhaseDurationDays: v.PhaseDurationDays,
			Description:       v.Description,
			Unit:              v.Unit,
		}

		responsesIrrigationsRecomendeds = append(responsesIrrigationsRecomendeds, responseIrrigationRecomended)
	}

	return responsesIrrigationsRecomendeds, nil
}

func (i *IrrigationRecomendedService) PostIrrigationRecomended(requestIrrigationRecomended requests.IrrigationRecomendedRequest) error {

	entityIrrigationRecomended := entities.IrrigationRecomendedEntity{
		PhenologicalPhase: requestIrrigationRecomended.PhenologicalPhase,
		PhaseDurationDays: requestIrrigationRecomended.PhaseDurationDays,
		IrrigationMax:     requestIrrigationRecomended.IrrigationMax,
		IrrigationMin:     requestIrrigationRecomended.IrrigationMin,
		Description:       requestIrrigationRecomended.Description,
		Unit:              requestIrrigationRecomended.Unit,
	}

	if err := i.irrigationRecomendedRepository.CreateIrrigationRecomended(entityIrrigationRecomended); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil

}

func (i *IrrigationRecomendedService) GetByIdIrrigationRecomended(id uint) (*responses.IrrigationRecomendedResponse, error) {

	entityIrrigationRecomended, err := i.irrigationRecomendedRepository.FindByIdIrrigationRecomended(id)
	if err != nil {
		return nil, fmt.Errorf("erro: %v", err)
	}

	responseIrrigationRecomended := responses.IrrigationRecomendedResponse{
		Id:                entityIrrigationRecomended.Id,
		PhenologicalPhase: entityIrrigationRecomended.PhenologicalPhase,
		PhaseDurationDays: entityIrrigationRecomended.PhaseDurationDays,
		IrrigationMax:     entityIrrigationRecomended.IrrigationMax,
		IrrigationMin:     entityIrrigationRecomended.IrrigationMin,
		Description:       entityIrrigationRecomended.Description,
		Unit:              entityIrrigationRecomended.Unit,
	}

	return &responseIrrigationRecomended, nil
}

func (i *IrrigationRecomendedService) PutIrrigationRecomended(id uint, requestIrrigationRecomended requests.IrrigationRecomendedRequest) error {

	entityIrrigationRecomended := entities.IrrigationRecomendedEntity{
		PhenologicalPhase: requestIrrigationRecomended.PhenologicalPhase,
		PhaseDurationDays: requestIrrigationRecomended.PhaseDurationDays,
		IrrigationMax:     requestIrrigationRecomended.IrrigationMax,
		IrrigationMin:     requestIrrigationRecomended.IrrigationMin,
		Unit:              requestIrrigationRecomended.Unit,
		Description:       requestIrrigationRecomended.Description,
	}

	if err := i.irrigationRecomendedRepository.UpdateIrrigationRecomended(id, entityIrrigationRecomended); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}

func (i *IrrigationRecomendedService) DeleteIrrigationRecomended(id uint) error {

	if err := i.irrigationRecomendedRepository.DeleteIrrigationRecomended(id); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}
