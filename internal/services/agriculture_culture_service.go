package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)


type AgricultureCultureInterface interface {

  FindAllAgricultureCultureService() ([]responses.AgricultureCultureResponse, error)
  FindByIdAgricultureCultureService() (responses.AgricultureCultureResponse, error)
  CreateAgricultureCultureService(agricultureCulture requests.AgricultureCultureRequest) (error)
  UpdateAgricultureCultureService(id uint, agricultureCulture requests.AgricultureCultureRequest) (error)
  DeleteAgricultureCultureService(id uint) (error)
}


type AgricultureCultureService struct {

  agricultureRepository *repositories.AgricultureCultureRepository
}

func NewAgricultureCultureService(agricultureCultureRepository *repositories.AgricultureCultureRepository) *AgricultureCultureService {

  return &AgricultureCultureService{agricultureRepository: agricultureCultureRepository}
}

func(a *AgricultureCultureService) FindAllAgricultureCultureService() ([]responses.AgricultureCultureResponse, error) {

  var agriculturesCulturesResponses []responses.AgricultureCultureResponse
  result, err := a.agricultureRepository.FindAllAgricultureCulture()

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

func(a *AgricultureCultureService) FindByIdAgricultureCultureService(id uint) (responses.AgricultureCultureResponse, error) {

  var agricultureCultureResponse responses.AgricultureCultureResponse
  result, err := a.agricultureRepository.FindByIdAgricultureCulture(id)

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

func(a *AgricultureCultureService) CreateAgricultureCultureService(agricultureCulture requests.AgricultureCultureRequest) error {

  agricultureCultureEntity := entities.AgricultureCultureEntity{
    Name: agricultureCulture.Name,
    NameCientific: agricultureCulture.NameCientific,
    SoilTypeId: agricultureCulture.SoilTypeId,
    PhIdealSoil: agricultureCulture.PhIdealSoil,
    MaxTemperature: agricultureCulture.MaxTemperature,
    MinTemperature: agricultureCulture.MinTemperature,
    ExcellentTemperature: agricultureCulture.ExcellentTemperature,
    WeeklyWaterRequirememntMax: agricultureCulture.WeeklyWaterRequirementMax,
    WeeklyWaterRequirememntMin: agricultureCulture.WeeklyWaterRequirementMin,
    SunlightRequirement: agricultureCulture.SunlightRequirement,
  }

  err := a.agricultureRepository.CreateAgricultureCulture(&agricultureCultureEntity)

  if err != nil {
    return fmt.Errorf("Erro ao criar cultura agrícola no repositório: %w", err)
  }

  return nil
}

func(a *AgricultureCultureService) PutAgricultureCultureService(id uint, agricultureCulture requests.AgricultureCultureRequest) error {

  agricultureCultureEntity := entities.AgricultureCultureEntity{
    Name: agricultureCulture.Name,
    NameCientific: agricultureCulture.NameCientific,
    SoilTypeId: agricultureCulture.SoilTypeId,
    PhIdealSoil: agricultureCulture.PhIdealSoil,
    MaxTemperature: agricultureCulture.MaxTemperature,
    MinTemperature: agricultureCulture.MinTemperature,
    ExcellentTemperature: agricultureCulture.ExcellentTemperature,
    WeeklyWaterRequirememntMax: agricultureCulture.WeeklyWaterRequirementMax,
    WeeklyWaterRequirememntMin: agricultureCulture.WeeklyWaterRequirementMin,
    SunlightRequirement: agricultureCulture.SunlightRequirement,
  }
 
  if err := a.agricultureRepository.UpdateAgricultureCulture(id, agricultureCultureEntity); err != nil {
    return fmt.Errorf("Erro ao atualizar cultura agrícola no repositório %w", err)
  }  

  return nil

}

func(a *AgricultureCultureService) DeleteAgricultureCultureService(id uint) error {

  if err := a.agricultureRepository.DeleteAgricultureCulture(id); err != nil {
    return fmt.Errorf(err.Error())
  }
  
  return nil
}
