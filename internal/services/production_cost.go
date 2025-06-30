package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type ProductionCostServiceInterface interface {
	GetAllProductionCost() ([]responses.ProductionCostResponse, error)
	PostProductionCost(requestProductionCost requests.ProductionCostRequest) error
	GetAllProductionCostByID(id uint) (*responses.ProductionCostResponse, error)
	PutProductionCost(id uint, requestProduction requests.ProductionCostRequest) error
	DeleteProductionCost(id uint) error
}

type ProductionCostService struct {
	productionCostRepository repositories.ProductionCostRepositoryInterface
}

func NewProductionCostService(productionCostRepository repositories.ProductionCostRepositoryInterface) ProductionCostServiceInterface {
	return &ProductionCostService{productionCostRepository: productionCostRepository}
}

func (p *ProductionCostService) GetAllProductionCost() ([]responses.ProductionCostResponse, error) {

	productionCostEntity, err := p.productionCostRepository.FindAllProductinCostRepository()
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

func (p *ProductionCostService) PostProductionCost(requestProductionCost requests.ProductionCostRequest) error {

	entityProductionCost := entities.ProductionCostEntity{
		PlantingID:  requestProductionCost.PlantingID,
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

func (p *ProductionCostService) GetAllProductionCostByID(id uint) (*responses.ProductionCostResponse, error) {

	productionCostEntity, err := p.productionCostRepository.FindProductionCostByID(id)
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

func (p *ProductionCostService) PutProductionCost(id uint, requestProduction requests.ProductionCostRequest) error {

	entityProductionCost := entities.ProductionCostEntity{
		PlantingID:  requestProduction.PlantingID,
		Item:        requestProduction.Item,
		Unit:        requestProduction.Unit,
		Quantity:    requestProduction.Quantity,
		CostPerUnit: requestProduction.CostPerUnit,
		CostDate:    requestProduction.CostDate,
	}

	if err := p.productionCostRepository.UpdateProductionCost(id, entityProductionCost); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (p *ProductionCostService) DeleteProductionCost(id uint) error {

	if err := p.productionCostRepository.DeleteProductionCost(id); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
