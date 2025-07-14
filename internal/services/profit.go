package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type ProfitServiceInterface interface {
	GetProfit(batchID, farmID, userID, plantingID uint) (*responses.ProfitResponse, error)
}

type ProfitService struct {
	profitRepository repositories.ProfitRepositoryInterface
}

func NewProfitService(profitRepository repositories.ProfitRepositoryInterface) ProfitServiceInterface {
	return &ProfitService{profitRepository: profitRepository}
}

func (p *ProfitService) GetProfit(batchID, farmID, userID, plantingID uint) (*responses.ProfitResponse, error) {

	profitResponse, err := p.profitRepository.FindProfit(batchID, farmID, userID, plantingID)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	return profitResponse, nil
}
