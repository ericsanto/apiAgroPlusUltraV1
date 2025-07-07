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

func SetupTestPerformancePlanting() (*mocks.PerformancePlantingRepository, PerformancePlantingService, entities.PerformancePlantingEntity,
	responses.DbResultPerformancePlanting, requests.PerformancePlantingRequest) {

	mockRepo := new(mocks.PerformancePlantingRepository)

	service := PerformancePlantingService{performanceCultureRepository: mockRepo}

	entityPerformancePlanting := entities.PerformancePlantingEntity{
		PlantingID:             uint(1),
		ProductionObtained:     500000,
		UnitProductionObtained: "kg",
		HarvestedArea:          300,
		HarvestedDate:          time.Date(2025, time.May, 3, 0, 0, 0, 0, time.UTC),
		UnitHarvestedArea:      "ha",
	}

	dbResultPlanting := responses.DbResultPerformancePlanting{
		PlantingID:                 uint(1),
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
		PlantingID:             uint(1),
		ProductionObtained:     500000,
		UnitProductionObtained: "kg",
		HarvestedArea:          300,
		HarvestedDate:          time.Date(2025, time.May, 3, 0, 0, 0, 0, time.UTC),
		UnitHarvestedArea:      "ha",
	}

	return mockRepo, service, entityPerformancePlanting, dbResultPlanting, requestPerfomancePlanting
}

func TestPostPerformancePlanting_Success(t *testing.T) {

	mockRepo, service, perfomoancePlantingEntity, _, perfomancePlantingRequest := SetupTestPerformancePlanting()

	mockRepo.On("CreatePerformancePlanting", perfomoancePlantingEntity).Return(nil)

	err := service.PostPerformancePlanting(perfomancePlantingRequest)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "CreatePerformancePlanting", perfomoancePlantingEntity)
	mockRepo.AssertExpectations(t)

}

func TestPostPerformanceCulture_ValidateUnitProductionObtainedEnumError(t *testing.T) {

	_, service, _, _, perfomancePlantingRequest := SetupTestPerformancePlanting()

	perfomancePlantingRequest.UnitProductionObtained = "teste"

	err := service.PostPerformancePlanting(perfomancePlantingRequest)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrEnumInvalid)
	assert.ErrorContains(t, err, "o campo unit_production_obtained")
}

func TestPostPerformanceCulture_ValidateUnitHarvestedAreaEnumError(t *testing.T) {

	_, service, _, _, perfomancePlantingRequest := SetupTestPerformancePlanting()

	perfomancePlantingRequest.UnitHarvestedArea = "teste"

	err := service.PostPerformancePlanting(perfomancePlantingRequest)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, myerror.ErrEnumInvalid)
	assert.ErrorContains(t, err, "o campo unit_harvested_area")
}

func TestPostPerfomanceCulture_Error(t *testing.T) {

	mockRepo, service, entityPerformancePlanting, _, perfomancePlantingRequest := SetupTestPerformancePlanting()

	mockRepo.On("CreatePerformancePlanting", entityPerformancePlanting).Return(fmt.Errorf("erro ao cadastrar performance da cultura"))

	err := service.PostPerformancePlanting(perfomancePlantingRequest)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro ao cadastrar performance da cultura")

	mockRepo.AssertCalled(t, "CreatePerformancePlanting", entityPerformancePlanting)
	mockRepo.AssertExpectations(t)
}

func TestGetAllPerformancePlanting_Success(t *testing.T) {

	mockRepo, service, _, dbResultPlanting, _ := SetupTestPerformancePlanting()

	var listDbResultPlanting []responses.DbResultPerformancePlanting
	listDbResultPlanting = append(listDbResultPlanting, dbResultPlanting)

	mockRepo.On("FindAll").Return(listDbResultPlanting, nil)

	result, err := service.GetAllPerformancePlanting()

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

	mockRepo.AssertCalled(t, "FindAll")
	mockRepo.AssertExpectations(t)
}

func TestGetAllPerformancePlanting_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPerformancePlanting()

	mockRepo.On("FindAll").Return([]responses.DbResultPerformancePlanting(nil), fmt.Errorf("erro ao buscar performance de plantação"))

	result, err := service.GetAllPerformancePlanting()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro ao buscar performance de plantação")

	mockRepo.AssertCalled(t, "FindAll")
	mockRepo.AssertExpectations(t)
}

func TestGetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID_Success(t *testing.T) {

	mockRepo, service, _, dbResultPlanting, _ := SetupTestPerformancePlanting()

	id := uint(1)

	mockRepo.On("FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID", id).Return(&dbResultPlanting, nil)

	response, err := service.GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(id)

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

	mockRepo.AssertCalled(t, "FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID", id)
	mockRepo.AssertExpectations(t)
}

func TestGetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPerformancePlanting()

	id := uint(1)

	mockRepo.On("FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID", id).Return(&responses.DbResultPerformancePlanting{}, errors.New("erro"))

	response, err := service.GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(id)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.ErrorContains(t, err, "erro")

	mockRepo.AssertCalled(t, "FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID", id)
	mockRepo.AssertExpectations(t)
}

func TestPutPerformancePlanting_Success(t *testing.T) {

	mockRepo, service, entityPerformancePlanting, _, requestPerformancePlanting := SetupTestPerformancePlanting()

	id := uint(1)

	mockRepo.On("UpdatePerformancePlanting", id, entityPerformancePlanting).Return(nil)

	err := service.PutPerformancePlanting(id, requestPerformancePlanting)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "UpdatePerformancePlanting", id, entityPerformancePlanting)
	mockRepo.AssertExpectations(t)
}

func TestPutPerformancePlanting_Error(t *testing.T) {

	mockRepo, service, entityPerformancePlanting, _, requestPerformancePlanting := SetupTestPerformancePlanting()

	id := uint(1)

	mockRepo.On("UpdatePerformancePlanting", id, entityPerformancePlanting).Return(errors.New("erro"))

	err := service.PutPerformancePlanting(id, requestPerformancePlanting)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "erro")

	mockRepo.AssertCalled(t, "UpdatePerformancePlanting", id, entityPerformancePlanting)
	mockRepo.AssertExpectations(t)

}

func TestDeletePerformancePlanting_Success(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPerformancePlanting()

	id := uint(1)

	mockRepo.On("DeletePerformancePlanting", id).Return(nil)

	err := service.DeletePerformancePlanting(id)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "DeletePerformancePlanting", id)
	mockRepo.AssertExpectations(t)

}

func TestDeletePerformancePlanting_Error(t *testing.T) {

	mockRepo, service, _, _, _ := SetupTestPerformancePlanting()

	id := uint(1)

	mockRepo.On("DeletePerformancePlanting", id).Return(errors.New("error"))

	err := service.DeletePerformancePlanting(id)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "error")

	mockRepo.AssertCalled(t, "DeletePerformancePlanting", id)
	mockRepo.AssertExpectations(t)

}
