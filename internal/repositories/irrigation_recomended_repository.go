package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type IrrigationRecomendedRepositoryInterface interface {
	FindAllIrrigationRecomended() ([]entities.IrrigationRecomendedEntity, error)
	FindByIdIrrigationRecomended(id uint) (*entities.IrrigationRecomendedEntity, error)
	CreateIrrigationRecomended(entityIrrigationRecomended entities.IrrigationRecomendedEntity) error
	UpdateIrrigationRecomended(id uint, entityIrrigationRecomended entities.IrrigationRecomendedEntity) error
	DeleteIrrigationRecomended(id uint)
}

type IrrigationRecomendedRepository struct {
	db *gorm.DB
}

func NewIrrigationRecomdedRepository(db *gorm.DB) *IrrigationRecomendedRepository {
	return &IrrigationRecomendedRepository{db: db}
}

func (i *IrrigationRecomendedRepository) FindAllIrrigationRecomended() ([]entities.IrrigationRecomendedEntity, error) {

	var entityIrrigationRecomended []entities.IrrigationRecomendedEntity

	if err := i.db.Find(&entityIrrigationRecomended).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar todas irrigações: %v", err)
	}

	return entityIrrigationRecomended, nil
}

func (i *IrrigationRecomendedRepository) CreateIrrigationRecomended(entityIrrigationRecomended entities.IrrigationRecomendedEntity) error {

	if err := i.db.Create(&entityIrrigationRecomended).Error; err != nil {
		return fmt.Errorf("erro ao criar irrigação: %v", err)
	}

	return nil
}

func (i *IrrigationRecomendedRepository) FindByIdIrrigationRecomended(id uint) (*entities.IrrigationRecomendedEntity, error) {

	var entityIrrigationRecomended entities.IrrigationRecomendedEntity

	if err := i.db.First(&entityIrrigationRecomended, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("objeto com o id fornecido não existe")
		}
		return nil, fmt.Errorf("erro ao buscar objeto: %v", err)
	}

	return &entityIrrigationRecomended, nil
}

func (i *IrrigationRecomendedRepository) UpdateIrrigationRecomended(id uint, entityIrrigationRecomended entities.IrrigationRecomendedEntity) error {

	_, err := i.FindByIdIrrigationRecomended(id)
	if err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	if err := i.db.Model(&entities.IrrigationRecomendedEntity{}).Where("id = ?", id).
		Updates(&entityIrrigationRecomended).Error; err != nil {
		return fmt.Errorf("não foi possível atualizar objeto: %v", err)
	}

	return nil
}

func (i *IrrigationRecomendedRepository) DeleteIrrigationRecomended(id uint) error {

	entityIrrigationRecomended, err := i.FindByIdIrrigationRecomended(id)
	if err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	if err := i.db.Where("id = ?", entityIrrigationRecomended.Id).Delete(&entityIrrigationRecomended).Error; err != nil {
		return fmt.Errorf("não foi possível deletar objeto: %v", err)
	}

	return nil
}
