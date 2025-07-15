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
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

const (
	perfomanceMockID = uint(1)
)

func SetupTestPerformancePlanting() (*mocks.PerformancePlantingRepository, PerformancePlantingService, entities.PerformancePlantingEntity,
	responses.DbResultPerformancePlanting, requests.PerformancePlantingRequest, *mocks.PlantingRepositoryMock) {

	mockRepo := new(mocks.PerformancePlantingRepository)

	mockRepoPlanting := new(mocks.PlantingRepositoryMock)

	service := PerformancePlantingService{performanceCultureRepository: mockRepo, plantingRepository: mockRepoPlanting}

	entityPerformancePlanting := entities.PerformancePlantingEntity{
		PlantingID:             plantingMOCKID,
		ProductionObtained:     500000,
		UnitProductionObtained: "kg",
		HarvestedArea:          300,
		HarvestedDate:          time.Date(2025, time.May, 3, 0, 0, 0, 0, time.UTC),
		UnitHarvestedArea:      "ha",
	}

	dbResultPlanting := responses.DbResultPerformancePlanting{
		BatchName:                  "lote14",
		AgricultureCultureName:     "milho",
		StartDatePlanting:          time.Now().AddDate(2025, 4, 3),
		IsPlanting:                 false,
		ProductionObtained:         500000,
		ProductionObtainedFormated: "500000kg",
		HarvestedArea:              200,
		HarvestedAreaFormated:      "200ha",
		HarvestedDate:              time.Date(2025, time.May, 3, 0, 0, 0, 0, time.UTC),
	}

	requestPerfomancePlanting := requests.PerformancePlantingRequest{
		ProductionObtained:     500000,
		UnitProductionObtained: "kg",
		HarvestedArea:          300,
		HarvestedDate:          time.Date(2025, time.May, 3, 0, 0, 0, 0, time.UTC),
		UnitHarvestedArea:      "ha",
	}

	return mockRepo, service, entityPerformancePlanting, dbResultPlanting, requestPerfomancePlanting, mockRepoPlanting
}

func TestPostPerformancePlanting_Success(t *testing.T) {

	mockRepo, service, perfomoancePlantingEntity, _, perfomancePlantingRequest, mockeRepoPlanting := SetupTestPerformancePlanting()

	mockeRepoPlanting.On("FindPlantingByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(entities.PlantingEntity{}, nil)

	mockRepo.On("CreatePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomoancePlantingEntity).Return(nil)

	err := service.PostPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomancePlantingRequest)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "CreatePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomoancePlantingEntity)
	mockRepo.AssertExpectations(t)

}

func TestPostPerformancePlanting_PlantingNotFound(t *testing.T) {

	mockRepo, service, _, _, perfomancePlantingRequest, mockeRepoPlanting := SetupTestPerformancePlanting()

	mockeRepoPlanting.On("FindPlantingByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(entities.PlantingEntity{}, errors.New("erro"))

	err := service.PostPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomancePlantingRequest)

	assert.NotNil(t, err)
	mockRepo.AssertExpectations(t)

}

func TestPostPerformanceCulture_ValidateUnitProductionObtainedEnumError(t *testing.T) {

	_, service, _, _, perfomancePlantingRequest, _ := SetupTestPerformancePlanting()

	perfomancePlantingRequest.UnitProductionObtained = "teste"

	err := service.PostPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomancePlantingRequest)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrEnumInvalid)
	assert.ErrorContains(t, err, "o campo unit_production_obtained")
}

func TestPostPerformanceCulture_ValidateUnitHarvestedAreaEnumError(t *testing.T) {

	_, service, _, _, perfomancePlantingRequest, _ := SetupTestPerformancePlanting()

	perfomancePlantingRequest.UnitHarvestedArea = "teste"

	err := service.PostPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomancePlantingRequest)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrEnumInvalid)
	assert.ErrorContains(t, err, "o campo unit_harvested_area")
}

