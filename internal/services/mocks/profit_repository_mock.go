package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
)

type ProfitRepositoryMock struct {
	mock.Mock
}

func (prm *ProfitRepositoryMock) FindProfit(plantingID, userID uint) (*responses.ProfitResponse, error) {

	args := prm.Called(plantingID, userID)

	return args.Get(0).(*responses.ProfitResponse), args.Error(1)
}
