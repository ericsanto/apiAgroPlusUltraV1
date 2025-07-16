package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type ProductionCostRepositoryMock struct {
	mock.Mock
}

func (pcrm *ProductionCostRepositoryMock) FindAllProductinCostRepository(batchID, farmID, userID, plantingID uint) ([]entities.ProductionCostEntity, error) {

	args := pcrm.Called(batchID, farmID, userID, plantingID)

	return args.Get(0).([]entities.ProductionCostEntity), args.Error(1)
}

func (pcrm *ProductionCostRepositoryMock) CreateProductionCost(entityProductionCost entities.ProductionCostEntity) error {

	args := pcrm.Called(entityProductionCost)

	return args.Error(0)
}

func (pcrm *ProductionCostRepositoryMock) FindProductionCostByID(batchID, farmID, userID, plantingID, productionCostID uint) (*entities.ProductionCostEntity, error) {

	args := pcrm.Called(batchID, farmID, userID, plantingID, productionCostID)

	response := args.Get(0).(*entities.ProductionCostEntity)

	if response == nil {
		return nil, args.Error(1)
	}

	return response, args.Error(1)
}

func (pcrm *ProductionCostRepositoryMock) UpdateProductionCost(batchID, farmID, userID, plantingID, productionCostID uint, entityProductCost entities.ProductionCostEntity) error {

	args := pcrm.Called(batchID, farmID, userID, plantingID, productionCostID, entityProductCost)

	return args.Error(0)
}

func (pcrm *ProductionCostRepositoryMock) DeleteProductionCost(batchID, farmID, userID, plantingID, productionCostID uint) error {

	args := pcrm.Called(batchID, farmID, userID, plantingID, productionCostID)

	return args.Error(0)
}
