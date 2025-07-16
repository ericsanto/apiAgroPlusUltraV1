package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type PlantingServiceInterface interface {
	GetByParam(userID, farmID, batchID uint) (*responses.PlantingResponse, error)
	PostPlanting(userID, farmID, batchID uint, requestPlanting requests.PlantingRequest) error
	GetByParamBatchNameOrIsActivePlanting(batchName string, active bool, userID, farmID uint) ([]responses.BatchPlantiesResponse, error)
	GetAllPlanting(batchID, farmID, userID uint) ([]responses.PlantingResponse, error)
	PutPlanting(batchID, farmID, userID, plantingID uint, requestPlanting requests.PlantingRequest) error
	DeletePlanting(batchID, farmID, userID, plantingID uint) error
}

type PlantingService struct {
	plantingRepository repositories.PlantingRepositoryInterface
}

func NewPlantingService(plantingRepository repositories.PlantingRepositoryInterface) PlantingServiceInterface {
	return &PlantingService{plantingRepository: plantingRepository}
}

func (p *PlantingService) GetByParam(userID, farmID, batchID uint) (*responses.PlantingResponse, error) {

	planting, err := p.plantingRepository.FindByParamPlanting(userID, farmID, batchID)

	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	plantingResponse := responses.PlantingResponse{
		ID:                   planting.ID,
		BatchID:              batchID,
		AgricultureCultureID: planting.AgricultureCultureID,
		IsPlanting:           planting.IsPlanting,
		StartDatePlanting:    planting.StartDatePlanting,
		SpaceBetweenPlants:   planting.SpaceBetweenPlants,
		SpaceBetweenRows:     planting.SpaceBetweenRows,
		IrrigationTypeID:     planting.IrrigationTypeID,
		ExpectedProduction:   planting.ExpectedProduction,
	}

	return &plantingResponse, nil

}

func (p *PlantingService) PostPlanting(userID, farmID, batchID uint, requestPlanting requests.PlantingRequest) error {

	planting, err := p.GetByParam(userID, farmID, batchID)

	if err != nil && !errors.Is(err, myerror.ErrNotFound) {
		return err
	}

	if planting != nil {
		return fmt.Errorf("erro ao cadastrar plantação. Lote já está sendo utilizado pela cultura %d", planting.AgricultureCultureID)
	}

	entityPlanting := entities.PlantingEntity{
		BatchID:              batchID,
		AgricultureCultureID: requestPlanting.AgricultureCultureID,
		StartDatePlanting:    time.Now(),
		IsPlanting:           *requestPlanting.IsPlanting,
		SpaceBetweenPlants:   requestPlanting.SpaceBetweenPlants,
		SpaceBetweenRows:     requestPlanting.SpaceBetweenRows,
		IrrigationTypeID:     requestPlanting.IrrigationTypeID,
		ExpectedProduction:   requestPlanting.ExpectedProduction,
	}

	if err := p.plantingRepository.CreatePlanting(entityPlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil

}

func (p *PlantingService) GetAllPlanting(batchID, farmID, userID uint) ([]responses.PlantingResponse, error) {

	var responseListPlanting []responses.PlantingResponse

	plantingsEntity, err := p.plantingRepository.FindAllPlanting(userID, farmID, batchID)
	if err != nil {
		return responseListPlanting, fmt.Errorf("erro: %w", err)
	}

	for _, v := range plantingsEntity {
		planting := responses.PlantingResponse{
			ID:                   v.ID,
			BatchID:              v.BatchID,
			AgricultureCultureID: v.AgricultureCultureID,
			IsPlanting:           v.IsPlanting,
			SpaceBetweenPlants:   v.SpaceBetweenPlants,
			SpaceBetweenRows:     v.SpaceBetweenRows,
			IrrigationTypeID:     v.IrrigationTypeID,
			ExpectedProduction:   v.ExpectedProduction,
		}

		responseListPlanting = append(responseListPlanting, planting)
	}

	return responseListPlanting, nil
}

func (p *PlantingService) GetByParamBatchNameOrIsActivePlanting(batchName string, active bool, userID, farmID uint) ([]responses.BatchPlantiesResponse, error) {

	plantingBatchResponse, err := p.plantingRepository.FindByParamBatchNameOrIsActivePlanting(batchName, active, userID, farmID)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	return plantingBatchResponse, nil
}

func (p *PlantingService) PutPlanting(batchID, farmID, userID, plantingID uint, requestPlanting requests.PlantingRequest) error {

	entityPlanting := entities.PlantingEntity{
		AgricultureCultureID: requestPlanting.AgricultureCultureID,
		IsPlanting:           *requestPlanting.IsPlanting,
		SpaceBetweenPlants:   requestPlanting.SpaceBetweenPlants,
		SpaceBetweenRows:     requestPlanting.SpaceBetweenRows,
		IrrigationTypeID:     requestPlanting.IrrigationTypeID,
		ExpectedProduction:   requestPlanting.ExpectedProduction,
	}

	if err := p.plantingRepository.UpdatePlanting(batchID, farmID, userID, plantingID, entityPlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *PlantingService) DeletePlanting(batchID, farmID, userID, plantingID uint) error {

	if err := p.plantingRepository.DeletePlanting(batchID, farmID, userID, plantingID); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
