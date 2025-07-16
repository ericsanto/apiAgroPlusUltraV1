package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
)

type PlantingRepositoryMock struct {
	mock.Mock
}

func (prm *PlantingRepositoryMock) FindByParamPlanting(userID, farmID, batchID uint) (entities.PlantingEntity, error) {

	args := prm.Called(userID, farmID, batchID)

	return args.Get(0).(entities.PlantingEntity), args.Error(1)
}

func (prm *PlantingRepositoryMock) CreatePlanting(entityPlanting entities.PlantingEntity) error {

	args := prm.Called(entityPlanting)

	return args.Error(0)
}

func (prm *PlantingRepositoryMock) FindByParamBatchNameOrIsActivePlanting(batchName string, active bool, userID, farmID uint) ([]responses.BatchPlantiesResponse, error) {

	args := prm.Called(batchName, active, userID, farmID)

	return args.Get(0).([]responses.BatchPlantiesResponse), args.Error(1)
}

func (prm *PlantingRepositoryMock) FindPlantingByID(batchID, farmID, userID, plantingID uint) (entities.PlantingEntity, error) {

	args := prm.Called(batchID, farmID, userID, plantingID)

	return args.Get(0).(entities.PlantingEntity), args.Error(1)
}

func (prm *PlantingRepositoryMock) UpdatePlanting(batchID, farmID, userID, plantingID uint, entityPlanting entities.PlantingEntity) error {

	args := prm.Called(batchID, farmID, userID, plantingID, entityPlanting)

	return args.Error(0)
}

func (prm *PlantingRepositoryMock) DeletePlanting(batchID, farmID, userID, plantingID uint) error {

	args := prm.Called(batchID, farmID, userID, plantingID)

	return args.Error(0)
}

func (prm *PlantingRepositoryMock) FindAllPlanting(userID, farmID, batchID uint) ([]entities.PlantingEntity, error) {

	args := prm.Called(userID, farmID, batchID)

	return args.Get(0).([]entities.PlantingEntity), args.Error(1)
}
