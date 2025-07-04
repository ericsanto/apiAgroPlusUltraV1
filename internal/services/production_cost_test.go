package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
)

func TestPostProductionCost_Success(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	entityProductCost := entities.ProductionCostEntity{
		ID:          uint(0),
		PlantingID:  uint(1),
		Item:        "TRATOR",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 5000000,
		CostDate:    time.Now(),
	}

	requestProductionCost := requests.ProductionCostRequest{
		PlantingID:  entityProductCost.PlantingID,
		Item:        entityProductCost.Item,
		Unit:        entityProductCost.Unit,
		Quantity:    entityProductCost.Quantity,
		CostPerUnit: entityProductCost.CostPerUnit,
		CostDate:    entityProductCost.CostDate,
	}

	mockRepo.On("CreateProductionCost", entityProductCost).Return(nil)

	err := service.PostProductionCost(requestProductionCost)

	assert.NoError(t, err)
	assert.Equal(t, entityProductCost.PlantingID, requestProductionCost.PlantingID)
	assert.EqualValues(t, entityProductCost.CostDate, requestProductionCost.CostDate)
	assert.Equal(t, entityProductCost.CostPerUnit, requestProductionCost.CostPerUnit)
	assert.EqualValues(t, entityProductCost.Item, entityProductCost.Item)
	assert.EqualValues(t, entityProductCost.Unit, requestProductionCost.Unit)
	assert.Equal(t, entityProductCost.Quantity, requestProductionCost.Quantity)

	mockRepo.AssertExpectations(t)
}

func TestPostProductionCost_Error(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	entityProductCost := entities.ProductionCostEntity{
		ID:          uint(0),
		PlantingID:  uint(1),
		Item:        "TRATOR",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 5000000,
		CostDate:    time.Now(),
	}

	requestProductionCost := requests.ProductionCostRequest{
		PlantingID:  entityProductCost.PlantingID,
		Item:        entityProductCost.Item,
		Unit:        entityProductCost.Unit,
		Quantity:    entityProductCost.Quantity,
		CostPerUnit: entityProductCost.CostPerUnit,
		CostDate:    entityProductCost.CostDate,
	}

	mockRepo.On("CreateProductionCost", entityProductCost).Return(fmt.Errorf("erro ao tentar criar custo"))

	err := service.PostProductionCost(requestProductionCost)

	assert.Contains(t, err.Error(), "erro ao tentar criar custo")

	mockRepo.AssertExpectations(t)

}

func TestGetAllProductionCost_Success(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	entityProductCost1 := entities.ProductionCostEntity{
		ID:          uint(0),
		PlantingID:  uint(1),
		Item:        "TRATOR",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 5000000,
		CostDate:    time.Now(),
	}

	entityProductCost2 := entities.ProductionCostEntity{
		ID:          uint(1),
		PlantingID:  uint(2),
		Item:        "Fungicida",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 500,
		CostDate:    time.Now(),
	}

	var entitiesProductionCosts []entities.ProductionCostEntity

	entitiesProductionCosts = append(entitiesProductionCosts, entityProductCost1)
	entitiesProductionCosts = append(entitiesProductionCosts, entityProductCost2)

	mockRepo.On("FindAllProductinCostRepository").Return(entitiesProductionCosts, nil)

	responseProductionCosts, err := service.GetAllProductionCost()

	assert.Nil(t, err)

	for i, v := range entitiesProductionCosts {
		assert.EqualValues(t, v.Item, responseProductionCosts[i].Item)
		assert.EqualValues(t, v.Unit, responseProductionCosts[i].Unit)
		assert.EqualValues(t, v.CostDate, responseProductionCosts[i].CostDate)
		assert.Equal(t, v.PlantingID, responseProductionCosts[i].PlantingID)
		assert.Equal(t, v.Quantity, responseProductionCosts[i].Quantity)
		assert.Equal(t, v.CostPerUnit, responseProductionCosts[i].CostPerUnit)
	}

	mockRepo.AssertExpectations(t)
}

func TestGetAllProductionCost_Error(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	mockRepo.On("FindAllProductinCostRepository").Return([]entities.ProductionCostEntity{}, fmt.Errorf("erro ao buscar custos de produtos"))

	responseProductionCosts, err := service.GetAllProductionCost()

	assert.Nil(t, responseProductionCosts)
	assert.Contains(t, err.Error(), "erro ao buscar custos de produtos")

	mockRepo.AssertExpectations(t)
}

func TestGetAllProductionCostByID_Success(t *testing.T) {

	mockeRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockeRepo}

	id := uint(4)

	entityProductCost := entities.ProductionCostEntity{
		ID:          id,
		PlantingID:  uint(1),
		Item:        "TRATOR",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 5000000,
		CostDate:    time.Now(),
	}

	mockeRepo.On("FindProductionCostByID", id).Return(&entityProductCost, nil)

	responseProductionCost, err := service.GetAllProductionCostByID(id)

	assert.Nil(t, err)
	assert.Equal(t, entityProductCost.ID, responseProductionCost.ID)
	assert.Equal(t, entityProductCost.PlantingID, responseProductionCost.PlantingID)
	assert.EqualValues(t, entityProductCost.Item, responseProductionCost.Item)
	assert.EqualValues(t, entityProductCost.Unit, responseProductionCost.Unit)
	assert.Equal(t, entityProductCost.Quantity, responseProductionCost.Quantity)
	assert.Equal(t, entityProductCost.CostPerUnit, responseProductionCost.CostPerUnit)
	assert.EqualValues(t, entityProductCost.CostDate, responseProductionCost.CostDate)

}

