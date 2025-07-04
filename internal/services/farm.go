package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type FarmServiceInterface interface {
	Create(farmRequest requests.FarmRequest) error
}

type FarmService struct {
	farmRepository repositories.FarmRepositoryInterface
}

func NewFarmService(farmRepository repositories.FarmRepositoryInterface) FarmServiceInterface {
	return &FarmService{farmRepository: farmRepository}
}

func (fs *FarmService) Create(farmRequest requests.FarmRequest) error {

	farmEntity := entities.FarmEntity{
		Name:   farmRequest.Name,
		UserID: farmRequest.UserID,
	}

	if err := fs.farmRepository.Create(farmEntity); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil

}
