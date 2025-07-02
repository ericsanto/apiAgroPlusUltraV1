package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
)

func TestPostSoilType_Success(t *testing.T) {

	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockRepo}

	request := requests.SoilTypeRequest{
		Name:        "arenoso",
		Description: "teste",
	}

	entity := entities.SoilTypeEntity{
		Id:          0,
		Name:        request.Name,
		Description: request.Description,
	}

	mockRepo.On("CreateSoilType", entity).Return(nil)

	err := service.PostSoilType(request)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPostSoilType_Error(t *testing.T) {

	//cria um ponteiro para mocks.SoilTypeRepositoryMock
	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockRepo}

	request := requests.SoilTypeRequest{
		Name:        "teste",
		Description: "teste",
	}

	mockRepo.On("CreateSoilType", mock.AnythingOfType("entities.SoilTypeEntity")).
		Return(fmt.Errorf("erro ao criar tipo de solo"))

	err := service.PostSoilType(request)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "erro ao criar tipo de solo")

	mockRepo.AssertExpectations(t)
}

func TestFindByIdSoilType_Success(t *testing.T) {

	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockRepo}

	id := uint(1)

	entitySoilType := entities.SoilTypeEntity{
		Id:          id,
		Name:        "teste",
		Description: "teste",
	}

	mockRepo.On("FindByIdSoilType", id).Return(&entitySoilType, nil)

	responseSoilType, err := service.GetSoilTypeFindById(id)

	assert.Nil(t, err)
	assert.NotNil(t, responseSoilType)
	assert.Equal(t, id, responseSoilType.Id)
	assert.Equal(t, entitySoilType.Id, responseSoilType.Id)
	assert.Equal(t, entitySoilType.Description, responseSoilType.Description)
	assert.Equal(t, entitySoilType.Name, responseSoilType.Name)

	mockRepo.AssertExpectations(t)
}

func TestFindByIdSoilType_Error(t *testing.T) {

	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockRepo}

	id := uint(1)

	mockRepo.On("FindByIdSoilType", id).Return(&entities.SoilTypeEntity{}, fmt.Errorf("erro ao buscar tipo de solo"))

	responseSoilType, err := service.GetSoilTypeFindById(id)

	assert.Contains(t, err.Error(), "erro ao buscar tipo de solo")
	assert.EqualValues(t, "", responseSoilType.Name)
	assert.EqualValues(t, 0, responseSoilType.Id)
	assert.EqualValues(t, "", responseSoilType.Description)

}

func TestUpdateSoilType_Success(t *testing.T) {

	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockRepo}

	id := uint(0)

	requestSoilType := requests.SoilTypeRequest{
		Name:        "teste",
		Description: "tesete",
	}

	expectedEntity := entities.SoilTypeEntity{
		Id:          0,
		Name:        requestSoilType.Name,
		Description: requestSoilType.Description,
	}

	mockRepo.On("UpdateSoilType", id, expectedEntity).Return(nil)

	err := service.PutSoilType(id, requestSoilType)

	assert.Nil(t, err)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)

}

func TestUpdateSoilType_Error(t *testing.T) {

	mockeRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockeRepo}

	id := uint(0)

	requestSoilType := requests.SoilTypeRequest{
		Name:        "TESTE",
		Description: "teste",
	}

	mockeRepo.On("UpdateSoilType", id, mock.AnythingOfType("entities.SoilTypeEntity")).Return(fmt.Errorf("não foi possivel atualizar. Id não existe"))

	err := service.PutSoilType(id, requestSoilType)

	assert.Contains(t, err.Error(), "não foi possivel atualizar. Id não existe")
	assert.Error(t, err)
}

func TestFindAllSoilType_Success(t *testing.T) {

	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockRepo}

	soilType1 := entities.SoilTypeEntity{
		Id:          uint(1),
		Name:        "teste1",
		Description: "teste1",
	}

	soilType2 := entities.SoilTypeEntity{
		Id:          uint(2),
		Name:        "teste2",
		Description: "teste2",
	}

	var soilTypes []entities.SoilTypeEntity

	soilTypes = append(soilTypes, soilType1)
	soilTypes = append(soilTypes, soilType2)

	mockRepo.On("FindAllSoilType").Return(soilTypes, nil)

	soilTypesResponse, err := service.GetAllSoilType()

	assert.Nil(t, err)
	assert.Equal(t, soilType1.Id, soilTypesResponse[0].Id)
	assert.Equal(t, soilType2.Id, soilTypesResponse[1].Id)
	assert.NotNil(t, soilTypesResponse)
	assert.EqualValues(t, soilType1.Name, soilTypesResponse[0].Name)
	assert.EqualValues(t, soilType2.Name, soilTypesResponse[1].Name)
	assert.EqualValues(t, soilType1.Description, soilTypesResponse[0].Description)
	assert.EqualValues(t, soilType2.Description, soilTypesResponse[1].Description)

	mockRepo.AssertExpectations(t)
}

func TestFindAllSoilType(t *testing.T) {

	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockRepo}

	mockRepo.On("FindAllSoilType").Return([]entities.SoilTypeEntity{}, fmt.Errorf("não foi possível buscar todos os tipos de solo"))

	_, err := service.GetAllSoilType()

	assert.Contains(t, err.Error(), "não foi possível buscar todos os tipos de solo")
	assert.Error(t, err)
}

func TestDeleteSoilTyp_Success(t *testing.T) {

	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{mockRepo}

	id := uint(0)

	mockRepo.On("DeleteSoilType", id).Return(nil)

	err := service.DeleteTypeSoil(id)

	assert.NoError(t, err)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)

}

func TestDeleteSoilTyp_Error(t *testing.T) {

	mockRepo := new(mocks.SoilTypeRepositoryMock)

	service := SoilTypeService{soilTypeRepository: mockRepo}

	id := uint(1)

	mockRepo.On("DeleteSoilType", id).Return(fmt.Errorf("não foi possível deletar tipo de solo"))

	err := service.DeleteTypeSoil(id)

	assert.Contains(t, err.Error(), "não foi possível deletar tipo de solo")
	assert.NotNil(t, err)

	mockRepo.AssertExpectations(t)
}
