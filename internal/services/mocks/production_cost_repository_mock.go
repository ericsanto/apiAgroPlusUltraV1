package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type ProductionCostRepositoryMock struct {
	mock.Mock
}

func (pcrm *ProductionCostRepositoryMock) FindAllProductinCostRepository() ([]entities.ProductionCostEntity, error) {

	args := pcrm.Called()

	return args.Get(0).([]entities.ProductionCostEntity), args.Error(1)
}

func (pcrm *ProductionCostRepositoryMock) CreateProductionCost(entityProductionCost entities.ProductionCostEntity) error {

	args := pcrm.Called(entityProductionCost)

	return args.Error(0)
}

func (pcrm *ProductionCostRepositoryMock) FindProductionCostByID(id uint) (*entities.ProductionCostEntity, error) {

	args := pcrm.Called(id)

	response := args.Get(0).(*entities.ProductionCostEntity)

	if response == nil {
		return nil, args.Error(1)
	}

	return response, args.Error(1)
}

func (pcrm *ProductionCostRepositoryMock) UpdateProductionCost(id uint, entityProductCost entities.ProductionCostEntity) error {

	args := pcrm.Called(id, entityProductCost)

	return args.Error(0)
}

func (pcrm *ProductionCostRepositoryMock) DeleteProductionCost(id uint) error {

	args := pcrm.Called(id)

	return args.Error(0)
}
