package repositories

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

type ProductionCostRepositoryInterface interface {
	FindAllProductinCostRepository() ([]entities.ProductionCostEntity, error)
	CreateProductionCost(entityProductionCost entities.ProductionCostEntity) error
	FindProductionCostByID(id uint) (*entities.ProductionCostEntity, error)
	UpdateProductionCost(id uint, entityProductCost entities.ProductionCostEntity) error
	DeleteProductionCost(id uint) error
}

type ProductionCostRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewProductionCostRepository(db interfaces.GORMRepositoryInterface) ProductionCostRepositoryInterface {
	return &ProductionCostRepository{db: db}
}

func (p *ProductionCostRepository) FindAllProductinCostRepository() ([]entities.ProductionCostEntity, error) {

	var entityProductionCost []entities.ProductionCostEntity

	if err := p.db.Find(&entityProductionCost).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar custos de produtos %w", err)
	}

	return entityProductionCost, nil
}

func (p *ProductionCostRepository) CreateProductionCost(entityProductionCost entities.ProductionCostEntity) error {

	if err := p.db.Create(&entityProductionCost).Error; err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return fmt.Errorf("plantio com o ID fornecido não existe")
		}
		return fmt.Errorf("erro ao tentar criar custo: %w", err)
	}

	return nil
}

func (p *ProductionCostRepository) FindProductionCostByID(id uint) (*entities.ProductionCostEntity, error) {

	var entityProductionCost entities.ProductionCostEntity

	if err := p.db.First(&entityProductionCost, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("não existe custo com o id %d", id)
		}

		return nil, fmt.Errorf("erro ao buscar custo: %w", err)
	}

	return &entityProductionCost, nil
}

func (p *ProductionCostRepository) UpdateProductionCost(id uint, entityProductCost entities.ProductionCostEntity) error {

	if _, err := p.FindProductionCostByID(id); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	if err := p.db.Model(&entities.ProductionCostEntity{}).Where("id = ?", id).Updates(&entityProductCost).Error; err != nil {
		return fmt.Errorf("erro ao atualizar custo")
	}

	return nil
}

func (p *ProductionCostRepository) DeleteProductionCost(id uint) error {

	productEntity, err := p.FindProductionCostByID(id)
	if err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	if err := p.db.Where("id = ?", id).Delete(&productEntity).Error; err != nil {
		return fmt.Errorf("erro ao tentar deletar custo")
	}

	return nil
}
