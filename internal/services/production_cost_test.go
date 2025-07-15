package services

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
)

const (
	productionCostMOCKID = uint(1)
)

func SetupTestProductioCost() (ProductionCostService, entities.ProductionCostEntity, requests.ProductionCostRequest,
	*mocks.ProductionCostRepositoryMock, *mocks.FarmRepositoryMock, *mocks.BatchRepositoryMock, *mocks.PlantingRepositoryMock,
	responses.FarmResponse, entities.BatchEntity, entities.PlantingEntity) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	mockRepoPlanting := new(mocks.PlantingRepositoryMock)

	servicePlanting := PlantingService{mockRepoPlanting}

	mockRepoFarm := new(mocks.FarmRepositoryMock)

	serviceFarm := FarmService{mockRepoFarm}

	mockeRepoBatch := new(mocks.BatchRepositoryMock)

	serviceBatch := BatchService{mockeRepoBatch}

	service := ProductionCostService{productionCostRepository: mockRepo, plantingService: &servicePlanting,
		farmService: &serviceFarm, batchService: &serviceBatch}

	timeCurrent := time.Date(2025, time.July, 3, 18, 19, 28, 674505796, time.Local)

	requestProductionCost := requests.ProductionCostRequest{
		Item:        "TRATOR",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 5000000,
		CostDate:    timeCurrent,
	}

	farmResponse := responses.FarmResponse{
		ID:   farmMOCKID,
		Name: "teste",
	}

	batchEntity := entities.BatchEntity{
		ID:     batchMOCKID,
		Name:   "teste",
		Area:   500,
		Unit:   "ha",
		FarmID: farmResponse.ID,
	}

	entityPlanting := entities.PlantingEntity{
		ID:                   plantingMOCKID,
		BatchID:              batchEntity.ID,
		AgricultureCultureID: uint(5),
		IsPlanting:           false,
		StartDatePlanting:    timeCurrent,
		ExpectedProduction:   0,
		SpaceBetweenPlants:   0.50,
		SpaceBetweenRows:     0.30,
		IrrigationTypeID:     uint(4),
	}

	entityProductionCost := entities.ProductionCostEntity{
		PlantingID:  entityPlanting.ID,
		Item:        requestProductionCost.Item,
		Unit:        requestProductionCost.Unit,
		Quantity:    requestProductionCost.Quantity,
		CostPerUnit: requestProductionCost.CostPerUnit,
		CostDate:    requestProductionCost.CostDate,
	}

	return service, entityProductionCost, requestProductionCost, mockRepo, mockRepoFarm, mockeRepoBatch,
		mockRepoPlanting, farmResponse, batchEntity, entityPlanting
}

func TestPostProductionCost_Success(t *testing.T) {

	service, entityProductionCost, request, mockRepo, mockRepoFarm, mockRepoBatch,
		mockRepoPlanting, farmResponse, batchEntity, plantingEntity := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&batchEntity, nil)

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.NotNil(t, batch)
	assert.Nil(t, err)

	mockRepoPlanting.On("FindByParamPlanting", userMOCKID, farm.ID, batch.ID).Return(plantingEntity, nil)

	planting, err := service.plantingService.GetByParam(userMOCKID, farm.ID, batch.ID)

	assert.NotNil(t, planting)
	assert.Nil(t, err)

	mockRepo.On("CreateProductionCost", entityProductionCost).Return(nil)

	err = service.PostProductionCost(batch.ID, farm.ID, userMOCKID, planting.ID, request)

	assert.NoError(t, err)
	assert.EqualValues(t, entityProductionCost.CostDate, request.CostDate)
	assert.Equal(t, entityProductionCost.CostPerUnit, request.CostPerUnit)
	assert.EqualValues(t, entityProductionCost.Item, request.Item)
	assert.EqualValues(t, entityProductionCost.Unit, request.Unit)
	assert.Equal(t, entityProductionCost.Quantity, request.Quantity)

	mockRepo.AssertExpectations(t)
}

