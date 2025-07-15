package services

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

const (
	userMOCKID     = uint(1)
	farmMOCKID     = uint(1)
	batchMOCKID    = uint(1)
	plantingMOCKID = uint(1)
)

func SetupTestPlantingService() (*mocks.PlantingRepositoryMock, PlantingService, entities.PlantingEntity, responses.PlantingResponse, requests.PlantingRequest) {

	mockRepo := new(mocks.PlantingRepositoryMock)

	service := PlantingService{plantingRepository: mockRepo}

	timeCurrent := time.Date(2025, time.July, 3, 18, 19, 28, 674505796, time.Local)

	entityPlanting := entities.PlantingEntity{
		BatchID:              batchMOCKID,
		AgricultureCultureID: uint(5),
		IsPlanting:           false,
		StartDatePlanting:    timeCurrent,
		ExpectedProduction:   0,
		SpaceBetweenPlants:   0.50,
		SpaceBetweenRows:     0.30,
		IrrigationTypeID:     uint(4),
	}

	ativo := false

	requestPlanting := requests.PlantingRequest{
		AgricultureCultureID: uint(5),
		IsPlanting:           &ativo,
		StartDatePlanting:    timeCurrent,
		SpaceBetweenPlants:   0.50,
		SpaceBetweenRows:     0.30,
		IrrigationTypeID:     uint(4),
		ExpectedProduction:   0,
	}

	responsePlanting := responses.PlantingResponse{
		BatchID:              batchMOCKID,
		AgricultureCultureID: uint(5),
		IsPlanting:           false,
		StartDatePlanting:    timeCurrent,
		SpaceBetweenPlants:   0.50,
		SpaceBetweenRows:     0.30,
		IrrigationTypeID:     uint(4),
	}

	return mockRepo, service, entityPlanting, responsePlanting, requestPlanting

}

func TestPostPlanting_Success(t *testing.T) {

	mockRepo, service, _, _, requestPlanting := SetupTestPlantingService()

	mockRepo.On("FindByParamPlanting", userMOCKID, farmMOCKID, batchMOCKID).Return(entities.PlantingEntity{}, myerror.ErrNotFound)

	planting, err := service.GetByParam(userMOCKID, farmMOCKID, batchMOCKID)

	assert.ErrorIs(t, err, myerror.ErrNotFound)
	assert.Nil(t, planting)

	matchPlanting := mock.MatchedBy(func(e entities.PlantingEntity) bool {
		return e.AgricultureCultureID == requestPlanting.AgricultureCultureID &&
			e.BatchID == batchMOCKID &&
			e.SpaceBetweenPlants == requestPlanting.SpaceBetweenPlants &&
			e.SpaceBetweenRows == requestPlanting.SpaceBetweenRows &&
			e.IrrigationTypeID == requestPlanting.IrrigationTypeID &&
			e.IsPlanting == *requestPlanting.IsPlanting &&
			e.ExpectedProduction == requestPlanting.ExpectedProduction
		// e.StartDatePlanting é ignorado na comparação
	})
	mockRepo.On("CreatePlanting", matchPlanting).Return(nil)

	err = service.PostPlanting(userMOCKID, farmMOCKID, batchMOCKID, requestPlanting)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "FindByParamPlanting", userMOCKID, farmMOCKID, batchMOCKID)
	mockRepo.AssertCalled(t, "CreatePlanting", matchPlanting)
	mockRepo.AssertExpectations(t)
}

func TestPostPlanting_ErrorBatchInUse(t *testing.T) {

	mockRepo, service, entityPlanting, _, requestPlanting := SetupTestPlantingService()

	entityPlanting.ID = uint(1)
	entityPlanting.IsPlanting = true

	mockRepo.On("FindByParamPlanting", userMOCKID, farmMOCKID, entityPlanting.BatchID).Return(entityPlanting, nil)

	response, err := service.GetByParam(userMOCKID, farmMOCKID, batchMOCKID)

	assert.Equal(t, uint(1), response.ID)
	assert.Nil(t, err)

	err = service.PostPlanting(userMOCKID, farmMOCKID, batchMOCKID, requestPlanting)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "ao cadastrar plantação. Lote já está sendo utilizado pela cultura")

	mockRepo.AssertCalled(t, "FindByParamPlanting", userMOCKID, farmMOCKID, entityPlanting.BatchID)

	mockRepo.AssertExpectations(t)

}

