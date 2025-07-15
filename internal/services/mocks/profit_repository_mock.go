package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
)

type ProfitRepositoryMock struct {
	mock.Mock
}

func (prm *ProfitRepositoryMock) FindProfit(batchID, farmID, userID, plantingID uint) (*responses.ProfitResponse, error) {

	args := prm.Called(batchID, farmID, userID, plantingID)

	return args.Get(0).(*responses.ProfitResponse), args.Error(1)
}
