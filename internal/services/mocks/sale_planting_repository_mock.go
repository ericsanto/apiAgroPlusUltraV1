package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type SalePlantingRepositoryMock struct {
	mock.Mock
}

func (sprm *SalePlantingRepositoryMock) CreateSalePlantingRepository(entitySalePlanting entities.SalePlantingEntity) error {

	args := sprm.Called(entitySalePlanting)

	return args.Error(0)
}

func (sprm *SalePlantingRepositoryMock) FindAllSalePlanting() ([]entities.SalePlantingEntity, error) {

	args := sprm.Called()

	return args.Get(0).([]entities.SalePlantingEntity), args.Error(1)
}

func (sprm *SalePlantingRepositoryMock) FindSalePlantingByID(id uint) (*entities.SalePlantingEntity, error) {

	args := sprm.Called(id)

	return args.Get(0).(*entities.SalePlantingEntity), args.Error(1)
}

func (sprm *SalePlantingRepositoryMock) UpdateSalePlanting(id uint, entitySalePlanting entities.SalePlantingEntity) error {

	args := sprm.Called(id, entitySalePlanting)

	return args.Error(0)
}

func (sprm *SalePlantingRepositoryMock) DeleteSalePlanting(id uint) error {

	args := sprm.Called(id)

	return args.Error(0)
}
