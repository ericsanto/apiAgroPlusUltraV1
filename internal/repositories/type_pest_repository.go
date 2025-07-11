package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

type TypePestRepositoryInterface interface {
	FindAllTypePest() ([]entities.TypePestEntity, error)
	FindByIdTypePest(id uint) (*entities.TypePestEntity, error)
	CreateTypePest(typePestEntity entities.TypePestEntity) error
	UpdateTypePest(id uint, typePestEntity entities.TypePestEntity) error
	DeleteTypePest(id uint) error
}

type TypePestRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewTypePestRepository(db interfaces.GORMRepositoryInterface) TypePestRepositoryInterface {
	return &TypePestRepository{db: db}
}

func (t *TypePestRepository) FindAllTypePest() ([]entities.TypePestEntity, error) {

	var typesPests []entities.TypePestEntity

	err := t.db.Find(&typesPests)
	if err.Error != nil {
		return nil, fmt.Errorf("erro ao buscar pragas no banco de dados")
	}

	return typesPests, nil
}

func (t *TypePestRepository) FindByIdTypePest(id uint) (*entities.TypePestEntity, error) {

	var typePest entities.TypePestEntity

	err := t.db.First(&typePest, id)
	if err.Error != nil {
		return &typePest, fmt.Errorf("não existe praga com esse id")
	}

	return &typePest, nil
}

func (t *TypePestRepository) CreateTypePest(typePestEntity entities.TypePestEntity) error {

	err := t.db.Create(&typePestEntity)
	if err.Error != nil {
		return fmt.Errorf("não foi possível cadastrar praga")
	}

	return nil
}

func (t *TypePestRepository) UpdateTypePest(id uint, entityTypePest entities.TypePestEntity) error {

	_, err := t.FindByIdTypePest(id)
	if err != nil {
		return fmt.Errorf("erro no repositório: %w", err)
	}

	result := t.db.Model(&entities.TypePestEntity{}).Where("id = ?", id).Updates(&entityTypePest)
	if result.Error != nil {
		return fmt.Errorf("erro ao atualizar: %w", err)
	}

	return nil
}

func (t *TypePestRepository) DeleteTypePest(id uint) error {

	typePestExists, err := t.FindByIdTypePest(id)
	if err != nil {
		return fmt.Errorf("erro ao buscar praga: %w", err)
	}

	result := t.db.Where("id = ?", id).Delete(&typePestExists)
	if result.Error != nil {
		return fmt.Errorf("erro ao deletar praga")
	}

	return nil
}
