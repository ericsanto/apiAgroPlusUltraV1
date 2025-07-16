package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

const (
	salePlantingMOCKID = uint(1)
)

func SetupTestSalePlantingService() (*mocks.SalePlantingRepositoryMock, *mocks.FarmRepositoryMock, *mocks.BatchRepositoryMock,
	*mocks.PlantingRepositoryMock, SalePlantingService, entities.SalePlantingEntity, requests.SalePlantingRequest,
	responses.FarmResponse, entities.BatchEntity, entities.PlantingEntity) {

	mockRepoPlanting := new(mocks.PlantingRepositoryMock)

	servicePlanting := PlantingService{mockRepoPlanting}

	mockRepoFarm := new(mocks.FarmRepositoryMock)

	serviceFarm := FarmService{mockRepoFarm}

	mockRepoBatch := new(mocks.BatchRepositoryMock)

	serviceBatch := BatchService{mockRepoBatch}

	mockRepo := new(mocks.SalePlantingRepositoryMock)
	service := SalePlantingService{salePlantingRepository: mockRepo, farmService: &serviceFarm,
		plantingService: &servicePlanting, batchService: &serviceBatch}

	requestSalePlanting := requests.SalePlantingRequest{
		ValueSale: 500000000,
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

	timeCurrent := time.Date(2025, time.July, 3, 18, 19, 28, 674505796, time.Local)

	entityPlanting := entities.PlantingEntity{
		BatchID:              batchEntity.ID,
		AgricultureCultureID: uint(5),
		IsPlanting:           false,
		StartDatePlanting:    timeCurrent,
		ExpectedProduction:   0,
		SpaceBetweenPlants:   0.50,
		SpaceBetweenRows:     0.30,
		IrrigationTypeID:     uint(4),
	}

	entitySalePlanting := entities.SalePlantingEntity{
		PlantingID: entityPlanting.ID,
		ValueSale:  requestSalePlanting.ValueSale,
	}

	return mockRepo, mockRepoFarm, mockRepoBatch, mockRepoPlanting, service, entitySalePlanting, requestSalePlanting, farmResponse, batchEntity, entityPlanting
}
func TestPostSalePlanting_Success(t *testing.T) {

	mockRepo, mockRepoFarm, mockRepoBatch, mockRepoPlanting, service, entitySalePlanting, requestSalePlanting, farmResponse,
		batchEntity, plantingEntity := SetupTestSalePlantingService()

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

	mockRepo.On("CreateSalePlantingRepository", batch.ID, farm.ID, userMOCKID, planting.ID, entitySalePlanting).Return(nil)

	err = service.PostSalePlanting(batch.ID, farm.ID, userMOCKID, plantingEntity.ID, requestSalePlanting)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "CreateSalePlantingRepository", batch.ID, farm.ID, userMOCKID, planting.ID, entitySalePlanting)
}

func TestPostSalePlanting_Error(t *testing.T) {

	mockRepo, mockRepoFarm, mockRepoBatch, mockRepoPlanting, service, entitySalePlanting, requestSalePlanting, farmResponse,
		batchEntity, plantingEntity := SetupTestSalePlantingService()

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

	mockRepo.On("CreateSalePlantingRepository", batch.ID, farm.ID, userMOCKID, planting.ID, entitySalePlanting).Return(fmt.Errorf("erro ao cadastrar venda"))

	err = service.PostSalePlanting(batch.ID, farm.ID, userMOCKID, planting.ID, requestSalePlanting)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro ao cadastrar venda")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "CreateSalePlantingRepository", batch.ID, farm.ID, userMOCKID, planting.ID, entitySalePlanting)
}

func TestPostSalePlanting_Return_ConstraintViolatedPlantingID(t *testing.T) {

	mockRepo, mockRepoFarm, mockRepoBatch, mockRepoPlanting, service, entitySalePlanting, requestSalePlanting, farmResponse,
		batchEntity, plantingEntity := SetupTestSalePlantingService()

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

	mockRepo.On("CreateSalePlantingRepository", batch.ID, farm.ID, userMOCKID, planting.ID, entitySalePlanting).Return(fmt.Errorf("%w", myerror.ErrDuplicateSale))

	err = service.PostSalePlanting(batch.ID, farm.ID, userMOCKID, planting.ID, requestSalePlanting)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrDuplicateSale)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "CreateSalePlantingRepository", batch.ID, farm.ID, userMOCKID, planting.ID, entitySalePlanting)
}

func TestGetAllSalePlanting_Success(t *testing.T) {

	mockRepo, _, _, _, service, entitySalePlanting, _, _, _, _ := SetupTestSalePlantingService()

	var entitiesSalePlanting []entities.SalePlantingEntity
	entitiesSalePlanting = append(entitiesSalePlanting, entitySalePlanting)

	mockRepo.On("FindAllSalePlanting", batchMOCKID, farmMOCKID, userMOCKID).Return(entitiesSalePlanting, nil)

	responsesSalePlanting, err := service.GetAllSalePlanting(batchMOCKID, farmMOCKID, userMOCKID)

	assert.Nil(t, err)
	assert.Equal(t, entitiesSalePlanting[0].ValueSale, responsesSalePlanting[0].ValueSale)
	assert.Equal(t, entitiesSalePlanting[0].PlantingID, responsesSalePlanting[0].PlantingID)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindAllSalePlanting", batchMOCKID, farmMOCKID, userMOCKID)

}

