package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
)

type FarmRepositoryMock struct {
	mock.Mock
}

func (frm *FarmRepositoryMock) FindByID(userID, id uint) (*responses.FarmResponse, error) {

	args := frm.Called(userID, id)

	return args.Get(0).(*responses.FarmResponse), args.Error(1)
}

func (frm *FarmRepositoryMock) FindAll(userID uint) ([]responses.FarmResponse, error) {

	args := frm.Called(userID)

	return args.Get(0).([]responses.FarmResponse), args.Error(1)
}

func (frm *FarmRepositoryMock) Create(ctx context.Context, farmEntity entities.FarmEntity) error {

	args := frm.Called(ctx, farmEntity)

	return args.Error(0)
}
