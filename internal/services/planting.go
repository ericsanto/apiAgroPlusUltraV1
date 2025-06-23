package services

import (
	"fmt"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type PlantingService struct {
	plantingRepository *repositories.PlantingRepository
}

func NewPlantingService(plantingRepository *repositories.PlantingRepository) *PlantingService {
	return &PlantingService{plantingRepository: plantingRepository}
}

func (p *PlantingService) GetByParam(batchID uint) (responses.PlantingResponse, error) {

	var plantingResponses responses.PlantingResponse

	planting, err := p.plantingRepository.FindByParamPlanting(batchID)

	if err != nil {
		return plantingResponses, fmt.Errorf("erro: %w", err)
	}

	plantingResponse := responses.PlantingResponse{
		ID:                   planting.ID,
		BatchID:              planting.BatchID,
		AgricultureCultureID: planting.AgricultureCultureID,
		IsPlanting:           planting.IsPlanting,
		StartDatePlanting:    planting.StartDatePlanting,
		SpaceBetweenPlants:   planting.SpaceBetweenPlants,
		SpaceBetweenRows:     planting.SpaceBetweenPlants,
		IrrigationTypeID:     planting.IrrigationTypeID,
	}

	return plantingResponse, nil

}

func (p *PlantingService) PostPlanting(requestPlanting requests.PlantingRequest) error {

	planting, _ := p.GetByParam(requestPlanting.BatchID)

	if planting.ID != 0 && planting.IsPlanting {
		return fmt.Errorf("erro ao cadastrar plantação. Lote já está sendo utilizado pela cultura %d", planting.AgricultureCultureID)
	}

	entityPlanting := entities.PlantingEntity{
		BatchID:              requestPlanting.BatchID,
		AgricultureCultureID: requestPlanting.AgricultureCultureID,
		StartDatePlanting:    time.Now(),
		IsPlanting:           *requestPlanting.IsPlanting,
		SpaceBetweenPlants:   requestPlanting.SpaceBetweenPlants,
		SpaceBetweenRows:     requestPlanting.SpaceBetweenRows,
		IrrigationTypeID:     requestPlanting.IrrigationTypeID,
	}

	if err := p.plantingRepository.CreatePlanting(entityPlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil

}

func (p *PlantingService) GetAllPlanting() ([]responses.PlantingResponse, error) {

	var responseListPlanting []responses.PlantingResponse

	plantingsEntity, err := p.plantingRepository.FindAllPlanting()
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
		}

		responseListPlanting = append(responseListPlanting, planting)
	}

	return responseListPlanting, nil
}

func (p *PlantingService) GetByParamBatchNameOrIsActivePlanting(batchName string, active bool) ([]responses.BatchPlantiesResponse, error) {

	plantingBatchResponse, err := p.plantingRepository.FindByParamBatchNameOrIsActivePlanting(batchName, active)
	if err != nil {
		return plantingBatchResponse, fmt.Errorf("erro: %w", err)
	}

	return plantingBatchResponse, nil
}

func (p *PlantingService) PutPlanting(id uint, requestPlanting requests.PlantingRequest) error {

	entityPlanting := entities.PlantingEntity{
		BatchID:              requestPlanting.BatchID,
		AgricultureCultureID: requestPlanting.AgricultureCultureID,
		IsPlanting:           *requestPlanting.IsPlanting,
		SpaceBetweenPlants:   requestPlanting.SpaceBetweenPlants,
		SpaceBetweenRows:     requestPlanting.SpaceBetweenRows,
		IrrigationTypeID:     requestPlanting.IrrigationTypeID,
	}

	if err := p.plantingRepository.UpdatePlanting(id, entityPlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *PlantingService) DeletePlanting(id uint) error {

	if err := p.plantingRepository.DeletePlanting(id); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
