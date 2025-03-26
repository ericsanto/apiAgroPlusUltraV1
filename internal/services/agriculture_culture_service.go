package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)


type AgricultureCultureInterface interface {

  FindAllAgricultureCulture() ([]responses.AgricultureCultureResponse, error)
  FindByIdAgricultureCulture() (responses.AgricultureCultureResponse, error)
  CreateAgricultureCulture(agricultureCulture responses.AgricultureCultureResponse) (error)
  UpdateAgricultureCulture(id uint, agricultureCulture responses.AgricultureCultureResponse) (error)
  DeleteAgricultureCulture(id uint) (error)
}


type AgricultureCultureService struct {

  agricultureRepository *repositories.AgricultureCultureRepository
}

func NewAgricultureCultureService(agricultureCultureRepository *repositories.AgricultureCultureRepository) *AgricultureCultureService {

  return &AgricultureCultureService{agricultureRepository: agricultureCultureRepository}
}

func(a *AgricultureCultureService) FindAllAgricultureCulture() ([]responses.AgricultureCultureResponse, error) {

  var agriculturesCulturesResponses []responses.AgricultureCultureResponse
  result, err := a.agricultureRepository.FindAll()

  if err != nil {
    return agriculturesCulturesResponses, fmt.Errorf("Erro ao buscar no repositório %w", err)
  }

  for _, v := range result{

    agricultureCultureResponse := responses.AgricultureCultureResponse{
      Id: v.Id,
      Name: v.Name,
      NameCientific: v.NameCientific,
      SoilTypeId: v.SoilTypeId,
      PhIdealSoil: v.PhIdealSoil,
      MaxTemperature: v.MaxTemperature,
      MinTemperature: v.MinTemperature,
      ExcellentTemperature: v.ExcellentTemperature,
      WeeklyWaterRequirememntMax: v.WeeklyWaterRequirememntMax,
      WeeklyWaterRequirememntMin: v.WeeklyWaterRequirememntMin,
      SunlightRequirement: v.SunlightRequirement,
    }
    
    agriculturesCulturesResponses = append(agriculturesCulturesResponses, agricultureCultureResponse)
    
  }

  return agriculturesCulturesResponses, nil
}

func(a *AgricultureCultureService) FindByIdAgricultureCulture(id uint) (responses.AgricultureCultureResponse, error) {

  var agricultureCultureResponse responses.AgricultureCultureResponse
  result, err := a.agricultureRepository.FindById(id)

  if err != nil {
    return agricultureCultureResponse, fmt.Errorf("Erro ao buscar cultura agrícola com esse id no repositório: %w", err)
  }

  agricultureCultureResponse = responses.AgricultureCultureResponse{
    Id: result.Id,
    Name: result.Name,
    NameCientific: result.NameCientific,
    SoilTypeId: result.SoilTypeId,
    PhIdealSoil: result.PhIdealSoil,
    MaxTemperature: result.MaxTemperature,
    MinTemperature: result.MinTemperature,
    ExcellentTemperature: result.ExcellentTemperature,
    WeeklyWaterRequirememntMax: result.WeeklyWaterRequirememntMax,
    WeeklyWaterRequirememntMin: result.WeeklyWaterRequirememntMin,
    SunlightRequirement: result.SunlightRequirement,
  }

  return agricultureCultureResponse, nil
}

func(a *AgricultureCultureService) CreateAgricultureCulture(agricultureCulture requests.AgricultureCultureRequest) error {

  agricultureCultureEntity := entities.AgricultureCultureEntity{
    Name: agricultureCulture.Name,
    NameCientific: agricultureCulture.NameCientific,
    SoilTypeId: agricultureCulture.SoilTypeId,
    PhIdealSoil: agricultureCulture.PhIdealSoil,
    MaxTemperature: agricultureCulture.MaxTemperature,
    MinTemperature: agricultureCulture.MinTemperature,
    ExcellentTemperature: agricultureCulture.ExcellentTemperature,
    WeeklyWaterRequirememntMax: agricultureCulture.WeeklyWaterRequirememntMax,
    WeeklyWaterRequirememntMin: agricultureCulture.WeeklyWaterRequirememntMin,
    SunlightRequirement: agricultureCulture.SunlightRequirement,
  }

  err := a.agricultureRepository.Create(&agricultureCultureEntity)

  if err != nil {
    return fmt.Errorf("Erro ao criar cultura agrícola no repositório: %w", err)
  }

  return nil
}
