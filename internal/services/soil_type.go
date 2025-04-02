package services

import (
	"fmt"
	"log"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type SoilTypeServiceInterface interface {
	GetAllSoilType() ([]responses.SoilTypeResponse, error)
	GetFindByIdSoilType(id uint) (responses.SoilTypeResponse, error)
	PostSoilType(soilTypeRequest requests.SoilTypeRequest) error
	PutSoilType(soilTypeRequest requests.SoilTypeRequest) error
}

type SoilTypeService struct {
	soilTypeRepository *repositories.SoilTypeRepository
}

func NewSoilTypeService(soilTypeRepository *repositories.SoilTypeRepository) *SoilTypeService {

	return &SoilTypeService{soilTypeRepository: soilTypeRepository}
}

func (s *SoilTypeService) GetAllSoilType() ([]responses.SoilTypeResponse, error) {

	soilTypes, err := s.soilTypeRepository.FindAllSoilType()
	if err != nil {
		log.Printf("erro ao buscar todos os tipos de solo: %v", err)
		return nil, fmt.Errorf("não foi possível buscar todos os tipos de solo: %w", err)
	}

	var soilTypesResponses []responses.SoilTypeResponse

	for _, v := range soilTypes {
		soilTypesResponse := responses.SoilTypeResponse{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
		}

		soilTypesResponses = append(soilTypesResponses, soilTypesResponse)

	}

	return soilTypesResponses, nil

}

func (s *SoilTypeService) GetSoilTypeFindById(id uint) (responses.SoilTypeResponse, error) {
	var soilTypeResponse responses.SoilTypeResponse

	soilTypes, err := s.soilTypeRepository.FindByIdSoilType(id)
	if err != nil {
		log.Printf("erro ao buscar tipo de solo: %v", err)
		return soilTypeResponse, fmt.Errorf("não foi possível buscar todos os tipos de solo: %w", err)
	}

	soilTypeResponse = responses.SoilTypeResponse{
		Id:          soilTypes.Id,
		Name:        soilTypes.Name,
		Description: soilTypes.Description,
	}

	return soilTypeResponse, nil
}

func (s *SoilTypeService) CreateSoilType(requestSoilType requests.SoilTypeRequest) error {

	soilTypeModel := entities.SoilTypeEntity{
		Name:        requestSoilType.Name,
		Description: requestSoilType.Description,
	}

	if err := s.soilTypeRepository.CreateSoilType(&soilTypeModel); err != nil {
		return fmt.Errorf("erro ao criar tipo de solo: %w", err)
	}

	return nil
}

func (s *SoilTypeService) PutSoilType(id uint, requestSoilType requests.SoilTypeRequest) error {

	soilTypeModel := entities.SoilTypeEntity{
		Name:        requestSoilType.Name,
		Description: requestSoilType.Description,
	}

	if err := s.soilTypeRepository.UpdateSoilType(id, soilTypeModel); err != nil {
		return fmt.Errorf("não foi possivel atualizar. Id não existe")
	}

	return nil
}

func (s *SoilTypeService) DeleteTypeSoil(id uint) error {

	if err := s.soilTypeRepository.DeleteSoilType(id); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}
