package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type SalePlantingService struct {
	salePlantingRepository *repositories.SalePlantingRepository
}

func NewSalePlantingService(salePlantingRepository *repositories.SalePlantingRepository) *SalePlantingService {
	return &SalePlantingService{salePlantingRepository: salePlantingRepository}
}

func (s *SalePlantingService) PostSalePlanting(requestSalePlanting requests.SalePlantingRequest) error {

	entitySalePlanting := entities.SalePlantingEntity{
		PlantingID: requestSalePlanting.PlantingID,
		ValueSale:  requestSalePlanting.ValueSale,
	}

	if err := s.salePlantingRepository.CreateSalePlantingRepository(entitySalePlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (s *SalePlantingService) GetAllSalePlanting() ([]responses.SalePlantingResponse, error) {

	var responsesSalePlanting []responses.SalePlantingResponse

	entitiesSalePlanting, err := s.salePlantingRepository.FindAllSalePlanting()
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	for _, v := range entitiesSalePlanting {
		responseSalePlanting := responses.SalePlantingResponse{
			ID:         v.ID,
			PlantingID: v.PlantingID,
			ValueSale:  v.ValueSale,
		}

		responsesSalePlanting = append(responsesSalePlanting, responseSalePlanting)
	}

	return responsesSalePlanting, nil
}

func (s *SalePlantingService) GetSalePlantingByID(id uint) (*responses.SalePlantingResponse, error) {

	entitySalePlanting, err := s.salePlantingRepository.FindSalePlantingByID(id)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	responseSalePlanting := responses.SalePlantingResponse{
		ID:         entitySalePlanting.ID,
		PlantingID: entitySalePlanting.PlantingID,
		ValueSale:  entitySalePlanting.ValueSale,
	}

	return &responseSalePlanting, nil
}

func (s *SalePlantingService) PutSalePlanting(id uint, requestSalePlanting requests.SalePlantingRequest) error {

	entitySalePlanting := entities.SalePlantingEntity{
		PlantingID: requestSalePlanting.PlantingID,
		ValueSale:  requestSalePlanting.ValueSale,
	}

	if err := s.salePlantingRepository.UpdateSalePlanting(id, entitySalePlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (s *SalePlantingService) DeleteSalePlanting(id uint) error {

	if err := s.salePlantingRepository.DeleteSalePlanting(id); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