func TestPostPlanting_ErrorNotFound(t *testing.T) {

	mockRepo, service, entityPlanting, _, requestPlanting := SetupTestPlantingService()

	mockRepo.On("FindByParamPlanting", userMOCKID, farmMOCKID, entityPlanting.BatchID).Return(entities.PlantingEntity{}, errors.New("erro"))

	response, err := service.GetByParam(userMOCKID, farmMOCKID, batchMOCKID)

	assert.Nil(t, response)
	assert.NotNil(t, err)

	err = service.PostPlanting(userMOCKID, farmMOCKID, batchMOCKID, requestPlanting)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro")

	mockRepo.AssertCalled(t, "FindByParamPlanting", userMOCKID, farmMOCKID, entityPlanting.BatchID)

	mockRepo.AssertExpectations(t)

}

// func TestPostPlanting_Error(t *testing.T) {

// 	mockRepo, service, entityPlanting, _, requestPlanting := SetupTestPlantingService()

// 	mockRepo.On("FindByParamPlanting", userMOCKID, farmMOCKID, batchMOCKID).Return(entities.PlantingEntity{}, nil)

// 	_, err := service.GetByParam(userMOCKID, farmMOCKID, batchMOCKID)

// 	assert.Nil(t, err)

// 	mockRepo.On("CreatePlanting", entityPlanting).Return(fmt.Errorf("erro"))

// 	err = service.PostPlanting(userMOCKID, farmMOCKID, batchMOCKID, requestPlanting)

// 	assert.NotNil(t, err)

// 	mockRepo.AssertCalled(t, "FindByParamPlanting", userMOCKID, farmMOCKID, batchMOCKID)

// 	mockRepo.AssertCalled(t, "CreatePlanting", entityPlanting)

// }

func TestGetByParam_Success(t *testing.T) {

	mockRepo, service, entityPlanting, _, _ := SetupTestPlantingService()

	mockRepo.On("FindByParamPlanting", userMOCKID, farmMOCKID, batchMOCKID).Return(entityPlanting, nil)

	response, err := service.GetByParam(userMOCKID, farmMOCKID, batchMOCKID)

	assert.Nil(t, err)
	assert.Equal(t, entityPlanting.ID, response.ID)
	assert.Equal(t, entityPlanting.BatchID, response.BatchID)
	assert.Equal(t, entityPlanting.ExpectedProduction, response.ExpectedProduction)
	assert.EqualValues(t, entityPlanting.IsPlanting, response.IsPlanting)
	assert.Equal(t, entityPlanting.IrrigationTypeID, response.IrrigationTypeID)
	assert.Equal(t, entityPlanting.SpaceBetweenPlants, response.SpaceBetweenPlants)
	assert.Equal(t, entityPlanting.AgricultureCultureID, response.AgricultureCultureID)
	assert.Equal(t, entityPlanting.SpaceBetweenRows, response.SpaceBetweenRows)

	mockRepo.AssertCalled(t, "FindByParamPlanting", userMOCKID, farmMOCKID, batchMOCKID)

	mockRepo.AssertExpectations(t)

}

func TestGetByParam_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPlantingService()

	mockRepo.On("FindByParamPlanting", userMOCKID, farmMOCKID, batchMOCKID).Return(entities.PlantingEntity{}, fmt.Errorf("erro ao buscar objeto"))

	response, err := service.GetByParam(userMOCKID, farmMOCKID, batchMOCKID)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro ao buscar objeto")
}

