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

type PerfomancePlantingService struct {
	perfomanceCultureRepository *repositories.PerfomancePlantingRepository
}

func NewPerfomancePlantingService(perfomanceCultureRepository *repositories.PerfomancePlantingRepository) *PerfomancePlantingService {
	return &PerfomancePlantingService{perfomanceCultureRepository: perfomanceCultureRepository}
}

func (p *PerfomancePlantingService) PostPerfomancePlanting(requestPerfomanceCulture requests.PerfomancePlantingRequest) error {

	if validateUnit := enums.IsValidateFieldUnitEnum(requestPerfomanceCulture.UnitProductionObtained); !validateUnit {
		return fmt.Errorf("o campo unit_production_obtained %w", myerror.ErrEnumInvalid)
	}

	if validateUnit := enums.IsValidateFieldUnitEnum(requestPerfomanceCulture.UnitHarvestedArea); !validateUnit {
		return fmt.Errorf("o campo unit_harvested_area  %w", myerror.ErrEnumInvalid)
	}

	entityPerfomanceCulture := entities.PerfomancePlantingEntity{
		PlantingID:             requestPerfomanceCulture.PlantingID,
		ProductionObtained:     requestPerfomanceCulture.ProductionObtained,
		UnitProductionObtained: requestPerfomanceCulture.UnitProductionObtained,
		HarvestedArea:          requestPerfomanceCulture.HarvestedArea,
		UnitHarvestedArea:      requestPerfomanceCulture.UnitHarvestedArea,
		HarvestedDate:          requestPerfomanceCulture.HarvestedDate,
	}

	if err := p.perfomanceCultureRepository.CreatePerfomancePlanting(entityPerfomanceCulture); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *PerfomancePlantingService) GetAllPerfomancePlanting() ([]responses.PerfomanceCultureResponse, error) {

	var reponsePerformancesCultures []responses.PerfomanceCultureResponse

	dbResult, err := p.perfomanceCultureRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	for _, v := range dbResult {
		responsePerfomanceCulture := responses.PerfomanceCultureResponse{
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

		reponsePerformancesCultures = append(reponsePerformancesCultures, responsePerfomanceCulture)
	}

	return reponsePerformancesCultures, nil
}

func (p *PerfomancePlantingService) PutPerformancePlanting(id uint, requestPerfomanceEntity requests.PerfomancePlantingRequest) error {

	entityPerfomancePlanting := entities.PerfomancePlantingEntity{
		PlantingID:             requestPerfomanceEntity.PlantingID,
		ProductionObtained:     requestPerfomanceEntity.ProductionObtained,
		UnitProductionObtained: requestPerfomanceEntity.UnitProductionObtained,
		HarvestedArea:          requestPerfomanceEntity.HarvestedArea,
		UnitHarvestedArea:      requestPerfomanceEntity.UnitHarvestedArea,
		HarvestedDate:          requestPerfomanceEntity.HarvestedDate,
	}

	if err := p.perfomanceCultureRepository.UpdatePerfomancePlanting(id, entityPerfomancePlanting); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *PerfomancePlantingService) GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(id uint) (*responses.PerfomanceCultureResponse, error) {

	dBResultPerfomancePlanting, err := p.perfomanceCultureRepository.FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID(id)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	responsePerfomancePlanting := responses.PerfomanceCultureResponse{
		Planting: responses.BatchPlantiesResponse{
			BatchName:              dBResultPerfomancePlanting.BatchName,
			AgricultureCultureName: dBResultPerfomancePlanting.AgricultureCultureName,
			IsPlanting:             dBResultPerfomancePlanting.IsPlanting,
			StartDatePlanting:      dBResultPerfomancePlanting.StartDatePlanting,
		},

		ID:                         dBResultPerfomancePlanting.ID,
		ProductionObtained:         dBResultPerfomancePlanting.ProductionObtained,
		ProductionObtainedFormated: dBResultPerfomancePlanting.ProductionObtainedFormated,
		HarvestedArea:              dBResultPerfomancePlanting.HarvestedArea,
		HarvestedAreaFormated:      dBResultPerfomancePlanting.HarvestedAreaFormated,
		HarvestedDate:              dBResultPerfomancePlanting.HarvestedDate,
	}

	return &responsePerfomancePlanting, nil
}

func (p *PerfomancePlantingService) DeletePerfomancePlanting(id uint) error {

	if err := p.perfomanceCultureRepository.DeletePerfomancePlanting(id); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
