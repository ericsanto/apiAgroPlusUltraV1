package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services/mocks"
)

func TestGetProfit_Success(t *testing.T) {

	mockeRepo := new(mocks.ProfitRepositoryMock)

	service := ProfitService{profitRepository: mockeRepo}

	plantingID := uint(1)
	userID := uint(1)

	responseProfit := responses.ProfitResponse{
		ValueSalePlantation: 5000000,
		TotalCost:           4000,
		Profit:              50,
		ProfitMargin:        30,
	}

	mockeRepo.On("FindProfit", plantingID, userID).Return(&responseProfit, nil)

	response, err := service.GetProfit(plantingID, userID)

	assert.Nil(t, err)
	assert.Equal(t, responseProfit.Profit, response.Profit)
	assert.Equal(t, responseProfit.ValueSalePlantation, response.ValueSalePlantation)
	assert.Equal(t, responseProfit.TotalCost, response.TotalCost)
	assert.Equal(t, responseProfit.ProfitMargin, response.ProfitMargin)

	mockeRepo.AssertCalled(t, "FindProfit", plantingID, userID)
	mockeRepo.AssertExpectations(t)

}

func TestGetProfit_Error(t *testing.T) {

	mockeRepo := new(mocks.ProfitRepositoryMock)

	service := ProfitService{profitRepository: mockeRepo}

	plantingID := uint(1)
	userID := uint(1)

	mockeRepo.On("FindProfit", plantingID, userID).Return(&responses.ProfitResponse{}, errors.New("erro"))

	response, err := service.GetProfit(plantingID, userID)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.ErrorContains(t, err, "erro")

	mockeRepo.AssertCalled(t, "FindProfit", plantingID, userID)
	mockeRepo.AssertExpectations(t)

}