func TestGetAllPlanting_Success(t *testing.T) {

	mockRepo, service, entityPlanting, _, _ := SetupTestPlantingService()

	var entitiesPlantings []entities.PlantingEntity

	entitiesPlantings = append(entitiesPlantings, entityPlanting)

	mockRepo.On("FindAllPlanting", userMOCKID, farmMOCKID, batchMOCKID).Return(entitiesPlantings, nil)

	responsesPlantings, err := service.GetAllPlanting(batchMOCKID, farmMOCKID, userMOCKID)

	for i := range responsesPlantings {
		assert.Equal(t, entitiesPlantings[i].AgricultureCultureID, responsesPlantings[i].AgricultureCultureID)
		assert.Equal(t, entitiesPlantings[i].BatchID, responsesPlantings[i].BatchID)
		assert.Equal(t, entitiesPlantings[i].ExpectedProduction, responsesPlantings[i].ExpectedProduction)
		assert.Equal(t, entitiesPlantings[i].ID, responsesPlantings[i].ID)
		assert.Equal(t, entitiesPlantings[i].IrrigationTypeID, responsesPlantings[i].IrrigationTypeID)
		assert.Equal(t, entitiesPlantings[i].IsPlanting, responsesPlantings[i].IsPlanting)
		assert.Equal(t, entitiesPlantings[i].SpaceBetweenPlants, responsesPlantings[i].SpaceBetweenPlants)
		assert.Equal(t, entitiesPlantings[i].SpaceBetweenRows, responsesPlantings[i].SpaceBetweenRows)
	}

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "FindAllPlanting", userMOCKID, farmMOCKID, batchMOCKID)

	mockRepo.AssertExpectations(t)

}

func TestGetAllPlanting_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPlantingService()

	mockRepo.On("FindAllPlanting", userMOCKID, farmMOCKID, batchMOCKID).Return([]entities.PlantingEntity{}, fmt.Errorf("erro ao buscar todas as plantações"))

	responsesPlantings, err := service.GetAllPlanting(batchMOCKID, farmMOCKID, userMOCKID)

	assert.Nil(t, responsesPlantings)
	assert.ErrorContains(t, err, "erro ao buscar todas as plantações")

	mockRepo.AssertCalled(t, "FindAllPlanting", userMOCKID, farmMOCKID, batchMOCKID)
	mockRepo.AssertExpectations(t)

}

func TestGetByParamBatchNameOrIsActivePlanting_Success(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPlantingService()

	isActive := true
	batchName := "lote 15"

	batchPlantingReponse := responses.BatchPlantiesResponse{
		BatchName:              batchName,
		IsPlanting:             isActive,
		AgricultureCultureName: "milho",
		SoilTypeName:           "argiloso",
		StartDatePlanting:      time.Now(),
		SpaceBetweenPlants:     0.50,
		SpaceBetweenRows:       0.30,
		IrrigationType:         "aspersao",
	}

	var listBatchPlantingResponse []responses.BatchPlantiesResponse
	listBatchPlantingResponse = append(listBatchPlantingResponse, batchPlantingReponse)

	mockRepo.On("FindByParamBatchNameOrIsActivePlanting", batchName, isActive, userMOCKID, farmMOCKID).Return(listBatchPlantingResponse, nil)

	responsesBatchPlanting, err := service.GetByParamBatchNameOrIsActivePlanting(batchName, isActive, userMOCKID, farmMOCKID)

	assert.Nil(t, err)

	for i := range responsesBatchPlanting {
		assert.Equal(t, listBatchPlantingResponse[i].AgricultureCultureName, responsesBatchPlanting[i].AgricultureCultureName)
		assert.Equal(t, listBatchPlantingResponse[i].BatchName, responsesBatchPlanting[i].BatchName)
		assert.Equal(t, listBatchPlantingResponse[i].IrrigationType, responsesBatchPlanting[i].IrrigationType)
		assert.Equal(t, listBatchPlantingResponse[i].IsPlanting, responsesBatchPlanting[i].IsPlanting)
		assert.Equal(t, listBatchPlantingResponse[i].SoilTypeName, responsesBatchPlanting[i].SoilTypeName)
		assert.Equal(t, listBatchPlantingResponse[i].SpaceBetweenPlants, responsesBatchPlanting[i].SpaceBetweenPlants)
		assert.Equal(t, listBatchPlantingResponse[i].SpaceBetweenRows, responsesBatchPlanting[i].SpaceBetweenRows)
	}

	mockRepo.AssertCalled(t, "FindByParamBatchNameOrIsActivePlanting", batchName, isActive, userMOCKID, farmMOCKID)
	mockRepo.AssertExpectations(t)

}

