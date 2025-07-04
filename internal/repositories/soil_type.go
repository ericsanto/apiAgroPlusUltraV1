package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

type SoilTypeInterface interface {
	FindAllSoilType() ([]entities.SoilTypeEntity, error)
	FindByIdSoilType(id uint) (*entities.SoilTypeEntity, error)
	CreateSoilType(soilTypeModel entities.SoilTypeEntity) error
	UpdateSoilType(id uint, soilTypeModel entities.SoilTypeEntity) error
	DeleteSoilType(id uint) error
}

type SoilTypeRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewSoilRepository(db interfaces.GORMRepositoryInterface) SoilTypeInterface {
	return &SoilTypeRepository{db: db}
}

func (r *SoilTypeRepository) FindAllSoilType() ([]entities.SoilTypeEntity, error) {

	var soilTypes []entities.SoilTypeEntity

	err := r.db.Find(&soilTypes).Error

	return soilTypes, err
}

func (r *SoilTypeRepository) FindByIdSoilType(id uint) (*entities.SoilTypeEntity, error) {

	var soilType entities.SoilTypeEntity

	err := r.db.First(&soilType, id).Error

	return &soilType, err
}

func (r *SoilTypeRepository) CreateSoilType(soilTypeModel entities.SoilTypeEntity) error {

	result := r.db.Create(&soilTypeModel)

	if result.Error != nil {
		return fmt.Errorf("erro ao criar no banco de dados %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nenhum registro foi criado")
	}

	return nil
}

func (r *SoilTypeRepository) UpdateSoilType(id uint, soilTypeModel entities.SoilTypeEntity) error {

	_, err := r.FindByIdSoilType(id)
	if err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	result := r.db.Model(&entities.SoilTypeEntity{}).Where("id = ?", id).Updates(&soilTypeModel)

	if result.Error != nil {
		return fmt.Errorf("erro ao atualizaer dados %w", result.Error)
	}

	return nil
}

func (r *SoilTypeRepository) DeleteSoilType(id uint) error {

	soilTypeModelExist, err := r.FindByIdSoilType(id)
	if err != nil {
		return fmt.Errorf("erro ao deletar tipo de solo. Id não existe")
	}

	result := r.db.Where("id = ?", id).Delete(&soilTypeModelExist)

	if result.Error != nil {
		return fmt.Errorf("não foi possível deletar tipo de solo %w", err)
	}

	return nil
}
