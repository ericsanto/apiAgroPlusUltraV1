package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"gorm.io/gorm"
)


type SoilTypeRepository struct {

  db *gorm.DB
}


func NewSoilRepository(db *gorm.DB) *SoilTypeRepository {
  return &SoilTypeRepository{db:db}
}

func(r *SoilTypeRepository) FindAll() ([]entities.SoilTypeEntity, error) {

  var soilTypes []entities.SoilTypeEntity

  err := r.db.Find(&soilTypes).Error

  return soilTypes, err
} 

func(r *SoilTypeRepository) FindById(id uint) (entities.SoilTypeEntity, error) {

  var soilType entities.SoilTypeEntity

  err := r.db.First(&soilType, id).Error

  return soilType, err
}

func(r *SoilTypeRepository) Create(soilTypeModel *entities.SoilTypeEntity) error {


  result := r.db.Create(soilTypeModel)

  if result.Error != nil {
    return fmt.Errorf("Erro ao criar no banco de dados %w", result.Error )
  } 

  if result.RowsAffected == 0 {
    return fmt.Errorf("nenhum registro foi criado")
  }

  return nil
}

func(r *SoilTypeRepository) Update(id uint, soilTypeModel entities.SoilTypeEntity) error {

  soilTypeModelExist, err := r.FindById(id)
  if err != nil {
    return fmt.Errorf("Erro: %W", err)
  }

  result := r.db.Model(&entities.SoilTypeEntity{}).Where("id = ?", soilTypeModelExist.Id).Updates(soilTypeModel)

  if result.Error != nil {
    return fmt.Errorf("Erro ao atualizaer dados %w", result.Error)
  }

  return nil
}

func(r *SoilTypeRepository) Delete(id uint) error {

  soilTypeModelExist, err := r.FindById(id);
  if err != nil {
    return fmt.Errorf("Erro ao deletar tipo de solo. Id não existe")
  }

  result := r.db.Where("id = ?", soilTypeModelExist.Id).Delete(&entities.SoilTypeEntity{})

  if result.Error != nil {
    return fmt.Errorf("Não foi possível deletar tipo de solo %w", err)
  }

  return nil
}
