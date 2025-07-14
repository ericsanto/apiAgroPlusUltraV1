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
	FindAllProductinCostRepository(batchID, farmID, userID, plantingID uint) ([]entities.ProductionCostEntity, error)
	CreateProductionCost(entityProductionCost entities.ProductionCostEntity) error
	FindProductionCostByID(batchID, farmID, userID, plantingID, productionCostID uint) (*entities.ProductionCostEntity, error)
	UpdateProductionCost(batchID, farmID, userID, plantingID, productionCostID uint, entityProductCost entities.ProductionCostEntity) error
	DeleteProductionCost(batchID, farmID, userID, plantingID, productionCostID uint) error
}

type ProductionCostRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewProductionCostRepository(db interfaces.GORMRepositoryInterface) ProductionCostRepositoryInterface {
	return &ProductionCostRepository{db: db}
}

func (p *ProductionCostRepository) FindAllProductinCostRepository(batchID, farmID, userID, plantingID uint) ([]entities.ProductionCostEntity, error) {

	var entityProductionCost []entities.ProductionCostEntity

	if err := p.db.Model(entities.ProductionCostEntity{}).
		Joins("JOIN planting_entities ON planting_entities.id = production_cost_entities.planting_id").
		Joins("JOIN batch_entities ON batch_entities.id = planting_entities.batch_id").
		Joins("JOIN farm_entities ON farm_entities.id = batch_entities.farm_id").
		Joins("JOIN user_models ON user_models.id = farm_entities.user_id").
		Where("planting_entities.id = ? AND batch_entities.id = ? AND user_models.id = ? AND farm_entities.id = ?",
			plantingID, batchID, userID, farmID).
		Find(&entityProductionCost).Error; err != nil {
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

func (p *ProductionCostRepository) FindProductionCostByID(batchID, farmID, userID, plantingID, productionCostID uint) (*entities.ProductionCostEntity, error) {

	var entityProductionCost entities.ProductionCostEntity

	if err := p.db.Model(entities.ProductionCostEntity{}).
		Joins("JOIN planting_entities ON planting_entities.id = production_cost_entities.planting_id").
		Joins("JOIN batch_entities ON batch_entities.id = planting_entities.batch_id").
		Joins("JOIN farm_entities ON farm_entities.id = batch_entities.farm_id").
		Joins("JOIN user_models ON user_models.id = farm_entities.user_id").
		Where("planting_entities.id = ? AND batch_entities.id = ? AND user_models.id = ? AND farm_entities.id = ? AND production_cost_entities.id = ?",
			plantingID, batchID, userID, farmID, productionCostID).First(&entityProductionCost).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("não existe custo com o id %d", productionCostID)
		}

		return nil, fmt.Errorf("erro ao buscar custo: %w", err)
	}

	return &entityProductionCost, nil
}

func (p *ProductionCostRepository) UpdateProductionCost(batchID, farmID, userID, plantingID, productionCostID uint, entityProductCost entities.ProductionCostEntity) error {

	if _, err := p.FindProductionCostByID(batchID, farmID, userID, plantingID, productionCostID); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	if err := p.db.Model(&entities.ProductionCostEntity{}).Where("id = ?", productionCostID).Updates(&entityProductCost).Error; err != nil {
		return fmt.Errorf("erro ao atualizar custo")
	}

	return nil
}

func (p *ProductionCostRepository) DeleteProductionCost(batchID, farmID, userID, plantingID, productionCostID uint) error {

	productEntity, err := p.FindProductionCostByID(batchID, farmID, userID, plantingID, productionCostID)
	if err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	if err := p.db.Where("id = ?", productionCostID).Delete(&productEntity).Error; err != nil {
		return fmt.Errorf("erro ao tentar deletar custo")
	}

	return nil
}
