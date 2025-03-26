package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"gorm.io/gorm"
)


type AgricultureCultureRepository struct {

  db *gorm.DB
}

func NewAgricultureCultureRepository(db *gorm.DB) *AgricultureCultureRepository {

  return &AgricultureCultureRepository{db:db}
}

func(a *AgricultureCultureRepository) FindAll() ([]entities.AgricultureCultureEntity, error) {

  var agricultureCultures []entities.AgricultureCultureEntity
  err := a.db.Find(&agricultureCultures)

  if err.Error != nil {
    return agricultureCultures, fmt.Errorf("Erro ao buscar dados")
  }

  return agricultureCultures, nil

}

func(a *AgricultureCultureRepository) FindById(id uint) (entities.AgricultureCultureEntity, error) {

  var agricultureCulture entities.AgricultureCultureEntity

  err := a.db.First(&agricultureCulture, id)

  if err.Error != nil {
    return agricultureCulture, fmt.Errorf("Erro ao buscar cultura agrícola. Id não existe no banco de dados")
  }

  return agricultureCulture, nil
}

func(a *AgricultureCultureRepository) Create(agriculutreCulture *entities.AgricultureCultureEntity) error {


  err := a.db.Create(&agriculutreCulture)

  if err.Error != nil {
    return fmt.Errorf("Não foi possivel salvar tipo de cultura no banco de dados %w", err.Error)
  }

  return nil
}


