package repositories

import (
	"fmt"
	"strings"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

type SustainablePestControlRepositoryInterface interface {
	FindAllSustainablePestControl() ([]entities.SustainablePestControlEntity, error)
	CreateSustainablePestControl(entitySustainablePestControl entities.SustainablePestControlEntity) error
	FindByIdSustainablePestControl(id uint) (entities.SustainablePestControlEntity, error)
	UpdateSustainablePestControl(id uint, newEntitySustainablePestControl entities.SustainablePestControlEntity) error
	DeleteSustainablePestControl(id uint) error
}

type SustainablePestControlRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewSustainablePestControlRepository(db interfaces.GORMRepositoryInterface) SustainablePestControlRepositoryInterface {
	return &SustainablePestControlRepository{db: db}
}

func (s *SustainablePestControlRepository) FindAllSustainablePestControl() ([]entities.SustainablePestControlEntity, error) {

	var sustainablesPestControl []entities.SustainablePestControlEntity

	if err := s.db.Find(&sustainablesPestControl).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar todos os objetos: %v", err)
	}

	return sustainablesPestControl, nil
}

func (s *SustainablePestControlRepository) CreateSustainablePestControl(entitySustainablePestControl entities.SustainablePestControlEntity) error {

	if err := s.db.Create(&entitySustainablePestControl).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("objeto já existe com esse nome")
		}
		return fmt.Errorf("erro ao criar objeto: %v", err)
	}

	return nil

}

func (s *SustainablePestControlRepository) FindByIdSustainablePestControl(id uint) (entities.SustainablePestControlEntity, error) {

	var entitySustainablePestControl entities.SustainablePestControlEntity

	if err := s.db.First(&entitySustainablePestControl, id).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return entitySustainablePestControl, fmt.Errorf(fmt.Sprintf("obejto com id %d não existe", id), err)
		}

		return entitySustainablePestControl, fmt.Errorf("erro ao buscar objeto %v", err)
	}

	return entitySustainablePestControl, nil
}

func (s *SustainablePestControlRepository) UpdateSustainablePestControl(id uint, newEntitySustainablePestControl entities.SustainablePestControlEntity) error {

	entitySustainablePestControl, err := s.FindByIdSustainablePestControl(id)
	if err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	if err := s.db.Model(&entities.SustainablePestControlEntity{}).Where("id = ?", entitySustainablePestControl.Id).
		Updates(&newEntitySustainablePestControl).Error; err != nil {
		return fmt.Errorf("erro ao atualizar objeto: %v", err)
	}

	return nil
}

func (s *SustainablePestControlRepository) DeleteSustainablePestControl(id uint) error {

	entitySustainablePestControl, err := s.FindByIdSustainablePestControl(id)
	if err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	if err := s.db.Where("id = ?", entitySustainablePestControl.Id).Delete(&entitySustainablePestControl).Error; err != nil {
		return fmt.Errorf("erro ao deletar objeto: %v", err)
	}

	return nil
}
