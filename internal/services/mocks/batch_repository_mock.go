package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type BatchRepositoryMock struct {
	mock.Mock
}

func (brm *BatchRepositoryMock) Create(userID, farmID uint, batchEntity entities.BatchEntity) error {

	args := brm.Called(userID, farmID, batchEntity)

	return args.Error(0)
}

func (brm *BatchRepositoryMock) FindAllBatch(userID, farmID uint) ([]entities.BatchEntity, error) {

	args := brm.Called(userID, farmID)

	return args.Get(0).([]entities.BatchEntity), args.Error(1)
}

func (brm *BatchRepositoryMock) BatchFindById(userID, farmID, batchID uint) (*entities.BatchEntity, error) {

	args := brm.Called(userID, farmID, batchID)

	return args.Get(0).(*entities.BatchEntity), args.Error(1)
}

func (brm *BatchRepositoryMock) Update(userID, farmID, batchID uint, entityBatch entities.BatchEntity) error {

	args := brm.Called(userID, farmID, batchID, entityBatch)

	return args.Error(0)
}

func (brm *BatchRepositoryMock) DeleteBatch(userID, farmID, batchID uint) error {

	args := brm.Called(userID, farmID, batchID)

	return args.Error(0)
}
