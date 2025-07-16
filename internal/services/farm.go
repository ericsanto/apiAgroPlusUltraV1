package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type FarmServiceInterface interface {
	Create(farmRequest requests.FarmRequest) error
	GetFarmByID(userID, id uint) (*responses.FarmResponse, error)
	GetAllFarm(userID uint) ([]responses.FarmResponse, error)
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

func (fs *FarmService) GetFarmByID(userID, id uint) (*responses.FarmResponse, error) {

	responseFarm, err := fs.farmRepository.FindByID(userID, id)

	if err != nil {
		return nil, fmt.Errorf("erro %w", err)
	}

	return responseFarm, nil
}

func (fs *FarmService) GetAllFarm(userID uint) ([]responses.FarmResponse, error) {

	listFarmResponse, err := fs.farmRepository.FindAll(userID)

	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	return listFarmResponse, nil

}
