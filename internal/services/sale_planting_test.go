package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

func SetupTestSalePlantingService() (*mocks.SalePlantingRepositoryMock, SalePlantingService,
	entities.SalePlantingEntity, requests.SalePlantingRequest) {

	mockRepo := new(mocks.SalePlantingRepositoryMock)
	service := SalePlantingService{salePlantingRepository: mockRepo}

	entitySalePlanting := entities.SalePlantingEntity{
		PlantingID: 1,
		ValueSale:  500000000,
	}

	requestSalePlanting := requests.SalePlantingRequest{
		PlantingID: 1,
		ValueSale:  500000000,
	}

	return mockRepo, service, entitySalePlanting, requestSalePlanting
}
func TestPostSalePlanting_Success(t *testing.T) {

	mockRepo, service, entitySalePlanting, requestSalePlanting := SetupTestSalePlantingService()

	mockRepo.On("CreateSalePlantingRepository", entitySalePlanting).Return(nil)

	err := service.PostSalePlanting(requestSalePlanting)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "CreateSalePlantingRepository", entitySalePlanting)
}

func TestPostSalePlanting_Error(t *testing.T) {

	mockRepo, service, entitySalePlanting, requestSalePlanting := SetupTestSalePlantingService()

	mockRepo.On("CreateSalePlantingRepository", entitySalePlanting).Return(fmt.Errorf("erro ao cadastrar venda"))

	err := service.PostSalePlanting(requestSalePlanting)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro ao cadastrar venda")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "CreateSalePlantingRepository", entitySalePlanting)
}

func TestPostSalePlanting_Return_ConstraintViolatedPlantingID(t *testing.T) {

	mockRepo, service, entitySalePlanting, requestSalePlanting := SetupTestSalePlantingService()

	mockRepo.On("CreateSalePlantingRepository", entitySalePlanting).Return(fmt.Errorf("%w", myerror.ErrDuplicateSale))

	err := service.PostSalePlanting(requestSalePlanting)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrDuplicateSale)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "CreateSalePlantingRepository", entitySalePlanting)
}

func TestGetAllSalePlanting_Success(t *testing.T) {

	mockRepo, service, entitySalePlanting, _ := SetupTestSalePlantingService()

	var entitiesSalePlanting []entities.SalePlantingEntity
	entitiesSalePlanting = append(entitiesSalePlanting, entitySalePlanting)

	mockRepo.On("FindAllSalePlanting").Return(entitiesSalePlanting, nil)

	responsesSalePlanting, err := service.GetAllSalePlanting()

	assert.Nil(t, err)
	assert.Equal(t, entitiesSalePlanting[0].ValueSale, responsesSalePlanting[0].ValueSale)
	assert.Equal(t, entitiesSalePlanting[0].PlantingID, responsesSalePlanting[0].PlantingID)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindAllSalePlanting")

}

func TestGetAllSalePlanting_Error(t *testing.T) {

	mockRepo, service, _, _ := SetupTestSalePlantingService()

	mockRepo.On("FindAllSalePlanting").Return([]entities.SalePlantingEntity(nil), fmt.Errorf("não foi possível buscar todas as vendas de plantações"))

	responsesSalePlanting, err := service.GetAllSalePlanting()

	assert.NotNil(t, err)
	assert.Nil(t, responsesSalePlanting)
	assert.Contains(t, err.Error(), "não foi possível buscar todas as vendas de plantações")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindAllSalePlanting")
}

func TestGetSalePlantingByID_Success(t *testing.T) {

	mockRepo, service, entitySalePlanting, _ := SetupTestSalePlantingService()

	id := uint(1)
	entitySalePlanting.ID = id

	mockRepo.On("FindSalePlantingByID", id).Return(&entitySalePlanting, nil)

	responseSalePlanting, err := service.GetSalePlantingByID(id)

	assert.Nil(t, err)
	assert.Equal(t, entitySalePlanting.PlantingID, responseSalePlanting.PlantingID)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindSalePlantingByID", id)

}

func TestGetSalePlantingByID_ErrorIDNotExist(t *testing.T) {

	mockRepo, service, _, _ := SetupTestSalePlantingService()

	id := uint(2)

	mockRepo.On("FindSalePlantingByID", id).Return(&entities.SalePlantingEntity{}, fmt.Errorf("%w %d", myerror.ErrNotFoundSale, id))

	responseSalePlanting, err := service.GetSalePlantingByID(id)

	assert.NotNil(t, err)
	assert.Nil(t, responseSalePlanting)
	assert.ErrorIs(t, err, myerror.ErrNotFoundSale)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindSalePlantingByID", id)
}

