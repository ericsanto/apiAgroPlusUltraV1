package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type IrrigationTypeRepository struct {
	db *gorm.DB
}

func NewIrrigationTypeRepository(db *gorm.DB) *IrrigationTypeRepository {

	return &IrrigationTypeRepository{db: db}
}

func (it *IrrigationTypeRepository) FindAllIrrigationType() ([]entities.IrrigationTypeEntity, error) {

	var entitiesIrrigationType []entities.IrrigationTypeEntity

	if err := it.db.Find(&entitiesIrrigationType).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar tipo de irrigacao %w", err)
	}

	return entitiesIrrigationType, nil
}

func (it *IrrigationTypeRepository) CreateIrrigationType(entityIrrigationType entities.IrrigationTypeEntity) error {

	if err := it.db.Create(&entityIrrigationType).Error; err != nil {
		return fmt.Errorf("erro ao tentar criar tipo de irrigacao: %w", err)
	}

	return nil
}

func (p *IrrigationTypeRepository) FindIrrigationTypetByID(id uint) (*entities.IrrigationTypeEntity, error) {

	var entityIrrigationType entities.IrrigationTypeEntity

	if err := p.db.First(&entityIrrigationType, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("não existe tipo de irrigacao com o id %d", id)
		}

		return nil, fmt.Errorf("erro ao buscar tipo de irrigacao: %w", err)
	}

	return &entityIrrigationType, nil
}

func (it *IrrigationTypeRepository) UpdateIrrigationType(id uint, entityIrrigationType entities.IrrigationTypeEntity) error {

	if _, err := it.FindIrrigationTypetByID(id); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	if err := it.db.Model(&entities.IrrigationTypeEntity{}).Where("id = ?", id).Updates(&entityIrrigationType).Error; err != nil {
		return fmt.Errorf("erro ao atualizar tipo de irrigacao")
	}

	return nil
}

func (it *IrrigationTypeRepository) DeleteIrrigationType(id uint) error {

	irrigationTypeEntity, err := it.FindIrrigationTypetByID(id)
	if err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	if err := it.db.Where("id = ?", id).Delete(&irrigationTypeEntity).Error; err != nil {
		return fmt.Errorf("erro ao tentar deletar tipo de irrigacao")
	}

	return nil
}
