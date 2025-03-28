package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"gorm.io/gorm"
)

type TypePestRepositoryInterface interface {

  FindAllTypePest() ([]entities.TypePestEntity, error)
  FindByIdTypePest(id uint) (*entities.TypePestEntity, error)
  CreateTypePest(typePestEntity entities.TypePestEntity) error 
  UpdateTypePest(id uint, typePestEntity entities.TypePestEntity) error
  DeleteTypePest(id uint) error
}


type TypePestRepository struct {

  db *gorm.DB
}


func NewTypePestRepository(db *gorm.DB) *TypePestRepository {
  return &TypePestRepository{db: db}
}

func(t *TypePestRepository) FindAllTypePest() ([]entities.TypePestEntity, error) {
  
  var typesPests []entities.TypePestEntity

  err := t.db.Find(&typesPests)
  if err.Error != nil {
    return nil, fmt.Errorf("Erro ao buscar pragas no banco de dados")  
  }

  return typesPests, nil 
}

func(t *TypePestRepository) FindByIdTypePest(id uint) (entities.TypePestEntity, error) {

  var typePest entities.TypePestEntity

  err := t.db.First(&typePest, id)
  if err.Error != nil {
    return typePest, fmt.Errorf("Não existe praga com esse id") 
  }

  return typePest, nil
}

func(t *TypePestRepository) CreateTypePest(typePestEntity *entities.TypePestEntity) error {

  err := t.db.Create(&typePestEntity)
  if err.Error != nil {
    return fmt.Errorf("Não foi possível cadastrar praga")
  }

  return nil
}

func(t *TypePestRepository) UpdateTypePest(id uint, entityTypePest entities.TypePestEntity) error {

  typePestExists, err := t.FindByIdTypePest(id)
  if err != nil {
    return fmt.Errorf("Erro no repositório: %w", err)
  }

  result := t.db.Model(&entities.TypePestEntity{}).Where("id = ?", typePestExists.Id).Updates(entityTypePest)
  if result.Error != nil {
    return fmt.Errorf("Erro ao atualizar: %w", err)
  }

  return nil
}

func(t *TypePestRepository) DeleteTypePest(id uint) error {

  typePestExists, err := t.FindByIdTypePest(id)
  if err != nil {
    return fmt.Errorf("Erro ao buscar praga: %w", err)
  }

  result := t.db.Where("id = ?", typePestExists.Id).Delete(typePestExists)
  if result.Error != nil {
    return fmt.Errorf("Erro ao deletar praga.")
  }

  return nil
}
