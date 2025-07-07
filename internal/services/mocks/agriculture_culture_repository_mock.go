package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type AgricultureCultureRepositoryMock struct {
	mock.Mock
}

func (acrm *AgricultureCultureRepositoryMock) FindAllAgricultureCulture() ([]entities.AgricultureCultureEntity, error) {
	args := acrm.Called()

	return args.Get(0).([]entities.AgricultureCultureEntity), args.Error(1)
}

func (acrm *AgricultureCultureRepositoryMock) FindByIdAgricultureCulture(id uint) (*entities.AgricultureCultureEntity, error) {

	args := acrm.Called(id)

	return args.Get(0).(*entities.AgricultureCultureEntity), args.Error(1)
}

func (acrm *AgricultureCultureRepositoryMock) CreateAgricultureCulture(agriculutreCulture entities.AgricultureCultureEntity) error {

	args := acrm.Called(agriculutreCulture)

	return args.Error(0)
}

func (acrm *AgricultureCultureRepositoryMock) UpdateAgricultureCulture(id uint, agricultureCultureEntity entities.AgricultureCultureEntity) error {

	args := acrm.Called(id, agricultureCultureEntity)

	return args.Error(0)
}

func (acrm *AgricultureCultureRepositoryMock) DeleteAgricultureCulture(id uint) error {

	args := acrm.Called(id)

	return args.Error(0)
}