func TestPostProductionCost_Error(t *testing.T) {

	service, entityProductionCost, request, mockRepo, mockRepoFarm, mockRepoBatch,
		mockRepoPlanting, farmResponse, batchEntity, plantingEntity := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&batchEntity, nil)

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.NotNil(t, batch)
	assert.Nil(t, err)

	mockRepoPlanting.On("FindByParamPlanting", userMOCKID, farm.ID, batch.ID).Return(plantingEntity, nil)

	planting, err := service.plantingService.GetByParam(userMOCKID, farm.ID, batch.ID)

	assert.NotNil(t, planting)
	assert.Nil(t, err)

	mockRepo.On("CreateProductionCost", entityProductionCost).Return(fmt.Errorf("erro ao tentar criar custo"))

	err = service.PostProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, request)

	assert.Contains(t, err.Error(), "erro ao tentar criar custo")

	mockRepo.AssertExpectations(t)

}

func TestPostProductionCost_PlantingNotFound(t *testing.T) {

	service, _, request, _, mockRepoFarm, mockRepoBatch,
		mockRepoPlanting, farmResponse, batchEntity, _ := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&batchEntity, nil)

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.NotNil(t, batch)
	assert.Nil(t, err)

	mockRepoPlanting.On("FindByParamPlanting", batch.ID, farm.ID, userMOCKID).Return(entities.PlantingEntity{}, errors.New("erro"))

	planting, err := service.plantingService.GetByParam(userMOCKID, farm.ID, batch.ID)

	assert.NotNil(t, err)
	assert.Nil(t, planting)

	err = service.PostProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, request)

	assert.NotNil(t, err)

	mockRepoFarm.AssertCalled(t, "FindByID", userMOCKID, farmMOCKID)
	mockRepoBatch.AssertCalled(t, "BatchFindById", userMOCKID, farm.ID, batchMOCKID)
	mockRepoPlanting.AssertCalled(t, "FindByParamPlanting", batch.ID, farm.ID, userMOCKID)

	mockRepoFarm.AssertExpectations(t)
	mockRepoBatch.AssertExpectations(t)
}

func TestPostProductionCost_FarmNotFound(t *testing.T) {

	service, _, request, _, mockRepoFarm, _,
		_, _, _, _ := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&responses.FarmResponse{}, errors.New("erro"))

	_, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro")

	err = service.PostProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, request)

	assert.NotNil(t, err)

	mockRepoFarm.AssertExpectations(t)
}

func TestPostProductionCost_BatchNotFound(t *testing.T) {

	service, _, request, _, mockRepoFarm, mockRepoBatch,
		_, farmResponse, _, _ := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&entities.BatchEntity{}, errors.New("erro"))

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.Nil(t, batch)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro")

	err = service.PostProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, request)

	assert.NotNil(t, err)

	mockRepoFarm.AssertCalled(t, "FindByID", userMOCKID, farmMOCKID)
	mockRepoBatch.AssertCalled(t, "BatchFindById", userMOCKID, farm.ID, batchMOCKID)

	mockRepoFarm.AssertExpectations(t)
	mockRepoBatch.AssertExpectations(t)
}

func TestGetAllProductionCost_Success(t *testing.T) {

	service, entityProductionCost1, _, mockRepo, _, _,
		_, _, _, _ := SetupTestProductioCost()

	entityProductionCost2 := entities.ProductionCostEntity{
		ID:          productionCostMOCKID,
		PlantingID:  uint(2),
		Item:        "Fungicida",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 500,
		CostDate:    time.Now(),
	}

	var entitiesProductionCosts []entities.ProductionCostEntity

	entitiesProductionCosts = append(entitiesProductionCosts, entityProductionCost1)
	entitiesProductionCosts = append(entitiesProductionCosts, entityProductionCost2)

	mockRepo.On("FindAllProductinCostRepository", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(entitiesProductionCosts, nil)

	responseProductionCosts, err := service.GetAllProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)

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

	mockRepo.On("FindAllProductinCostRepository", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return([]entities.ProductionCostEntity{}, fmt.Errorf("erro ao buscar custos de produtos"))

	responseProductionCosts, err := service.GetAllProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)

	assert.Nil(t, responseProductionCosts)
	assert.Contains(t, err.Error(), "erro ao buscar custos de produtos")

	mockRepo.AssertExpectations(t)
}

