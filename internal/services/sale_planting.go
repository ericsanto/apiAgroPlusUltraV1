package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type SalePlantingServiceInterface interface {
	PostSalePlanting(batchID, farmID, userID, plantingID uint, requestSalePlanting requests.SalePlantingRequest) error
	GetAllSalePlanting(batchID, farmID, userID uint) ([]responses.SalePlantingResponse, error)
	GetSalePlantingByID(batchID, farmID, userID, salePlantingID uint) (*responses.SalePlantingResponse, error)
	PutSalePlanting(batchID, farmID, userID, salePlantingID uint, requestSalePlanting requests.SalePlantingRequest) error
	DeleteSalePlanting(batchID, farmID, userID, salePlantingID uint) error
}

type SalePlantingService struct {
	salePlantingRepository repositories.SalePlantingRepositoryInterface
	plantingService        PlantingServiceInterface
	batchService           BatchServiceInterface
	farmService            FarmServiceInterface
}

func NewSalePlantingService(salePlantingRepository repositories.SalePlantingRepositoryInterface,
	plantingService PlantingServiceInterface, batchService BatchServiceInterface,
	farmService FarmServiceInterface) SalePlantingServiceInterface {
	return &SalePlantingService{salePlantingRepository: salePlantingRepository, plantingService: plantingService,
		batchService: batchService, farmService: farmService}
}

func (s *SalePlantingService) PostSalePlanting(batchID, farmID, userID, plantingID uint, requestSalePlanting requests.SalePlantingRequest) error {

	_, err := s.farmService.GetFarmByID(userID, farmID)

	if err != nil {
		return err
	}

	_, err = s.batchService.GetBatchFindById(userID, farmID, batchID)

	if err != nil {
		return err
	}

	_, err = s.plantingService.GetByParam(userID, farmID, batchID)

	if err != nil {
		return err
	}

	entitySalePlanting := entities.SalePlantingEntity{
		PlantingID: plantingID,
		ValueSale:  requestSalePlanting.ValueSale,
	}

	if err := s.salePlantingRepository.CreateSalePlantingRepository(batchID, farmID, userID, plantingID, entitySalePlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (s *SalePlantingService) GetAllSalePlanting(batchID, farmID, userID uint) ([]responses.SalePlantingResponse, error) {

	var responsesSalePlanting []responses.SalePlantingResponse

	entitiesSalePlanting, err := s.salePlantingRepository.FindAllSalePlanting(batchID, farmID, userID)
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

func (s *SalePlantingService) GetSalePlantingByID(batchID, farmID, userID, salePlantingID uint) (*responses.SalePlantingResponse, error) {

	entitySalePlanting, err := s.salePlantingRepository.FindSalePlantingByID(batchID, farmID, userID, salePlantingID)
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

func (s *SalePlantingService) PutSalePlanting(batchID, farmID, userID, salePlantingID uint, requestSalePlanting requests.SalePlantingRequest) error {

	entitySalePlanting := entities.SalePlantingEntity{
		ValueSale: requestSalePlanting.ValueSale,
	}

	if err := s.salePlantingRepository.UpdateSalePlanting(batchID, farmID, userID, salePlantingID, entitySalePlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (s *SalePlantingService) DeleteSalePlanting(batchID, farmID, userID, salePlantingID uint) error {

	if err := s.salePlantingRepository.DeleteSalePlanting(batchID, farmID, userID, salePlantingID); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
