package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"gorm.io/gorm"
)

type AgricultureCultureInterface interface {
  
  FindAllAgricultureCulture() ([]entities.AgricultureCultureEntity, error)
  FindByIdAgricultureCulture(id uint) (entities.AgricultureCultureEntity, error)
  CreateAgricultureCulture(agricultureCulture *entities.AgricultureCultureEntity)
  UpdateAgricultureCulture(id uint, agricultureCulture entities.AgricultureCultureEntity) error
  DeleteAgricultureCulture(id uint)
}

type AgricultureCultureRepository struct {

  db *gorm.DB
}

func NewAgricultureCultureRepository(db *gorm.DB) *AgricultureCultureRepository {

  return &AgricultureCultureRepository{db:db}
}

func(a *AgricultureCultureRepository) FindAllAgricultureCulture() ([]entities.AgricultureCultureEntity, error) {

  var agricultureCultures []entities.AgricultureCultureEntity
  err := a.db.Find(&agricultureCultures)

  if err.Error != nil {
    return agricultureCultures, fmt.Errorf("Erro ao buscar dados")
  }

  return agricultureCultures, nil

}

func(a *AgricultureCultureRepository) FindByIdAgricultureCulture(id uint) (entities.AgricultureCultureEntity, error) {

  var agricultureCulture entities.AgricultureCultureEntity

  err := a.db.First(&agricultureCulture, id)

  if err.Error != nil {
    return agricultureCulture, fmt.Errorf("Erro ao buscar cultura agrícola. Id não existe no banco de dados")
  }

  return agricultureCulture, nil
}

func(a *AgricultureCultureRepository) CreateAgricultureCulture(agriculutreCulture *entities.AgricultureCultureEntity) error {


  err := a.db.Create(&agriculutreCulture)

  if err.Error != nil {
    return fmt.Errorf("Não foi possivel salvar tipo de cultura no banco de dados %w", err.Error)
  }

  return nil
}


func(r *AgricultureCultureRepository) UpdateAgricultureCulture(id uint, agricultureCultureEntity entities.AgricultureCultureEntity) error {

  agricultureCultureExists, err := r.FindByIdAgricultureCulture(id)
  if err != nil {
    return fmt.Errorf(err.Error())
  }

  result := r.db.Model(&entities.AgricultureCultureEntity{}).Where("id = ?", agricultureCultureExists.Id).Updates(agricultureCultureEntity) 
  if result.Error != nil {
    return fmt.Errorf("Erro ao atualizar cultura agrícola: %w", result.Error) 
  }
  
  return nil
}

func(r *AgricultureCultureRepository) DeleteAgricultureCulture(id uint) error {

  agricultureCultureExists, err := r.FindByIdAgricultureCulture(id)
  if err != nil {
    return fmt.Errorf(err.Error())
  }

  result := r.db.Where("id = ?", agricultureCultureExists.Id).Delete(&entities.AgricultureCultureEntity{})
  if result.Error != nil {
    return fmt.Errorf("Erro no repositório: %w", result.Error)
  }

  return nil
}

