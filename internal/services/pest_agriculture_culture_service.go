package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type PestAgricultureCultureServiceInterface interface {
	GetAllPestAgricultureCulture() ([]responses.PestAgricultureCultureResponse, error)
	GetFindByIdPestAgricultureCulture(pestId, cultureId uint) (responses.PestAgricultureCultureResponse, error)
	PostPestAgricultureCulture(requestPestAgricultureCulture requests.PestAgricultureCultureRequest) error
	PutPestAgricultureCulture(pestId, cultureId uint, requestAgricultureCulture requests.PestAgricultureCultureRequest) error
	DeletePestAgricultureCulture(pestId, cultureId uint)
}

type PestAgricultureCultureService struct {
	pestAgricultureCultureRepository *repositories.PestAgricultureCultureRepository
}

func NewPestAgricultureCultureService(pestAgricultureCultureRepository *repositories.PestAgricultureCultureRepository) *PestAgricultureCultureService {
	return &PestAgricultureCultureService{pestAgricultureCultureRepository: pestAgricultureCultureRepository}
}

func (p *PestAgricultureCultureService) GetAllPestAgricultureCulture() ([]responses.PestAgricultureCultureResponse, error) {

	result, err := p.pestAgricultureCultureRepository.FindAllPestAgricultureCulture()
	if err != nil {
		return result, fmt.Errorf("erro: %v", err)
	}

	return result, nil
}

func (p *PestAgricultureCultureService) GetFindByIdPestAgricultureCulture(pestId, cultureId uint) (responses.PestAgricultureCultureResponse, error) {

	result, err := p.pestAgricultureCultureRepository.FindByIdPestAgricultureCulture(pestId, cultureId)
	if err != nil {
		return result, fmt.Errorf("erro: %v", err)
	}

	return result, nil
}

func (p *PestAgricultureCultureService) PostPestAgricultureCulture(requestPestAgricultureCulture requests.PestAgricultureCultureRequest) error {

	entityPestAgricultureCulture := entities.PestAgricultureCulture{
		PestId:               requestPestAgricultureCulture.PestId,
		AgricultureCultureId: requestPestAgricultureCulture.AgricultureCultureId,
		Description:          requestPestAgricultureCulture.Description,
		ImageUrl:             requestPestAgricultureCulture.ImageUrl,
	}

	result := p.pestAgricultureCultureRepository.CreatePestAgricultureCulture(entityPestAgricultureCulture)
	if result != nil {
		return fmt.Errorf("erro no repositório: %v", result)
	}

	return nil
}

func (p *PestAgricultureCultureService) PutPestAgricultureCulture(pestId, cultureId uint, requestAgricultureCulture requests.PestAgricultureCultureRequest) error {

	entityPestAgricultureCulture := entities.PestAgricultureCulture{
		PestId:               requestAgricultureCulture.PestId,
		AgricultureCultureId: requestAgricultureCulture.AgricultureCultureId,
		Description:          requestAgricultureCulture.Description,
		ImageUrl:             requestAgricultureCulture.ImageUrl,
	}

	if err := p.pestAgricultureCultureRepository.UpdatePestAgricultureCulture(pestId, cultureId, entityPestAgricultureCulture); err != nil {
		return fmt.Errorf("erro no repositório: %v", err)
	}

	return nil
}

func (p *PestAgricultureCultureService) DeletePestAgricultureCulture(pestId, cultureId uint) error {

	if err := p.pestAgricultureCultureRepository.DeletePestAgricultureCulture(pestId, cultureId); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}