func TestGetByParamBatchNameOrIsActivePlanting_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPlantingService()

	isActive := true
	batchName := "lote 15"

	mockRepo.On("FindByParamBatchNameOrIsActivePlanting", batchName, isActive, userMOCKID, farmMOCKID).Return([]responses.BatchPlantiesResponse{}, fmt.Errorf("erro ao buscar dados"))

	responsesBatchPlanting, err := service.GetByParamBatchNameOrIsActivePlanting(batchName, isActive, userMOCKID, farmMOCKID)

	assert.Nil(t, responsesBatchPlanting)
	assert.ErrorContains(t, err, "ao buscar dados")

	mockRepo.AssertCalled(t, "FindByParamBatchNameOrIsActivePlanting", batchName, isActive, userMOCKID, farmMOCKID)
	mockRepo.AssertExpectations(t)

}

func TestPutPlanting_Success(t *testing.T) {

	mockRepo, service, _, _, requestPlanting := SetupTestPlantingService()

	matchPlanting := mock.MatchedBy(func(e entities.PlantingEntity) bool {
		return e.AgricultureCultureID == requestPlanting.AgricultureCultureID &&
			e.SpaceBetweenPlants == requestPlanting.SpaceBetweenPlants &&
			e.SpaceBetweenRows == requestPlanting.SpaceBetweenRows &&
			e.IrrigationTypeID == requestPlanting.IrrigationTypeID &&
			e.IsPlanting == *requestPlanting.IsPlanting &&
			e.ExpectedProduction == requestPlanting.ExpectedProduction
	})

	mockRepo.On("UpdatePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, matchPlanting).Return(nil)

	err := service.PutPlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, requestPlanting)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "UpdatePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, matchPlanting)
	mockRepo.AssertExpectations(t)

}

func TestPutPlanting_Error(t *testing.T) {

	mockRepo, service, entityPlanting, _, requestPlanting := SetupTestPlantingService()

	entityPlanting.ID = plantingMOCKID

	matchPlanting := mock.MatchedBy(func(e entities.PlantingEntity) bool {
		return e.AgricultureCultureID == requestPlanting.AgricultureCultureID &&
			e.SpaceBetweenPlants == requestPlanting.SpaceBetweenPlants &&
			e.SpaceBetweenRows == requestPlanting.SpaceBetweenRows &&
			e.IrrigationTypeID == requestPlanting.IrrigationTypeID &&
			e.IsPlanting == *requestPlanting.IsPlanting &&
			e.ExpectedProduction == requestPlanting.ExpectedProduction
	})

	mockRepo.On("UpdatePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, matchPlanting).Return(fmt.Errorf("erro: erro ao atualilzar plantação"))

	err := service.PutPlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, requestPlanting)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "ao atualilzar plantação")

	mockRepo.AssertCalled(t, "UpdatePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, matchPlanting)
	mockRepo.AssertExpectations(t)

}

func TestDeletePlanting_Success(t *testing.T) {

	mockRepo, service, entitiesPlanting, _, _ := SetupTestPlantingService()

	id := uint(1)
	entitiesPlanting.ID = id

	mockRepo.On("DeletePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(nil)

	err := service.DeletePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "DeletePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)
	mockRepo.AssertExpectations(t)
}

func TestDeletePlanting_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPlantingService()

	mockRepo.On("DeletePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(fmt.Errorf("erro ao tentar deletar plantação"))

	err := service.DeletePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "ao tentar deletar plantação")

	mockRepo.AssertCalled(t, "DeletePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)
	mockRepo.AssertExpectations(t)

}
