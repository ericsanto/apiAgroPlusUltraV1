package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type ProductionCostServiceInterface interface {
	GetAllProductionCost(batchID, farmID, userID, plantingID uint) ([]responses.ProductionCostResponse, error)
	PostProductionCost(batchID, farmID, userID, plantingID uint, requestProductionCost requests.ProductionCostRequest) error
	GetAllProductionCostByID(batchID, farmID, userID, plantingID, productionCostID uint) (*responses.ProductionCostResponse, error)
	PutProductionCost(batchID, farmID, userID, plantingID, productionCostID uint, requestProduction requests.ProductionCostRequest) error
	DeleteProductionCost(batchID, farmID, userID, plantingID, productionCostID uint) error
}

type ProductionCostService struct {
	productionCostRepository repositories.ProductionCostRepositoryInterface
	plantingService          PlantingServiceInterface
	batchService             BatchServiceInterface
	farmService              FarmServiceInterface
}

func NewProductionCostService(productionCostRepository repositories.ProductionCostRepositoryInterface,
	plantingService PlantingServiceInterface,
	batchService BatchServiceInterface,
	farmService FarmServiceInterface) ProductionCostServiceInterface {
	return &ProductionCostService{productionCostRepository: productionCostRepository,
		plantingService: plantingService,
		batchService:    batchService,
		farmService:     farmService}
}

func (p *ProductionCostService) GetAllProductionCost(batchID, farmID, userID, plantingID uint) ([]responses.ProductionCostResponse, error) {

	productionCostEntity, err := p.productionCostRepository.FindAllProductinCostRepository(batchID, farmID, userID, plantingID)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	var productionCostResponseList []responses.ProductionCostResponse

	for _, v := range productionCostEntity {
		productionCost := responses.ProductionCostResponse{
			ID:          v.ID,
			PlantingID:  v.PlantingID,
			Item:        v.Item,
			Unit:        v.Unit,
			Quantity:    v.Quantity,
			CostPerUnit: v.CostPerUnit,
			CostDate:    v.CostDate,
		}

		productionCostResponseList = append(productionCostResponseList, productionCost)
	}

	return productionCostResponseList, nil
}

func (p *ProductionCostService) PostProductionCost(batchID, farmID, userID, plantingID uint, requestProductionCost requests.ProductionCostRequest) error {

	if _, err := p.farmService.GetFarmByID(userID, farmID); err != nil {
		return err
	}

	if _, err := p.batchService.GetBatchFindById(userID, farmID, batchID); err != nil {
		return err
	}

	if _, err := p.plantingService.GetByParam(userID, farmID, batchID); err != nil {
		return err
	}

	entityProductionCost := entities.ProductionCostEntity{
		PlantingID:  plantingID,
		Item:        requestProductionCost.Item,
		Unit:        requestProductionCost.Unit,
		Quantity:    requestProductionCost.Quantity,
		CostPerUnit: requestProductionCost.CostPerUnit,
		CostDate:    requestProductionCost.CostDate,
	}

	if err := p.productionCostRepository.CreateProductionCost(entityProductionCost); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *ProductionCostService) GetAllProductionCostByID(batchID, farmID, userID, plantingID, productionCost uint) (*responses.ProductionCostResponse, error) {

	productionCostEntity, err := p.productionCostRepository.FindProductionCostByID(batchID, farmID, userID, plantingID, productionCost)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	productionCostResponse := responses.ProductionCostResponse{
		ID:          productionCostEntity.ID,
		PlantingID:  productionCostEntity.PlantingID,
		Item:        productionCostEntity.Item,
		Unit:        productionCostEntity.Unit,
		Quantity:    productionCostEntity.Quantity,
		CostPerUnit: productionCostEntity.CostPerUnit,
		CostDate:    productionCostEntity.CostDate,
	}

	return &productionCostResponse, nil
}

func (p *ProductionCostService) PutProductionCost(batchID, farmID, userID, plantingID, productionCostID uint, requestProduction requests.ProductionCostRequest) error {

	if _, err := p.farmService.GetFarmByID(userID, farmID); err != nil {
		return err
	}

	if _, err := p.batchService.GetBatchFindById(userID, farmID, batchID); err != nil {
		return err
	}

	if _, err := p.plantingService.GetByParam(userID, farmID, batchID); err != nil {
		return err
	}

	entityProductionCost := entities.ProductionCostEntity{
		PlantingID:  plantingID,
		Item:        requestProduction.Item,
		Unit:        requestProduction.Unit,
		Quantity:    requestProduction.Quantity,
		CostPerUnit: requestProduction.CostPerUnit,
		CostDate:    requestProduction.CostDate,
	}

	if err := p.productionCostRepository.UpdateProductionCost(batchID, farmID, userID, plantingID, productionCostID, entityProductionCost); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *ProductionCostService) DeleteProductionCost(batchID, farmID, userID, plantingID, productionCostID uint) error {

	if err := p.productionCostRepository.DeleteProductionCost(batchID, farmID, userID, plantingID, productionCostID); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
