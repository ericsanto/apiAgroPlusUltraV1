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

	responseProfit := responses.ProfitResponse{
		ValueSalePlantation: 5000000,
		TotalCost:           4000,
		Profit:              50,
		ProfitMargin:        30,
	}

	mockeRepo.On("FindProfit", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(&responseProfit, nil)

	response, err := service.GetProfit(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)

	assert.Nil(t, err)
	assert.Equal(t, responseProfit.Profit, response.Profit)
	assert.Equal(t, responseProfit.ValueSalePlantation, response.ValueSalePlantation)
	assert.Equal(t, responseProfit.TotalCost, response.TotalCost)
	assert.Equal(t, responseProfit.ProfitMargin, response.ProfitMargin)

	mockeRepo.AssertCalled(t, "FindProfit", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)
	mockeRepo.AssertExpectations(t)

}

func TestGetProfit_Error(t *testing.T) {

	mockeRepo := new(mocks.ProfitRepositoryMock)

	service := ProfitService{profitRepository: mockeRepo}

	mockeRepo.On("FindProfit", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID).Return(&responses.ProfitResponse{}, errors.New("erro"))

	response, err := service.GetProfit(batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.ErrorContains(t, err, "erro")

	mockeRepo.AssertCalled(t, "FindProfit", batchMOCKID, farmMOCKID, userMOCKID, plantingMOCKID)
	mockeRepo.AssertExpectations(t)

}
