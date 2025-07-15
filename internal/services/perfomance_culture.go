package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type PerformancePlantingServiceInterface interface {
	PostPerformancePlanting(batchID, farmID, userI, plantingID uint, requestPerformanceCulture requests.PerformancePlantingRequest) error
	GetAllPerformancePlanting(batchID, farmID, userID uint) ([]responses.PerformanceCultureResponse, error)
	PutPerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint, requestPerformanceEntity requests.PerformancePlantingRequest) error
	GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(batchID, farmID, userID, plantingID, performanceID uint) (*responses.PerformanceCultureResponse, error)
	DeletePerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint) error
}

type PerformancePlantingService struct {
	performanceCultureRepository repositories.PerformancePlantingRepositoryInterface
	plantingRepository           repositories.PlantingRepositoryInterface
}

func NewPerformancePlantingService(performanceCultureRepository repositories.PerformancePlantingRepositoryInterface,
	plantingRepository repositories.PlantingRepositoryInterface) PerformancePlantingServiceInterface {
	return &PerformancePlantingService{performanceCultureRepository: performanceCultureRepository,
		plantingRepository: plantingRepository}
}

func (p *PerformancePlantingService) PostPerformancePlanting(batchID, farmID, userID, plantingID uint, requestPerformanceCulture requests.PerformancePlantingRequest) error {

	if validateUnit := enums.IsValidateFieldUnitEnum(requestPerformanceCulture.UnitProductionObtained); !validateUnit {
		return fmt.Errorf("o campo unit_production_obtained %w", myerror.ErrEnumInvalid)
	}

	if validateUnit := enums.IsValidateFieldUnitEnum(requestPerformanceCulture.UnitHarvestedArea); !validateUnit {
		return fmt.Errorf("o campo unit_harvested_area  %w", myerror.ErrEnumInvalid)
	}

	_, err := p.plantingRepository.FindPlantingByID(batchID, farmID, userID, plantingID)

	if err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	entityPerformanceCulture := entities.PerformancePlantingEntity{
		PlantingID:             plantingID,
		ProductionObtained:     requestPerformanceCulture.ProductionObtained,
		UnitProductionObtained: requestPerformanceCulture.UnitProductionObtained,
		HarvestedArea:          requestPerformanceCulture.HarvestedArea,
		UnitHarvestedArea:      requestPerformanceCulture.UnitHarvestedArea,
		HarvestedDate:          requestPerformanceCulture.HarvestedDate,
	}

	if err := p.performanceCultureRepository.CreatePerformancePlanting(batchID, farmID, userID, plantingID, entityPerformanceCulture); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *PerformancePlantingService) GetAllPerformancePlanting(batchID, farmID, userID uint) ([]responses.PerformanceCultureResponse, error) {

	var reponsePerformancesCultures []responses.PerformanceCultureResponse

	dbResult, err := p.performanceCultureRepository.FindAll(batchID, farmID, userID)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	for _, v := range dbResult {
		responsePerformanceCulture := responses.PerformanceCultureResponse{
			Planting: responses.BatchPlantiesResponse{
				BatchName:              v.BatchName,
				AgricultureCultureName: v.AgricultureCultureName,
				StartDatePlanting:      v.StartDatePlanting,
				IsPlanting:             v.IsPlanting,
			},
			ID:                         v.ID,
			ProductionObtained:         v.ProductionObtained,
			ProductionObtainedFormated: v.ProductionObtainedFormated,
			HarvestedArea:              v.HarvestedArea,
			HarvestedAreaFormated:      v.HarvestedAreaFormated,
			HarvestedDate:              v.HarvestedDate,
		}

		reponsePerformancesCultures = append(reponsePerformancesCultures, responsePerformanceCulture)
	}

	return reponsePerformancesCultures, nil
}

func (p *PerformancePlantingService) PutPerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint, requestPerformanceEntity requests.PerformancePlantingRequest) error {

	entityPerformancePlanting := entities.PerformancePlantingEntity{
		ID:                     performanceID,
		PlantingID:             plantingID,
		ProductionObtained:     requestPerformanceEntity.ProductionObtained,
		UnitProductionObtained: requestPerformanceEntity.UnitProductionObtained,
		HarvestedArea:          requestPerformanceEntity.HarvestedArea,
		UnitHarvestedArea:      requestPerformanceEntity.UnitHarvestedArea,
		HarvestedDate:          requestPerformanceEntity.HarvestedDate,
	}

	if err := p.performanceCultureRepository.UpdatePerformancePlanting(batchID, farmID, userID, plantingID, performanceID, entityPerformancePlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *PerformancePlantingService) GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(batchID, farmID, userID, plantingID, performanceID uint) (*responses.PerformanceCultureResponse, error) {

	dBResultPerformancePlanting, err := p.performanceCultureRepository.FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID(batchID, farmID, userID, plantingID, performanceID)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	responsePerformancePlanting := responses.PerformanceCultureResponse{
		Planting: responses.BatchPlantiesResponse{
			BatchName:              dBResultPerformancePlanting.BatchName,
			AgricultureCultureName: dBResultPerformancePlanting.AgricultureCultureName,
			IsPlanting:             dBResultPerformancePlanting.IsPlanting,
			StartDatePlanting:      dBResultPerformancePlanting.StartDatePlanting,
		},

		ID:                         dBResultPerformancePlanting.ID,
		ProductionObtained:         dBResultPerformancePlanting.ProductionObtained,
		ProductionObtainedFormated: dBResultPerformancePlanting.ProductionObtainedFormated,
		HarvestedArea:              dBResultPerformancePlanting.HarvestedArea,
		HarvestedAreaFormated:      dBResultPerformancePlanting.HarvestedAreaFormated,
		HarvestedDate:              dBResultPerformancePlanting.HarvestedDate,
	}

	return &responsePerformancePlanting, nil
}

func (p *PerformancePlantingService) DeletePerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint) error {

	if err := p.performanceCultureRepository.DeletePerformancePlanting(batchID, farmID, userID, plantingID, performanceID); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
