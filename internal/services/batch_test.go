package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
)

func SetupTestBatch() (*mocks.BatchRepositoryMock, BatchService, entities.BatchEntity, responses.BatchResponse, requests.BatchRequest) {

	mockRepo := new(mocks.BatchRepositoryMock)

	service := BatchService{mockRepo}

	batchEntity := entities.BatchEntity{
		Name:   "teste",
		Area:   5000,
		Unit:   "ha",
		FarmID: farmMOCKID,
	}

	batchRequest := requests.BatchRequest{
		Name: "teste",
		Area: 5000,
		Unit: "ha",
	}

	batchResponse := responses.BatchResponse{

		Name: "teste",
		Area: 5000,
		Unit: "ha",
	}

	return mockRepo, service, batchEntity, batchResponse, batchRequest

}

func TestPostBatch_Success(t *testing.T) {

	mockRepo, service, batchEntity, _, batchRequest := SetupTestBatch()

	mockRepo.On("Create", userMOCKID, farmMOCKID, batchEntity).Return(nil)

	err := service.PostBatchService(userMOCKID, farmMOCKID, batchRequest)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "Create", userMOCKID, farmMOCKID, batchEntity)
	mockRepo.AssertExpectations(t)

}

func TestPostBatch_Error(t *testing.T) {

	mockRepo, service, batchEntity, _, batchRequest := SetupTestBatch()

	mockRepo.On("Create", userMOCKID, farmMOCKID, batchEntity).Return(errors.New("erro"))

	err := service.PostBatchService(userMOCKID, farmMOCKID, batchRequest)

	assert.NotNil(t, err)

	mockRepo.AssertCalled(t, "Create", userMOCKID, farmMOCKID, batchEntity)
	mockRepo.AssertExpectations(t)

}

func TestGetAllBatch_Success(t *testing.T) {

	mockRepo, service, batchEntity, _, _ := SetupTestBatch()

	var listEntityBatch []entities.BatchEntity
	listEntityBatch = append(listEntityBatch, batchEntity)

	mockRepo.On("FindAllBatch", userMOCKID, farmMOCKID).Return(listEntityBatch, nil)

	listResponseBatch, err := service.GetAllBatch(userMOCKID, farmMOCKID)

	assert.Nil(t, err)

	for i := range listEntityBatch {
		assert.Equal(t, listEntityBatch[i].Name, listResponseBatch[i].Name)
		assert.Equal(t, listEntityBatch[i].Area, listResponseBatch[i].Area)
		assert.Equal(t, listEntityBatch[i].Unit, listResponseBatch[i].Unit)
	}

	mockRepo.AssertCalled(t, "FindAllBatch", userMOCKID, farmMOCKID)
	mockRepo.AssertExpectations(t)

}

func TestGetAllBatch_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestBatch()

	mockRepo.On("FindAllBatch", userMOCKID, farmMOCKID).Return([]entities.BatchEntity{}, errors.New("erro"))

	listResponseBatch, err := service.GetAllBatch(userMOCKID, farmMOCKID)

	assert.NotNil(t, err)
	assert.Nil(t, listResponseBatch)

	mockRepo.AssertCalled(t, "FindAllBatch", userMOCKID, farmMOCKID)
	mockRepo.AssertExpectations(t)

}

func TestPutBatch_Success(t *testing.T) {

	mockRepo, service, batchEntity, _, batchRequest := SetupTestBatch()

	batchEntity.FarmID = 0

	mockRepo.On("Update", userMOCKID, farmMOCKID, batchMOCKID, batchEntity).Return(nil)

	err := service.PutBatch(userMOCKID, farmMOCKID, batchMOCKID, batchRequest)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "Update", userMOCKID, farmMOCKID, batchMOCKID, batchEntity)
	mockRepo.AssertExpectations(t)

}

func TestPutBatch_Error(t *testing.T) {

	mockRepo, service, batchEntity, _, batchRequest := SetupTestBatch()

	batchEntity.FarmID = 0

	mockRepo.On("Update", userMOCKID, farmMOCKID, batchMOCKID, batchEntity).Return(errors.New("erro"))

	err := service.PutBatch(userMOCKID, farmMOCKID, batchMOCKID, batchRequest)

	assert.NotNil(t, err)

	mockRepo.AssertCalled(t, "Update", userMOCKID, farmMOCKID, batchMOCKID, batchEntity)
	mockRepo.AssertExpectations(t)

}

func TestDeleteBatch_Success(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestBatch()

	mockRepo.On("DeleteBatch", userMOCKID, farmMOCKID, batchMOCKID).Return(nil)

	err := service.DeleteBatch(userMOCKID, farmMOCKID, batchMOCKID)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "DeleteBatch", userMOCKID, farmMOCKID, batchMOCKID)
	mockRepo.AssertExpectations(t)

}

func TestDeleteBatch_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestBatch()

	mockRepo.On("DeleteBatch", userMOCKID, farmMOCKID, batchMOCKID).Return(errors.New("erro"))

	err := service.DeleteBatch(userMOCKID, farmMOCKID, batchMOCKID)

	assert.NotNil(t, err)

	mockRepo.AssertCalled(t, "DeleteBatch", userMOCKID, farmMOCKID, batchMOCKID)
	mockRepo.AssertExpectations(t)

}
