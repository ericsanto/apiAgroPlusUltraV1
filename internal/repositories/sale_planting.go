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
	CreateSalePlantingRepository(entitySalePlanting entities.SalePlantingEntity) error
	FindAllSalePlanting() ([]entities.SalePlantingEntity, error)
	FindSalePlantingByID(id uint) (*entities.SalePlantingEntity, error)
	UpdateSalePlanting(id uint, entitySalePlanting entities.SalePlantingEntity) error
	DeleteSalePlanting(id uint) error
}

type SalePlantingRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewSalePlantingRepository(db interfaces.GORMRepositoryInterface) SalePlantingRepositoryInterface {
	return &SalePlantingRepository{db: db}
}

func (s *SalePlantingRepository) CreateSalePlantingRepository(entitySalePlanting entities.SalePlantingEntity) error {

	if err := s.db.Create(&entitySalePlanting).Error; err != nil {
		if myerror.IsUniqueConstraintViolated(err) {
			return fmt.Errorf("%w %d", myerror.ErrDuplicateSale, entitySalePlanting.PlantingID)
		}
		return fmt.Errorf("erro ao cadastrar venda")
	}

	return nil

}

func (s *SalePlantingRepository) FindAllSalePlanting() ([]entities.SalePlantingEntity, error) {

	var entitiesSalePlanting []entities.SalePlantingEntity

	if err := s.db.Find(&entitiesSalePlanting).Error; err != nil {
		return nil, fmt.Errorf("não foi possível buscar todas as vendas de plantações")
	}

	return entitiesSalePlanting, nil
}

func (s *SalePlantingRepository) FindSalePlantingByID(id uint) (*entities.SalePlantingEntity, error) {

	var entitySalePlanting entities.SalePlantingEntity

	if err := s.db.First(&entitySalePlanting, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w %d", myerror.ErrNotFoundSale, id)
		}

		return nil, fmt.Errorf("erro ao buscar venda com id %d", id)
	}

	return &entitySalePlanting, nil
}

func (s *SalePlantingRepository) UpdateSalePlanting(id uint, entitySalePlanting entities.SalePlantingEntity) error {

	if _, err := s.FindSalePlantingByID(id); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := s.db.Model(&entities.SalePlantingEntity{}).Where("id = ?", id).Updates(&entitySalePlanting).Error; err != nil {

		switch {
		case myerror.IsUniqueConstraintViolated(err):
			return fmt.Errorf("%w %d", myerror.ErrDuplicateSale, id)

		case myerror.IsViolatedForeingKeyConstraint(err):
			message := myerror.InterpolationErrViolatedForeingKey("plantio com id", entitySalePlanting.PlantingID)
			return fmt.Errorf("%w %s", myerror.ErrViolatedForeingKey, message)

		default:
			return fmt.Errorf("erro ao atualizar venda")
		}

	}

	return nil
}

func (s *SalePlantingRepository) DeleteSalePlanting(id uint) error {

	entitySalePlanting, err := s.FindSalePlantingByID(id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := s.db.Where("id = ?", id).Delete(&entitySalePlanting).Error; err != nil {
		return fmt.Errorf("erro ao deletar venda")
	}

	return nil
}
