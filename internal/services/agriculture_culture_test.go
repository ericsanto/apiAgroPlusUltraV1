package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
)

func SetupTestAgricultureCulture() (*mocks.AgricultureCultureRepositoryMock, AgricultureCultureService,
	entities.AgricultureCultureEntity, requests.AgricultureCultureRequest) {

	mockRepo := new(mocks.AgricultureCultureRepositoryMock)

	service := AgricultureCultureService{agricultureRepository: mockRepo}

	requestAgricultureCulture := requests.AgricultureCultureRequest{
		Name:                      "milho",
		Variety:                   "teste",
		UseType:                   "GRAO",
		Region:                    "NORDESTE",
		SoilTypeId:                uint(1),
		PhIdealSoil:               5.4,
		MaxTemperature:            30,
		MinTemperature:            18,
		ExcellentTemperature:      25,
		WeeklyWaterRequirementMax: 8.8,
		WeeklyWaterRequirementMin: 5.3,
		SunlightRequirement:       6,
	}

	entityAgricultureCulture := entities.AgricultureCultureEntity{
		Name:                      "milho",
		Variety:                   "teste",
		UseType:                   "GRAO",
		Region:                    "NORDESTE",
		SoilTypeId:                uint(1),
		PhIdealSoil:               5.4,
		MaxTemperature:            30,
		MinTemperature:            18,
		ExcellentTemperature:      25,
		WeeklyWaterRequirementMax: 8.8,
		WeeklyWaterRequirementMin: 5.3,
		SunlightRequirement:       6,
	}

	return mockRepo, service, entityAgricultureCulture, requestAgricultureCulture
}

func TestFindAllAgricultureCulture_Success(t *testing.T) {

	mockRepo, service, entityAgricultureCulture, _ := SetupTestAgricultureCulture()

	var listAgricultureCultureEntity []entities.AgricultureCultureEntity
	listAgricultureCultureEntity = append(listAgricultureCultureEntity, entityAgricultureCulture)

	mockRepo.On("FindAllAgricultureCulture").Return(listAgricultureCultureEntity, nil)

	listResponseAgricultureCulture, err := service.FindAllAgricultureCultureService()

	assert.Nil(t, err)

	for i := range listAgricultureCultureEntity {
		assert.EqualValues(t, listAgricultureCultureEntity[i].Name, listResponseAgricultureCulture[i].Name)
		assert.EqualValues(t, listAgricultureCultureEntity[i].Variety, listResponseAgricultureCulture[i].Variety)
		assert.EqualValues(t, listAgricultureCultureEntity[i].UseType, listResponseAgricultureCulture[i].UseType)
		assert.EqualValues(t, listAgricultureCultureEntity[i].Region, listResponseAgricultureCulture[i].Region)
		assert.EqualValues(t, listAgricultureCultureEntity[i].SoilTypeId, listResponseAgricultureCulture[i].SoilTypeId)
		assert.EqualValues(t, listAgricultureCultureEntity[i].PhIdealSoil, listResponseAgricultureCulture[i].PhIdealSoil)
		assert.EqualValues(t, listAgricultureCultureEntity[i].MaxTemperature, listResponseAgricultureCulture[i].MaxTemperature)
		assert.EqualValues(t, listAgricultureCultureEntity[i].MinTemperature, listResponseAgricultureCulture[i].MinTemperature)
		assert.EqualValues(t, listAgricultureCultureEntity[i].ExcellentTemperature, listResponseAgricultureCulture[i].ExcellentTemperature)
		assert.EqualValues(t, listAgricultureCultureEntity[i].WeeklyWaterRequirementMax, listResponseAgricultureCulture[i].WeeklyWaterRequirememntMax)
		assert.EqualValues(t, listAgricultureCultureEntity[i].WeeklyWaterRequirementMin, listResponseAgricultureCulture[i].WeeklyWaterRequirememntMin)
		assert.EqualValues(t, listAgricultureCultureEntity[i].SunlightRequirement, listResponseAgricultureCulture[i].SunlightRequirement)
	}

	mockRepo.AssertCalled(t, "FindAllAgricultureCulture")
	mockRepo.AssertExpectations(t)
}

func TestFindAllAgricultureCulture_Error(t *testing.T) {

	mockRepo, service, _, _ := SetupTestAgricultureCulture()

	mockRepo.On("FindAllAgricultureCulture").Return([]entities.AgricultureCultureEntity{}, errors.New(""))

	response, err := service.FindAllAgricultureCultureService()

	assert.Nil(t, response)
	assert.NotNil(t, err)

	mockRepo.AssertCalled(t, "FindAllAgricultureCulture")
	mockRepo.AssertExpectations(t)
}

