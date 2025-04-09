package services

import (
	"fmt"
	"strings"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type SustainablePestControlService struct {
	sustainablePestControlRepository *repositories.SustainablePestControlRepository
}

func NewSustainablePestControlService(sustainablePestControlRepository *repositories.SustainablePestControlRepository) *SustainablePestControlService {
	return &SustainablePestControlService{sustainablePestControlRepository: sustainablePestControlRepository}
}

func (s *SustainablePestControlService) GetAllSustainablePestControl() ([]responses.SustainablePestControlResponse, error) {

	var responseSustainablesPestControl []responses.SustainablePestControlResponse
	entitySustainablePestControl, err := s.sustainablePestControlRepository.FindAllSustainablePestControl()
	if err != nil {
		return nil, fmt.Errorf("erro: %v", err)
	}

	for _, v := range entitySustainablePestControl {
		sustainablePestControl := responses.SustainablePestControlResponse{
			Id:   v.Id,
			Name: v.Name,
		}

		responseSustainablesPestControl = append(responseSustainablesPestControl, sustainablePestControl)
	}

	return responseSustainablesPestControl, nil

}

func (s *SustainablePestControlService) PostSustainablePestControl(requestSustainablePestControl requests.SustainablePestControlRequest) error {

	entitySustainablePestControl := entities.SustainablePestControlEntity{
		Name: strings.ToUpper(requestSustainablePestControl.Name),
	}

	if err := s.sustainablePestControlRepository.CreateSustainablePestControl(entitySustainablePestControl); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}

func (s *SustainablePestControlService) GetFindByIdSustainablePestControl(id uint) (responses.SustainablePestControlResponse, error) {

	var responseSustainablesPestControl responses.SustainablePestControlResponse

	entitySustainablePestControl, err := s.sustainablePestControlRepository.FindByIdSustainablePestControl(id)
	if err != nil {
		return responseSustainablesPestControl, fmt.Errorf("erro: %v", err)
	}

	responseSustainablesPestControl = responses.SustainablePestControlResponse{
		Id:   entitySustainablePestControl.Id,
		Name: entitySustainablePestControl.Name,
	}

	return responseSustainablesPestControl, nil

}

func (s *SustainablePestControlService) PutSustainablePestControl(id uint, requestSustainablePestControl requests.SustainablePestControlRequest) error {

	entitySustainablePestControl := entities.SustainablePestControlEntity{
		Name: requestSustainablePestControl.Name,
	}

	if err := s.sustainablePestControlRepository.UpdateSustainablePestControl(id, entitySustainablePestControl); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}

func (s *SustainablePestControlService) DeleteSustainablePestControl(id uint) error {

	if err := s.sustainablePestControlRepository.DeleteSustainablePestControl(id); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}
