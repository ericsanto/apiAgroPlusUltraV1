package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type ProfitService struct {
	profitRepository *repositories.ProfitRepository
}

func NewProfitService(profitRepository *repositories.ProfitRepository) *ProfitService {
	return &ProfitService{profitRepository: profitRepository}
}

func (p *ProfitService) GetProfit(plantingID, userID uint) (*responses.ProfitResponse, error) {

	profitResponse, err := p.profitRepository.FindProfit(plantingID, userID)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	return profitResponse, nil
}