func TestGetAllProductionCostByID_Success(t *testing.T) {

	mockeRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockeRepo}

	entityProductCost := entities.ProductionCostEntity{
		ID:          productionCostMOCKID,
		PlantingID:  uint(1),
		Item:        "TRATOR",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 5000000,
		CostDate:    time.Now(),
	}

	mockeRepo.On("FindProductionCostByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID).Return(&entityProductCost, nil)

	responseProductionCost, err := service.GetAllProductionCostByID(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID)

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

	mockRepo.On("FindProductionCostByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID).Return(&entities.ProductionCostEntity{}, fmt.Errorf("não existe custo com o id"))

	responseProductionCost, err := service.GetAllProductionCostByID(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID)

	assert.Contains(t, err.Error(), "não existe custo com o id")
	assert.Nil(t, responseProductionCost)
}

func TestGetAllProductionCostByID_ErrorSerachProducitonCost(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	entityProductCost := entities.ProductionCostEntity{
		ID:          productionCostMOCKID,
		PlantingID:  uint(1),
		Item:        "TRATOR",
		Unit:        "1",
		Quantity:    1,
		CostPerUnit: 5000000,
		CostDate:    time.Now(),
	}

	mockRepo.On("FindProductionCostByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID).Return(&entityProductCost, fmt.Errorf("erro ao buscar custo"))

	responseProductionCost, err := service.GetAllProductionCostByID(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID)

	assert.Contains(t, err.Error(), "erro ao buscar custo")
	assert.Nil(t, responseProductionCost)

	mockRepo.AssertExpectations(t)
}

func TestPutProductionCost_Success(t *testing.T) {

	service, entityProductionCost, request, mockRepo, mockRepoFarm, mockRepoBatch,
		mockRepoPlanting, farmResponse, batchEntity, plantingEntity := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&batchEntity, nil)

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.NotNil(t, batch)
	assert.Nil(t, err)

	mockRepoPlanting.On("FindByParamPlanting", userMOCKID, farm.ID, batch.ID).Return(plantingEntity, nil)

	planting, err := service.plantingService.GetByParam(userMOCKID, farm.ID, batch.ID)

	assert.NotNil(t, planting)
	assert.Nil(t, err)

	mockRepo.On("UpdateProductionCost", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, entityProductionCost).Return(nil)

	err = service.PutProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, request)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestPutProductionCost_ErrorNotIDExist(t *testing.T) {

	service, entityProductionCost, request, mockRepo, mockRepoFarm, mockRepoBatch,
		mockRepoPlanting, farmResponse, batchEntity, plantingEntity := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&batchEntity, nil)

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.NotNil(t, batch)
	assert.Nil(t, err)

	mockRepoPlanting.On("FindByParamPlanting", userMOCKID, farm.ID, batch.ID).Return(plantingEntity, nil)

	planting, err := service.plantingService.GetByParam(userMOCKID, farm.ID, batch.ID)

	assert.NotNil(t, planting)
	assert.Nil(t, err)

	mockRepo.On("UpdateProductionCost", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, entityProductionCost).Return(fmt.Errorf("não existe custo com o id"))

	err = service.PutProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, request)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "não existe custo com o id")

	mockRepo.AssertExpectations(t)
}

func TestPutProductionCost_Error(t *testing.T) {

	service, entityProductionCost, request, mockRepo, mockRepoFarm, mockRepoBatch,
		mockRepoPlanting, farmResponse, batchEntity, plantingEntity := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&batchEntity, nil)

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.NotNil(t, batch)
	assert.Nil(t, err)

	mockRepoPlanting.On("FindByParamPlanting", userMOCKID, farm.ID, batch.ID).Return(plantingEntity, nil)

	planting, err := service.plantingService.GetByParam(userMOCKID, farm.ID, batch.ID)

	assert.NotNil(t, planting)
	assert.Nil(t, err)

	mockRepo.On("UpdateProductionCost", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, entityProductionCost).Return(fmt.Errorf("erro ao atualizar custo"))

	err = service.PutProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, request)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro ao atualizar custo")

	mockRepo.AssertExpectations(t)
}