func TestGetAllSalePlanting_Error(t *testing.T) {

	mockRepo, _, _, _, service, _, _, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("FindAllSalePlanting", batchMOCKID, farmMOCKID, userMOCKID).Return([]entities.SalePlantingEntity(nil), fmt.Errorf("não foi possível buscar todas as vendas de plantações"))

	responsesSalePlanting, err := service.GetAllSalePlanting(batchMOCKID, farmMOCKID, userMOCKID)

	assert.NotNil(t, err)
	assert.Nil(t, responsesSalePlanting)
	assert.Contains(t, err.Error(), "não foi possível buscar todas as vendas de plantações")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindAllSalePlanting", batchMOCKID, farmMOCKID, userMOCKID)
}

func TestGetSalePlantingByID_Success(t *testing.T) {

	mockRepo, _, _, _, service, entitySalePlanting, _, _, _, _ := SetupTestSalePlantingService()

	entitySalePlanting.ID = salePlantingMOCKID

	mockRepo.On("FindSalePlantingByID", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID).Return(&entitySalePlanting, nil)

	responseSalePlanting, err := service.GetSalePlantingByID(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)

	assert.Nil(t, err)
	assert.Equal(t, entitySalePlanting.PlantingID, responseSalePlanting.PlantingID)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindSalePlantingByID", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)

}

func TestGetSalePlantingByID_ErrorIDNotExist(t *testing.T) {

	mockRepo, _, _, _, service, _, _, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("FindSalePlantingByID", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID).Return(&entities.SalePlantingEntity{}, fmt.Errorf("%w %d", myerror.ErrNotFoundSale, salePlantingMOCKID))

	responseSalePlanting, err := service.GetSalePlantingByID(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)

	assert.NotNil(t, err)
	assert.Nil(t, responseSalePlanting)
	assert.ErrorIs(t, err, myerror.ErrNotFoundSale)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindSalePlantingByID", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)
}

func TestGetSalePlantingByID_Error(t *testing.T) {

	mockRepo, _, _, _, service, _, _, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("FindSalePlantingByID", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID).Return(&entities.SalePlantingEntity{}, fmt.Errorf("ao buscar venda com id %d", salePlantingMOCKID))

	responseEntity, err := service.GetSalePlantingByID(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)

	assert.Nil(t, responseEntity)
	assert.Contains(t, err.Error(), fmt.Sprintf("ao buscar venda com id %d", salePlantingMOCKID))

	mockRepo.AssertCalled(t, "FindSalePlantingByID", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)
	mockRepo.AssertExpectations(t)
}

func TestPutSalePlanting_Success(t *testing.T) {

	mockRepo, _, _, _, service, entitySalePlanting, requestSalePlanting, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("UpdateSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, entitySalePlanting).Return(nil)

	err := service.PutSalePlanting(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, requestSalePlanting)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "UpdateSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, entitySalePlanting)
	mockRepo.AssertExpectations(t)
}

func TestPutSalePlanting_ConstraintViolated(t *testing.T) {

	mockRepo, _, _, _, service, entitySalePlanting, requestSalePlanting, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("UpdateSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, entitySalePlanting).Return(fmt.Errorf("%w", myerror.ErrDuplicateSale))

	err := service.PutSalePlanting(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, requestSalePlanting)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrDuplicateSale)

	mockRepo.AssertCalled(t, "UpdateSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, entitySalePlanting)
	mockRepo.AssertExpectations(t)
}

func TestPutSalePlanting_ViolatedForeignKey(t *testing.T) {

	mockRepo, _, _, _, service, entitySalePlanting, requestSalePlanting, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("UpdateSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, entitySalePlanting).Return(fmt.Errorf("%w", myerror.ErrViolatedForeingKey))

	err := service.PutSalePlanting(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, requestSalePlanting)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrViolatedForeingKey)

	mockRepo.AssertCalled(t, "UpdateSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, entitySalePlanting)
	mockRepo.AssertExpectations(t)

}

func TestPutSalePlanting_Error(t *testing.T) {

	mockRepo, _, _, _, service, entitySalePlanting, requestSalePlanting, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("UpdateSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, entitySalePlanting).Return(fmt.Errorf("erro ao atualizar venda"))

	err := service.PutSalePlanting(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, requestSalePlanting)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro ao atualizar venda")

	mockRepo.AssertCalled(t, "UpdateSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID, entitySalePlanting)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSalePlanting_Success(t *testing.T) {

	mockRepo, _, _, _, service, _, _, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("DeleteSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID).Return(nil)

	err := service.DeleteSalePlanting(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "DeleteSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSalePlanting_IDNotExists(t *testing.T) {

	mockRepo, _, _, _, service, _, _, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("DeleteSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID).Return(fmt.Errorf("%w", myerror.ErrNotFoundSale))

	err := service.DeleteSalePlanting(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrNotFoundSale)

	mockRepo.AssertCalled(t, "DeleteSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSalePlanting_Error(t *testing.T) {

	mockRepo, _, _, _, service, _, _, _, _, _ := SetupTestSalePlantingService()

	mockRepo.On("DeleteSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID).Return(fmt.Errorf("erro ao deletar venda"))

	err := service.DeleteSalePlanting(batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "ao deletar venda")

	mockRepo.AssertCalled(t, "DeleteSalePlanting", batchMOCKID, farmMOCKID, userMOCKID, salePlantingMOCKID)
	mockRepo.AssertExpectations(t)
}
