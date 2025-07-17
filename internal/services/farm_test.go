package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
)

func SetupTestFarm() (*mocks.FarmRepositoryMock, FarmService, entities.FarmEntity, responses.FarmResponse,
	requests.FarmRequest, context.Context) {

	mockRepoFarm := new(mocks.FarmRepositoryMock)

	service := FarmService{mockRepoFarm}

	entityFarm := entities.FarmEntity{
		Name:   "teste",
		UserID: userMOCKID,
	}

	requestFarm := requests.FarmRequest{
		Name:   "teste",
		UserID: userMOCKID,
	}

	responseFarm := responses.FarmResponse{
		ID:   farmMOCKID,
		Name: "teste",
	}

	return mockRepoFarm, service, entityFarm, responseFarm, requestFarm, context.Background()

}

func TestPostFarm_Success(t *testing.T) {

	mockRepo, service, entityFarm, _, requestFarm, ctx := SetupTestFarm()

	mockRepo.On("Create", ctx, entityFarm).Return(nil)

	err := service.Create(ctx, requestFarm)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "Create", ctx, entityFarm)
	mockRepo.AssertExpectations(t)

}

func TestPostFarm_Error(t *testing.T) {

	mockRepo, service, entityFarm, _, requestFarm, ctx := SetupTestFarm()

	mockRepo.On("Create", ctx, entityFarm).Return(errors.New("erro"))

	err := service.Create(ctx, requestFarm)

	assert.NotNil(t, err)

	mockRepo.AssertCalled(t, "Create", ctx, entityFarm)
	mockRepo.AssertExpectations(t)

}

func TestGetAllFarm_Success(t *testing.T) {

	mockRepo, service, _, responseFarm, _, _ := SetupTestFarm()

	var listResponseFarm []responses.FarmResponse
	listResponseFarm = append(listResponseFarm, responseFarm)

	mockRepo.On("FindAll", userMOCKID).Return(listResponseFarm, nil)

	listResponse, err := service.GetAllFarm(userMOCKID)

	assert.Nil(t, err)

	for i := range listResponseFarm {
		assert.Equal(t, listResponseFarm[i].ID, listResponse[i].ID)
		assert.Equal(t, listResponseFarm[i].Name, listResponse[i].Name)
	}

	mockRepo.AssertCalled(t, "FindAll", userMOCKID)
	mockRepo.AssertExpectations(t)

}

func TestGetAllFarm_Error(t *testing.T) {

	mockRepo, service, _, _, _, _ := SetupTestFarm()

	mockRepo.On("FindAll", userMOCKID).Return([]responses.FarmResponse{}, errors.New("erro"))

	listResponse, err := service.GetAllFarm(userMOCKID)

	assert.NotNil(t, err)
	assert.Nil(t, listResponse)

	mockRepo.AssertCalled(t, "FindAll", userMOCKID)
	mockRepo.AssertExpectations(t)

}
