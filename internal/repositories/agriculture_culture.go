package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

type AgricultureCultureInterface interface {
	FindAllAgricultureCulture() ([]entities.AgricultureCultureEntity, error)
	FindByIdAgricultureCulture(id uint) (*entities.AgricultureCultureEntity, error)
	CreateAgricultureCulture(agriculutreCulture entities.AgricultureCultureEntity) error
	UpdateAgricultureCulture(id uint, agricultureCultureEntity entities.AgricultureCultureEntity) error
	DeleteAgricultureCulture(id uint) error
}

type AgricultureCultureRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewAgricultureCultureRepository(db interfaces.GORMRepositoryInterface) AgricultureCultureInterface {

	return &AgricultureCultureRepository{db: db}
}

func (a *AgricultureCultureRepository) FindAllAgricultureCulture() ([]entities.AgricultureCultureEntity, error) {

	var agricultureCultures []entities.AgricultureCultureEntity
	if err := a.db.Find(&agricultureCultures).Error; err != nil {
		return agricultureCultures, fmt.Errorf("erro ao buscar dados %v", err)

	}

	return agricultureCultures, nil

}

func (a *AgricultureCultureRepository) FindByIdAgricultureCulture(id uint) (*entities.AgricultureCultureEntity, error) {

	var agricultureCulture entities.AgricultureCultureEntity

	if err := a.db.First(&agricultureCulture, id).Error; err != nil {
		return &agricultureCulture, fmt.Errorf("erro ao buscar cultura agrícola. Id não existe no banco de dados")

	}

	return &agricultureCulture, nil
}

func (a *AgricultureCultureRepository) CreateAgricultureCulture(agriculutreCulture entities.AgricultureCultureEntity) error {

	if err := a.db.Create(&agriculutreCulture).Error; err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return fmt.Errorf("já existe cultura agrícola com essa variedade")
		}
		return fmt.Errorf("não foi possivel salvar tipo de cultura no banco de dados %v", err)
	}

	return nil
}

func (r *AgricultureCultureRepository) UpdateAgricultureCulture(id uint, agricultureCultureEntity entities.AgricultureCultureEntity) error {

	_, err := r.FindByIdAgricultureCulture(id)
	if err != nil {
		return fmt.Errorf("erro ao buscar cultura agricola: %v", err)
	}

	if err := r.db.Model(&entities.AgricultureCultureEntity{}).Where("id = ?", id).Updates(&agricultureCultureEntity).Error; err != nil {
		return fmt.Errorf("erro ao atualizar cultura agrícola: %v", err)
	}

	return nil
}

func (r *AgricultureCultureRepository) DeleteAgricultureCulture(id uint) error {

	agricultureCulture, err := r.FindByIdAgricultureCulture(id)
	if err != nil {
		return fmt.Errorf("erro ao buscar cultura agricola: %v", err)
	}

	if err := r.db.Where("id = ?", id).Delete(&agricultureCulture).Error; err != nil {
		return fmt.Errorf("erro ao deletar cultura agricola: %v", err)
	}

	return nil
}
