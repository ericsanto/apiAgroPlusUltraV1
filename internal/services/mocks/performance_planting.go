package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
)

type PerformancePlantingRepository struct {
	mock.Mock
}

func (pcr *PerformancePlantingRepository) CreatePerformancePlanting(entityPerformanceCulutre entities.PerformancePlantingEntity) error {

	args := pcr.Called(entityPerformanceCulutre)

	return args.Error(0)
}

func (pcr *PerformancePlantingRepository) FindAll() ([]responses.DbResultPerformancePlanting, error) {

	args := pcr.Called()

	return args.Get(0).([]responses.DbResultPerformancePlanting), args.Error(1)
}

func (pcr *PerformancePlantingRepository) FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID(id uint) (*responses.DbResultPerformancePlanting, error) {

	args := pcr.Called(id)

	return args.Get(0).(*responses.DbResultPerformancePlanting), args.Error(1)
}

func (pcr *PerformancePlantingRepository) FindPerformancePlantingByID(id uint) (*entities.PerformancePlantingEntity, error) {

	args := pcr.Called(id)

	return args.Get(0).(*entities.PerformancePlantingEntity), args.Error(1)
}

func (pcr *PerformancePlantingRepository) UpdatePerformancePlanting(id uint, entityPerformancePlanting entities.PerformancePlantingEntity) error {

	args := pcr.Called(id, entityPerformancePlanting)

	return args.Error(0)
}

func (pcr *PerformancePlantingRepository) DeletePerformancePlanting(id uint) error {

	args := pcr.Called(id)

	return args.Error(0)
}
