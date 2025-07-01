package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type SoilTypeRepositoryMock struct {
	mock.Mock
}

func (strm *SoilTypeRepositoryMock) CreateSoilType(soilTypeModel entities.SoilTypeEntity) error {

	args := strm.Called(soilTypeModel)

	return args.Error(0)
}

func (strm *SoilTypeRepositoryMock) FindAllSoilType() ([]entities.SoilTypeEntity, error) {

	args := strm.Called()

	return args.Get(0).([]entities.SoilTypeEntity), args.Error(1)
}

func (strm *SoilTypeRepositoryMock) FindByIdSoilType(id uint) (*entities.SoilTypeEntity, error) {
	args := strm.Called(id)

	result := args.Get(0)

	if result == nil {
		return nil, args.Error(1)
	}

	return result.(*entities.SoilTypeEntity), args.Error(1)

}

func (strm *SoilTypeRepositoryMock) UpdateSoilType(id uint, soilTypeModel entities.SoilTypeEntity) error {

	args := strm.Called(id, soilTypeModel)

	return args.Error(0)
}

func (strm *SoilTypeRepositoryMock) DeleteSoilType(id uint) error {

	return nil
}
