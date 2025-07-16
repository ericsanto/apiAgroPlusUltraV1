package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type SalePlantingRepositoryMock struct {
	mock.Mock
}

func (sprm *SalePlantingRepositoryMock) CreateSalePlantingRepository(batchID, farmID, userID, plantingID uint, entitySalePlanting entities.SalePlantingEntity) error {

	args := sprm.Called(batchID, farmID, userID, plantingID, entitySalePlanting)

	return args.Error(0)
}

func (sprm *SalePlantingRepositoryMock) FindAllSalePlanting(batchID, farmID, userID uint) ([]entities.SalePlantingEntity, error) {

	args := sprm.Called(batchID, farmID, userID)

	return args.Get(0).([]entities.SalePlantingEntity), args.Error(1)
}

func (sprm *SalePlantingRepositoryMock) FindSalePlantingByID(batchID, farmID, userID, salePlantingID uint) (*entities.SalePlantingEntity, error) {

	args := sprm.Called(batchID, farmID, userID, salePlantingID)

	return args.Get(0).(*entities.SalePlantingEntity), args.Error(1)
}

func (sprm *SalePlantingRepositoryMock) UpdateSalePlanting(batchID, farmID, userID, salePlantingID uint, entitySalePlanting entities.SalePlantingEntity) error {

	args := sprm.Called(batchID, farmID, userID, salePlantingID, entitySalePlanting)

	return args.Error(0)
}

func (sprm *SalePlantingRepositoryMock) DeleteSalePlanting(batchID, farmID, userID, salePlantingID uint) error {

	args := sprm.Called(batchID, farmID, userID, salePlantingID)

	return args.Error(0)
}