func TestGetSalePlantingByID_Error(t *testing.T) {

	mockRepo, service, _, _ := SetupTestSalePlantingService()

	id := uint(1)

	mockRepo.On("FindSalePlantingByID", id).Return(&entities.SalePlantingEntity{}, fmt.Errorf("ao buscar venda com id %d", id))

	responseEntity, err := service.GetSalePlantingByID(id)

	assert.Nil(t, responseEntity)
	assert.Contains(t, err.Error(), fmt.Sprintf("ao buscar venda com id %d", id))

	mockRepo.AssertCalled(t, "FindSalePlantingByID", id)
	mockRepo.AssertExpectations(t)
}

func TestPutSalePlanting_Success(t *testing.T) {

	mockRepo, service, entitySalePlanting, requestSalePlanting := SetupTestSalePlantingService()

	id := uint(3)

	mockRepo.On("UpdateSalePlanting", id, entitySalePlanting).Return(nil)

	err := service.PutSalePlanting(id, requestSalePlanting)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "UpdateSalePlanting", id, entitySalePlanting)
	mockRepo.AssertExpectations(t)
}

func TestPutSalePlanting_ConstraintViolated(t *testing.T) {

	mockRepo, service, entitySalePlanting, requestSalePlanting := SetupTestSalePlantingService()

	id := uint(1)

	mockRepo.On("UpdateSalePlanting", id, entitySalePlanting).Return(fmt.Errorf("%w", myerror.ErrDuplicateSale))

	err := service.PutSalePlanting(id, requestSalePlanting)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrDuplicateSale)

	mockRepo.AssertCalled(t, "UpdateSalePlanting", id, entitySalePlanting)
	mockRepo.AssertExpectations(t)
}

func TestPutSalePlanting_ViolatedForeignKey(t *testing.T) {

	mockRepo, service, entitySalePlanting, requestSalePlanting := SetupTestSalePlantingService()

	id := uint(2)

	mockRepo.On("UpdateSalePlanting", id, entitySalePlanting).Return(fmt.Errorf("%w", myerror.ErrViolatedForeingKey))

	err := service.PutSalePlanting(id, requestSalePlanting)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrViolatedForeingKey)

	mockRepo.AssertCalled(t, "UpdateSalePlanting", id, entitySalePlanting)
	mockRepo.AssertExpectations(t)

}

func TestPutSalePlanting_Error(t *testing.T) {

	mockRepo, service, entitySalePlanting, requestSalePlanting := SetupTestSalePlantingService()

	id := uint(2)

	mockRepo.On("UpdateSalePlanting", id, entitySalePlanting).Return(fmt.Errorf("erro ao atualizar venda"))

	err := service.PutSalePlanting(id, requestSalePlanting)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro ao atualizar venda")

	mockRepo.AssertCalled(t, "UpdateSalePlanting", id, entitySalePlanting)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSalePlanting_Success(t *testing.T) {

	mockRepo, service, _, _ := SetupTestSalePlantingService()

	id := uint(5)

	mockRepo.On("DeleteSalePlanting", id).Return(nil)

	err := service.DeleteSalePlanting(id)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "DeleteSalePlanting", id)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSalePlanting_IDNotExists(t *testing.T) {

	mockRepo, service, _, _ := SetupTestSalePlantingService()

	id := uint(3)

	mockRepo.On("DeleteSalePlanting", id).Return(fmt.Errorf("%w", myerror.ErrNotFoundSale))

	err := service.DeleteSalePlanting(id)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrNotFoundSale)

	mockRepo.AssertCalled(t, "DeleteSalePlanting", id)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSalePlanting_Error(t *testing.T) {

	mockRepo, service, _, _ := SetupTestSalePlantingService()

	id := uint(3)

	mockRepo.On("DeleteSalePlanting", id).Return(fmt.Errorf("erro ao deletar venda"))

	err := service.DeleteSalePlanting(id)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "ao deletar venda")

	mockRepo.AssertCalled(t, "DeleteSalePlanting", id)
	mockRepo.AssertExpectations(t)
}
