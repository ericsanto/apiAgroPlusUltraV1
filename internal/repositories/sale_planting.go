package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type SalePlantingRepositoryInterface interface {
	CreateSalePlantingRepository(batchID, farmID, userID, plantingID uint, entitySalePlanting entities.SalePlantingEntity) error
	FindAllSalePlanting(batchID, farmID, userID uint) ([]entities.SalePlantingEntity, error)
	FindSalePlantingByID(batchID, farmID, userID, salePlantingID uint) (*entities.SalePlantingEntity, error)
	UpdateSalePlanting(batchID, farmID, userID, salePlantingID uint, entitySalePlanting entities.SalePlantingEntity) error
	DeleteSalePlanting(batchID, farmID, userID, salePlantingID uint) error
}

type SalePlantingRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewSalePlantingRepository(db interfaces.GORMRepositoryInterface) SalePlantingRepositoryInterface {
	return &SalePlantingRepository{db: db}
}

func (s *SalePlantingRepository) CreateSalePlantingRepository(batchID, farmID, userID, plantingID uint, entitySalePlanting entities.SalePlantingEntity) error {

	if err := s.db.Model(entities.SalePlantingEntity{}).
		Joins("JOIN planting_entities ON planting_entities.id = sale_planting_entities.planting_id").
		Joins("JOIN batch_entities ON  batch_entities.id = planting_entities.batch_id").
		Joins("JOIN farm_entities ON farm_entities.id = batch_entities.farm_id").
		Joins("JOIN user_models ON user_models.id = farm_entities.user_id").
		Where("planting_entities.id = ? AND batch_entities.id = ? AND farm_entities.id = ? AND user_models.id = ?",
			plantingID, batchID, farmID, userID).
		Create(&entitySalePlanting).Error; err != nil {
		if myerror.IsUniqueConstraintViolated(err) {
			return fmt.Errorf("%w %d", myerror.ErrDuplicateSale, entitySalePlanting.PlantingID)
		}
		return fmt.Errorf("erro ao cadastrar venda")
	}

	return nil

}

func (s *SalePlantingRepository) FindAllSalePlanting(batchID, farmID, userID uint) ([]entities.SalePlantingEntity, error) {

	var entitiesSalePlanting []entities.SalePlantingEntity

	if err := s.db.Model(entities.SalePlantingEntity{}).
		Joins("JOIN planting_entities ON planting_entities.id = sale_planting_entities.planting_id").
		Joins("JOIN batch_entities ON  batch_entities.id = planting_entities.batch_id").
		Joins("JOIN farm_entities ON farm_entities.id = batch_entities.farm_id").
		Joins("JOIN user_models ON user_models.id = farm_entities.user_id").
		Where("batch_entities.id = ? AND farm_entities.id = ? AND user_models.id = ?",
			batchID, farmID, userID).Find(&entitiesSalePlanting).Error; err != nil {
		return nil, fmt.Errorf("não foi possível buscar todas as vendas de plantações")
	}

	return entitiesSalePlanting, nil
}

func (s *SalePlantingRepository) FindSalePlantingByID(batchID, farmID, userID, salePlantingID uint) (*entities.SalePlantingEntity, error) {

	var entitySalePlanting entities.SalePlantingEntity

	if err := s.db.Model(entities.SalePlantingEntity{}).
		Joins("JOIN planting_entities ON planting_entities.id = sale_planting_entities.planting_id").
		Joins("JOIN batch_entities ON  batch_entities.id = planting_entities.batch_id").
		Joins("JOIN farm_entities ON farm_entities.id = batch_entities.farm_id").
		Joins("JOIN user_models ON user_models.id = farm_entities.user_id").
		Where("batch_entities.id = ? AND farm_entities.id = ? AND user_models.id = ? AND sale_planting_entities.id = ?",
			batchID, farmID, userID, salePlantingID).First(&entitySalePlanting).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w %d", myerror.ErrNotFoundSale, salePlantingID)
		}

		return nil, fmt.Errorf("erro ao buscar venda com id %d", salePlantingID)
	}

	return &entitySalePlanting, nil
}

func (s *SalePlantingRepository) UpdateSalePlanting(batchID, farmID, userID, salePlantingID uint, entitySalePlanting entities.SalePlantingEntity) error {

	if _, err := s.FindSalePlantingByID(batchID, farmID, userID, salePlantingID); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := s.db.Model(&entities.SalePlantingEntity{}).Where("id = ?", salePlantingID).Updates(&entitySalePlanting).Error; err != nil {

		switch {
		case myerror.IsUniqueConstraintViolated(err):
			return fmt.Errorf("%w %d", myerror.ErrDuplicateSale, salePlantingID)

		case myerror.IsViolatedForeingKeyConstraint(err):
			message := myerror.InterpolationErrViolatedForeingKey("plantio com id", entitySalePlanting.PlantingID)
			return fmt.Errorf("%w %s", myerror.ErrViolatedForeingKey, message)

		default:
			return fmt.Errorf("erro ao atualizar venda")
		}

	}

	return nil
}

func (s *SalePlantingRepository) DeleteSalePlanting(batchID, farmID, userID, salePlantingID uint) error {

	entitySalePlanting, err := s.FindSalePlantingByID(batchID, farmID, userID, salePlantingID)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := s.db.Where("id = ?", salePlantingID).Delete(&entitySalePlanting).Error; err != nil {
		return fmt.Errorf("erro ao deletar venda")
	}

	return nil
}
