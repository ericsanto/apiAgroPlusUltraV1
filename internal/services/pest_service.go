package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type PestServiceInterface interface {
	GetAllPest() ([]responses.PestResponse, error)
	GetFindByIdPest(id uint) (responses.PestResponse, error)
	PostPest(requestPest requests.PestRequest) error
	PutPest(id uint, requestPest requests.PestRequest) error
	DeletePest(id uint) error
}

type PestService struct {
	pestRepository *repositories.PestRepository
}

func NewPestService(pestRepository *repositories.PestRepository) *PestService {
	return &PestService{pestRepository: pestRepository}
}

func (p *PestService) GetAllPest() ([]responses.PestResponse, error) {

	var responsesPests []responses.PestResponse

	result, err := p.pestRepository.FindAllPest()
	if err != nil {
		return responsesPests, fmt.Errorf("Erro: %w", err)
	}

	for _, v := range result {
		responsePest := responses.PestResponse{
			Id:         v.Id,
			Name:       v.Name,
			TypePestId: v.TypePestId,
		}

		responsesPests = append(responsesPests, responsePest)
	}

	return responsesPests, nil
}

func (p *PestService) GetFindByIdPest(id uint) (responses.PestResponse, error) {

	var responsePest responses.PestResponse

	result, err := p.pestRepository.FindByIdPest(id)
	if err != nil {
		return responsePest, fmt.Errorf("Erro: %w", err)
	}

	responsePest = responses.PestResponse{
		Id:         result.Id,
		Name:       result.Name,
		TypePestId: result.TypePestId,
	}

	return responsePest, nil
}

func (p *PestService) PostPest(requestPest requests.PestRequest) error {
	entityPest := entities.PestEntity{
		Name:       requestPest.Name,
		TypePestId: requestPest.TypePestId,
	}

	if err := p.pestRepository.CreatePest(entityPest); err != nil {
		return fmt.Errorf("Erro: %w", err)
	}

	return nil
}

func (p *PestService) PutPest(id uint, pestRequest requests.PestRequest) error {

	pestEntity := entities.PestEntity{
		Name:       pestRequest.Name,
		TypePestId: pestRequest.TypePestId,
	}

	if err := p.pestRepository.UpdatePest(id, pestEntity); err != nil {
		return fmt.Errorf("Erro: %w", err)
	}

	return nil
}

func (p *PestService) DeletePest(id uint) error {

	if err := p.pestRepository.DeletePest(id); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
