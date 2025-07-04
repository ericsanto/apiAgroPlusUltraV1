package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
)

type PlantingRepositoryMock struct {
	mock.Mock
}

func (prm *PlantingRepositoryMock) FindByParamPlanting(batchID uint) (entities.PlantingEntity, error) {

	args := prm.Called(batchID)

	return args.Get(0).(entities.PlantingEntity), args.Error(1)
}

func (prm *PlantingRepositoryMock) CreatePlanting(entityPlanting entities.PlantingEntity) error {

	args := prm.Called(entityPlanting)

	return args.Error(0)
}

func (prm *PlantingRepositoryMock) FindByParamBatchNameOrIsActivePlanting(batchName string, active bool) ([]responses.BatchPlantiesResponse, error) {

	args := prm.Called(batchName, active)

	return args.Get(0).([]responses.BatchPlantiesResponse), args.Error(1)
}

func (prm *PlantingRepositoryMock) FindPlantingByID(id uint) (entities.PlantingEntity, error) {

	args := prm.Called(id)

	return args.Get(0).(entities.PlantingEntity), args.Error(1)
}

func (prm *PlantingRepositoryMock) UpdatePlanting(id uint, entityPlanting entities.PlantingEntity) error {

	args := prm.Called(id, entityPlanting)

	return args.Error(0)
}

func (prm *PlantingRepositoryMock) DeletePlanting(id uint) error {

	args := prm.Called(id)

	return args.Error(0)
}

func (prm *PlantingRepositoryMock) FindAllPlanting() ([]entities.PlantingEntity, error) {

	args := prm.Called()

	return args.Get(0).([]entities.PlantingEntity), args.Error(1)
}