func TestGetAllProductionCostByID_ErrorIDNotExist(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	id := uint(1)

	mockRepo.On("FindProductionCostByID", id).Return(&entities.ProductionCostEntity{}, fmt.Errorf("não existe custo com o id"))

	responseProductionCost, err := service.GetAllProductionCostByID(id)

	assert.Contains(t, err.Error(), "não existe custo com o id")
	assert.Nil(t, responseProductionCost)
}

func TestGetAllProductionCostByID_ErrorSerachProducitonCost(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	id := uint(1)

	entityProductCost := entities.ProductionCostEntity{
		ID:          id,
		PlantingID:  uint(1),
		Item:        "TRATOR",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 5000000,
		CostDate:    time.Now(),
	}

	mockRepo.On("FindProductionCostByID", id).Return(&entityProductCost, fmt.Errorf("erro ao buscar custo"))

	responseProductionCost, err := service.GetAllProductionCostByID(id)

	assert.Contains(t, err.Error(), "erro ao buscar custo")
	assert.Nil(t, responseProductionCost)

	mockRepo.AssertExpectations(t)
}

func TestPutProductionCost_Success(t *testing.T) {

	mockeRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockeRepo}

	id := uint(5)

	requestProductionCost := requests.ProductionCostRequest{
		PlantingID:  uint(3),
		Item:        "adubo 10-10",
		Unit:        "kg",
		Quantity:    5,
		CostPerUnit: 150,
		CostDate:    time.Now(),
	}

	expectedEntityProductionCost := entities.ProductionCostEntity{
		PlantingID:  requestProductionCost.PlantingID,
		Item:        requestProductionCost.Item,
		Unit:        requestProductionCost.Unit,
		Quantity:    requestProductionCost.Quantity,
		CostPerUnit: requestProductionCost.CostPerUnit,
		CostDate:    requestProductionCost.CostDate,
	}

	mockeRepo.On("UpdateProductionCost", id, expectedEntityProductionCost).Return(nil)

	err := service.PutProductionCost(id, requestProductionCost)

	assert.Nil(t, err)

	mockeRepo.AssertExpectations(t)
}

func TestPutProductionCost_ErrorNotIDExist(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	id := uint(5)

	requestProductionCost := requests.ProductionCostRequest{
		PlantingID:  uint(3),
		Item:        "adubo 10-10",
		Unit:        "kg",
		Quantity:    5,
		CostPerUnit: 150,
		CostDate:    time.Now(),
	}

	expectedEntityProductionCost := entities.ProductionCostEntity{
		PlantingID:  requestProductionCost.PlantingID,
		Item:        requestProductionCost.Item,
		Unit:        requestProductionCost.Unit,
		Quantity:    requestProductionCost.Quantity,
		CostPerUnit: requestProductionCost.CostPerUnit,
		CostDate:    requestProductionCost.CostDate,
	}

	mockRepo.On("UpdateProductionCost", id, expectedEntityProductionCost).Return(fmt.Errorf("não existe custo com o id"))

	err := service.PutProductionCost(id, requestProductionCost)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "não existe custo com o id")

	mockRepo.AssertExpectations(t)
}

func TestPutProductionCost_Error(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	id := uint(5)

	requestProductionCost := requests.ProductionCostRequest{
		PlantingID:  uint(3),
		Item:        "adubo 10-10",
		Unit:        "kg",
		Quantity:    5,
		CostPerUnit: 150,
		CostDate:    time.Now(),
	}

	expectedEntityProductionCost := entities.ProductionCostEntity{
		PlantingID:  requestProductionCost.PlantingID,
		Item:        requestProductionCost.Item,
		Unit:        requestProductionCost.Unit,
		Quantity:    requestProductionCost.Quantity,
		CostPerUnit: requestProductionCost.CostPerUnit,
		CostDate:    requestProductionCost.CostDate,
	}

	mockRepo.On("UpdateProductionCost", id, expectedEntityProductionCost).Return(fmt.Errorf("erro ao atualizar custo"))

	err := service.PutProductionCost(id, requestProductionCost)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro ao atualizar custo")

	mockRepo.AssertExpectations(t)
}

func TestDeleteProductionCost_Success(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	id := uint(5)

	mockRepo.On("DeleteProductionCost", id).Return(nil)

	err := service.DeleteProductionCost(id)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteProductionCost_NotIDExist(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	id := uint(5)

	mockRepo.On("DeleteProductionCost", id).Return(fmt.Errorf("não existe custo com o id"))

	err := service.DeleteProductionCost(id)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "não existe custo com o id")

	mockRepo.AssertExpectations(t)
}

func TestDeleteProductionCost_Error(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	id := uint(5)

	mockRepo.On("DeleteProductionCost", id).Return(fmt.Errorf("erro ao tentar deletar custo"))

	err := service.DeleteProductionCost(id)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro ao tentar deletar custo")

	mockRepo.AssertExpectations(t)
}
