package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"

)


type TypePestService struct {

  typePestRepository *repositories.TypePestRepository
}

func NewTypePestService(typePestRepository *repositories.TypePestRepository) *TypePestService {

  return &TypePestService{typePestRepository:typePestRepository}
}

func(s *TypePestService) GetAllTypePest() ([]responses.TypePestResponse, error) {

  var typePests []responses.TypePestResponse

  result, err := s.typePestRepository.FindAllTypePest()
  if err != nil {
    return typePests, fmt.Errorf("Erro: %w", err)
  }

  for _, v := range result{
    
    typePest := responses.TypePestResponse{
      Id: v.Id,
      Name: v.Name,
    }

    typePests = append(typePests, typePest)
  }

  return typePests, nil
}

func(s *TypePestService) GetTypePestFindById(id uint) (responses.TypePestResponse, error){

  var typePest responses.TypePestResponse
  result, err := s.typePestRepository.FindByIdTypePest(id)
  if err != nil {
    return typePest, fmt.Errorf("Erro: %w", err)
  }

  typePest = responses.TypePestResponse{
    Id: result.Id,
    Name: result.Name,
  }

  return typePest, nil
}

func(s *TypePestService) CreateTypePest(requestTypePest requests.TypePestRequest) error {

  typePestEntity := entities.TypePestEntity{
    Name: requestTypePest.Name,
  }

  err := s.typePestRepository.CreateTypePest(&typePestEntity)
  if err != nil {
    return fmt.Errorf("Erro: %w", err)
  }

  return nil
}

func(s *TypePestService) PutTypePest(id uint, requestTypePest requests.TypePestRequest) error {
  
  typePestEntity := entities.TypePestEntity{
    Name: requestTypePest.Name,
  }

  err := s.typePestRepository.UpdateTypePest(id, typePestEntity)
  if err != nil {
    return fmt.Errorf("Erro: %w", err)
  }

  return nil
}

func(s *TypePestService) DeleteTypePest(id uint) error {

  err := s.typePestRepository.DeleteTypePest(id)
  if err != nil {
    return fmt.Errorf("Erro: %w", err)
  }

  return nil
}