func TestFindByIdAgricultureCultureService_Success(t *testing.T) {

	mockRepo, service, entityAgricultureCulture, _ := SetupTestAgricultureCulture()

	id := uint(1)

	mockRepo.On("FindByIdAgricultureCulture", id).Return(&entityAgricultureCulture, nil)

	responseAgricultureCulture, err := service.FindByIdAgricultureCultureService(id)

	assert.Nil(t, err)
	assert.EqualValues(t, entityAgricultureCulture.Name, responseAgricultureCulture.Name)
	assert.EqualValues(t, entityAgricultureCulture.Variety, responseAgricultureCulture.Variety)
	assert.EqualValues(t, entityAgricultureCulture.Region, responseAgricultureCulture.Region)
	assert.EqualValues(t, entityAgricultureCulture.UseType, responseAgricultureCulture.UseType)
	assert.EqualValues(t, entityAgricultureCulture.SoilTypeId, responseAgricultureCulture.SoilTypeId)
	assert.EqualValues(t, entityAgricultureCulture.PhIdealSoil, responseAgricultureCulture.PhIdealSoil)
	assert.EqualValues(t, entityAgricultureCulture.MaxTemperature, responseAgricultureCulture.MaxTemperature)
	assert.EqualValues(t, entityAgricultureCulture.MinTemperature, responseAgricultureCulture.MinTemperature)
	assert.EqualValues(t, entityAgricultureCulture.ExcellentTemperature, responseAgricultureCulture.ExcellentTemperature)
	assert.EqualValues(t, entityAgricultureCulture.WeeklyWaterRequirementMax, responseAgricultureCulture.WeeklyWaterRequirememntMax)
	assert.EqualValues(t, entityAgricultureCulture.WeeklyWaterRequirementMin, responseAgricultureCulture.WeeklyWaterRequirememntMin)
	assert.EqualValues(t, entityAgricultureCulture.SunlightRequirement, responseAgricultureCulture.SunlightRequirement)

	mockRepo.AssertCalled(t, "FindByIdAgricultureCulture", id)
	mockRepo.AssertExpectations(t)
}

func TestFindByIdAgricultureCultureService_Error(t *testing.T) {

	mockRepo, service, _, _ := SetupTestAgricultureCulture()

	id := uint(1)

	mockRepo.On("FindByIdAgricultureCulture", id).Return(&entities.AgricultureCultureEntity{}, errors.New(""))

	responseAgricultureCulture, err := service.FindByIdAgricultureCultureService(id)

	assert.EqualValues(t, "", responseAgricultureCulture.Name)
	assert.NotNil(t, err)

	mockRepo.AssertCalled(t, "FindByIdAgricultureCulture", id)
	mockRepo.AssertExpectations(t)

}

func TestCreateAgricultureCultureService_Success(t *testing.T) {

	mockRepo, service, entityAgricultureCulture, requestAgricultureCulture := SetupTestAgricultureCulture()

	mockRepo.On("CreateAgricultureCulture", entityAgricultureCulture).Return(nil)

	err := service.CreateAgricultureCultureService(requestAgricultureCulture)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "CreateAgricultureCulture", entityAgricultureCulture)
	mockRepo.AssertExpectations(t)
}

func TestCreateAgricultureCultureService_Error(t *testing.T) {

	mockRepo, service, entityAgricultureCulture, requestAgricultureCulture := SetupTestAgricultureCulture()

	mockRepo.On("CreateAgricultureCulture", entityAgricultureCulture).Return(errors.New(""))

	err := service.CreateAgricultureCultureService(requestAgricultureCulture)

	assert.NotNil(t, err)
	assert.Error(t, err)

	mockRepo.AssertCalled(t, "CreateAgricultureCulture", entityAgricultureCulture)
	mockRepo.AssertExpectations(t)
}

func TestPutAgricultureCultureService_Success(t *testing.T) {

	mockRepo, service, entityAgricultureCulture, requestAgricultureCulture := SetupTestAgricultureCulture()

	id := uint(1)

	mockRepo.On("UpdateAgricultureCulture", id, entityAgricultureCulture).Return(nil)

	err := service.PutAgricultureCultureService(id, requestAgricultureCulture)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "UpdateAgricultureCulture", id, entityAgricultureCulture)
	mockRepo.AssertExpectations(t)
}

func TestPutAgricultureCultureService_Error(t *testing.T) {

	mockRepo, service, entityAgricultureCulture, requestAgricultureCulture := SetupTestAgricultureCulture()

	id := uint(1)

	mockRepo.On("UpdateAgricultureCulture", id, entityAgricultureCulture).Return(errors.New(""))

	err := service.PutAgricultureCultureService(id, requestAgricultureCulture)

	assert.NotNil(t, err)

	mockRepo.AssertCalled(t, "UpdateAgricultureCulture", id, entityAgricultureCulture)
	mockRepo.AssertExpectations(t)
}

func TestDeleteAgricultureCultureService_Success(t *testing.T) {

	mockRepo, service, _, _ := SetupTestAgricultureCulture()

	id := uint(1)

	mockRepo.On("DeleteAgricultureCulture", id).Return(nil)

	err := service.DeleteAgricultureCultureService(id)

	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "DeleteAgricultureCulture", id)
	mockRepo.AssertExpectations(t)
}

func TestDeleteAgricultureCultureService_Error(t *testing.T) {

	mockRepo, service, _, _ := SetupTestAgricultureCulture()

	id := uint(1)

	mockRepo.On("DeleteAgricultureCulture", id).Return(errors.New(""))

	err := service.DeleteAgricultureCultureService(id)

	assert.NotNil(t, err)

	mockRepo.AssertCalled(t, "DeleteAgricultureCulture", id)
	mockRepo.AssertExpectations(t)
}