func TestPutProductionCost_FarmNotFound(t *testing.T) {

	service, _, request, _, mockRepoFarm, _,
		_, _, _, _ := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&responses.FarmResponse{}, errors.New("erro"))

	_, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro")

	err = service.PutProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, request)

	assert.NotNil(t, err)

	mockRepoFarm.AssertExpectations(t)
}

func TestPutProductionCost_BatchNotFound(t *testing.T) {

	service, _, request, _, mockRepoFarm, mockRepoBatch,
		_, farmResponse, _, _ := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&entities.BatchEntity{}, errors.New("erro"))

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.Nil(t, batch)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro")

	err = service.PutProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, request)

	assert.NotNil(t, err)

	mockRepoFarm.AssertCalled(t, "FindByID", userMOCKID, farmMOCKID)
	mockRepoBatch.AssertCalled(t, "BatchFindById", userMOCKID, farm.ID, batchMOCKID)

	mockRepoFarm.AssertExpectations(t)
	mockRepoBatch.AssertExpectations(t)
}

func TestPutProductionCost_PlantingNotFound(t *testing.T) {

	service, _, request, _, mockRepoFarm, mockRepoBatch,
		mockRepoPlanting, farmResponse, batchEntity, _ := SetupTestProductioCost()

	mockRepoFarm.On("FindByID", userMOCKID, farmMOCKID).Return(&farmResponse, nil)

	farm, err := service.farmService.GetFarmByID(userMOCKID, farmMOCKID)

	assert.NotNil(t, farm)
	assert.Nil(t, err)

	mockRepoBatch.On("BatchFindById", userMOCKID, farm.ID, batchMOCKID).Return(&batchEntity, nil)

	batch, err := service.batchService.GetBatchFindById(userMOCKID, farm.ID, batchMOCKID)

	assert.NotNil(t, batch)
	assert.Nil(t, err)

	mockRepoPlanting.On("FindByParamPlanting", batch.ID, farm.ID, userMOCKID).Return(entities.PlantingEntity{}, errors.New("erro"))

	planting, err := service.plantingService.GetByParam(userMOCKID, farm.ID, batch.ID)

	assert.NotNil(t, err)
	assert.Nil(t, planting)

	err = service.PutProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID, request)

	assert.NotNil(t, err)

	mockRepoFarm.AssertCalled(t, "FindByID", userMOCKID, farmMOCKID)
	mockRepoBatch.AssertCalled(t, "BatchFindById", userMOCKID, farm.ID, batchMOCKID)
	mockRepoPlanting.AssertCalled(t, "FindByParamPlanting", batch.ID, farm.ID, userMOCKID)

	mockRepoFarm.AssertExpectations(t)
	mockRepoBatch.AssertExpectations(t)
}

func TestDeleteProductionCost_Success(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	mockRepo.On("DeleteProductionCost", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID).Return(nil)

	err := service.DeleteProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteProductionCost_NotIDExist(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	mockRepo.On("DeleteProductionCost", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID).Return(fmt.Errorf("não existe custo com o id"))

	err := service.DeleteProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "não existe custo com o id")

	mockRepo.AssertExpectations(t)
}

func TestDeleteProductionCost_Error(t *testing.T) {

	mockRepo := new(mocks.ProductionCostRepositoryMock)

	service := ProductionCostService{productionCostRepository: mockRepo}

	mockRepo.On("DeleteProductionCost", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID).Return(fmt.Errorf("erro ao tentar deletar custo"))

	err := service.DeleteProductionCost(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, productionCostMOCKID)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro ao tentar deletar custo")

	mockRepo.AssertExpectations(t)
}