func TestPostPerfomanceCulture_Error(t *testing.T) {

	mockRepo, service, entityPerformancePlanting, _, perfomancePlantingRequest, mockeRepoPlanting := SetupTestPerformancePlanting()

	mockeRepoPlanting.On("FindPlantingByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(entities.PlantingEntity{}, nil)

	mockRepo.On("CreatePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, entityPerformancePlanting).Return(fmt.Errorf("erro ao cadastrar performance da cultura"))

	err := service.PostPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomancePlantingRequest)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro ao cadastrar performance da cultura")

	mockRepo.AssertCalled(t, "CreatePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, entityPerformancePlanting)
	mockRepo.AssertExpectations(t)
}

func TestGetAllPerformancePlanting_Success(t *testing.T) {

	mockRepo, service, _, dbResultPlanting, _, mockeRepoPlanting := SetupTestPerformancePlanting()

	var listDbResultPlanting []responses.DbResultPerformancePlanting
	listDbResultPlanting = append(listDbResultPlanting, dbResultPlanting)

	mockeRepoPlanting.On("FindPlantingByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(entities.PlantingEntity{}, nil)

	mockRepo.On("FindAll", batchMOCKID, farmMOCKID, userMOCKID).Return(listDbResultPlanting, nil)

	result, err := service.GetAllPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID)

	assert.Nil(t, err)

	for i := range result {
		assert.Equal(t, listDbResultPlanting[i].BatchName, result[i].Planting.BatchName)
		assert.Equal(t, listDbResultPlanting[i].AgricultureCultureName, result[i].Planting.AgricultureCultureName)
		assert.Equal(t, listDbResultPlanting[i].IsPlanting, result[i].Planting.IsPlanting)
		assert.Equal(t, listDbResultPlanting[i].StartDatePlanting, result[i].Planting.StartDatePlanting)
		assert.Equal(t, listDbResultPlanting[i].HarvestedArea, result[i].HarvestedArea)
		assert.Equal(t, listDbResultPlanting[i].HarvestedAreaFormated, result[i].HarvestedAreaFormated)
		assert.Equal(t, listDbResultPlanting[i].HarvestedDate, result[i].HarvestedDate)
		assert.Equal(t, listDbResultPlanting[i].ProductionObtained, result[i].ProductionObtained)
		assert.Equal(t, listDbResultPlanting[i].ProductionObtainedFormated, result[i].ProductionObtainedFormated)
	}

	mockRepo.AssertCalled(t, "FindAll", batchMOCKID, farmMOCKID, userMOCKID)
	mockRepo.AssertExpectations(t)
}

func TestGetAllPerformancePlanting_Error(t *testing.T) {

	mockRepo, service, _, _, _, _ := SetupTestPerformancePlanting()

	mockRepo.On("FindAll", batchMOCKID, farmMOCKID, userMOCKID).Return([]responses.DbResultPerformancePlanting(nil), fmt.Errorf("erro ao buscar performance de plantação"))

	result, err := service.GetAllPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro ao buscar performance de plantação")

	mockRepo.AssertCalled(t, "FindAll", batchMOCKID, farmMOCKID, userMOCKID)
	mockRepo.AssertExpectations(t)
}

func TestGetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID_Success(t *testing.T) {

	mockRepo, service, _, dbResultPlanting, _, _ := SetupTestPerformancePlanting()

	mockRepo.On("FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID).Return(&dbResultPlanting, nil)

	response, err := service.GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID)

	assert.Nil(t, err)
	assert.Equal(t, dbResultPlanting.BatchName, response.Planting.BatchName)
	assert.Equal(t, dbResultPlanting.AgricultureCultureName, response.Planting.AgricultureCultureName)
	assert.Equal(t, dbResultPlanting.IsPlanting, response.Planting.IsPlanting)
	assert.Equal(t, dbResultPlanting.StartDatePlanting, response.Planting.StartDatePlanting)
	assert.Equal(t, dbResultPlanting.HarvestedArea, response.HarvestedArea)
	assert.Equal(t, dbResultPlanting.HarvestedAreaFormated, response.HarvestedAreaFormated)
	assert.Equal(t, dbResultPlanting.HarvestedDate, response.HarvestedDate)
	assert.Equal(t, dbResultPlanting.ProductionObtained, response.ProductionObtained)
	assert.Equal(t, dbResultPlanting.ProductionObtainedFormated, response.ProductionObtainedFormated)
	assert.Equal(t, dbResultPlanting.StartDatePlanting, response.Planting.StartDatePlanting)

	mockRepo.AssertCalled(t, "FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID)
	mockRepo.AssertExpectations(t)
}

func TestGetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID_Error(t *testing.T) {

	mockRepo, service, _, _, _, _ := SetupTestPerformancePlanting()

	mockRepo.On("FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID).Return(&responses.DbResultPerformancePlanting{}, errors.New("erro"))

	response, err := service.GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.ErrorContains(t, err, "erro")

	mockRepo.AssertCalled(t, "FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID)
	mockRepo.AssertExpectations(t)
}

func TestPutPerformancePlanting_Success(t *testing.T) {

	mockRepo, service, entityPerformancePlanting, _, requestPerformancePlanting, _ := SetupTestPerformancePlanting()
	entityPerformancePlanting.PlantingID = plantingMOCKID
	entityPerformancePlanting.ID = perfomanceMockID

	mockRepo.On("UpdatePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID, entityPerformancePlanting).Return(nil)

	err := service.PutPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID, requestPerformancePlanting)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "UpdatePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID, entityPerformancePlanting)
	mockRepo.AssertExpectations(t)
}

func TestPutPerformancePlanting_Error(t *testing.T) {

	mockRepo, service, entityPerformancePlanting, _, requestPerformancePlanting, _ := SetupTestPerformancePlanting()
	entityPerformancePlanting.ID = perfomanceMockID
	entityPerformancePlanting.PlantingID = plantingMOCKID

	mockRepo.On("UpdatePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID, entityPerformancePlanting).Return(errors.New("erro"))

	err := service.PutPerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID, requestPerformancePlanting)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro")

	mockRepo.AssertCalled(t, "UpdatePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID, entityPerformancePlanting)
	mockRepo.AssertExpectations(t)

}

func TestDeletePerformancePlanting_Success(t *testing.T) {

	mockRepo, service, _, _, _, _ := SetupTestPerformancePlanting()

	mockRepo.On("DeletePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID).Return(nil)

	err := service.DeletePerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "DeletePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID)
	mockRepo.AssertExpectations(t)

}

func TestDeletePerformancePlanting_Error(t *testing.T) {

	mockRepo, service, _, _, _, _ := SetupTestPerformancePlanting()

	mockRepo.On("DeletePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID).Return(errors.New("error"))

	err := service.DeletePerformancePlanting(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "error")

	mockRepo.AssertCalled(t, "DeletePerformancePlanting", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID, perfomanceMockID)
	mockRepo.AssertExpectations(t)

}
