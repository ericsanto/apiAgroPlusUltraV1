package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
)

type PerformancePlantingRepository struct {
	mock.Mock
}

func (pcr *PerformancePlantingRepository) CreatePerformancePlanting(batchID, farmID, userID, plantingID uint, entityPerformanceCulutre entities.PerformancePlantingEntity) error {

	args := pcr.Called(batchID, farmID, userID, plantingID, entityPerformanceCulutre)

	return args.Error(0)
}

func (pcr *PerformancePlantingRepository) FindAll(batchID, farmID, userID uint) ([]responses.DbResultPerformancePlanting, error) {

	args := pcr.Called(batchID, farmID, userID)

	return args.Get(0).([]responses.DbResultPerformancePlanting), args.Error(1)
}

func (pcr *PerformancePlantingRepository) FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID(batchID, farmID, userID, plantingID, performanceID uint) (*responses.DbResultPerformancePlanting, error) {

	args := pcr.Called(batchID, farmID, userID, plantingID, performanceID)

	return args.Get(0).(*responses.DbResultPerformancePlanting), args.Error(1)
}

func (pcr *PerformancePlantingRepository) FindPerformancePlantingByID(batchID, farmID, userID, plantingID, performanceID uint) (*entities.PerformancePlantingEntity, error) {

	args := pcr.Called(batchID, farmID, userID, plantingID, performanceID)

	return args.Get(0).(*entities.PerformancePlantingEntity), args.Error(1)
}

func (pcr *PerformancePlantingRepository) UpdatePerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint, entityPerformancePlanting entities.PerformancePlantingEntity) error {

	args := pcr.Called(batchID, farmID, userID, plantingID, performanceID, entityPerformancePlanting)

	return args.Error(0)
}

func (pcr *PerformancePlantingRepository) DeletePerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint) error {

	args := pcr.Called(batchID, farmID, userID, plantingID, performanceID)

	return args.Error(0)
}
