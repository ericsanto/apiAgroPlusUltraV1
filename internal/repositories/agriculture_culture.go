package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"gorm.io/gorm"
)

type AgricultureCultureInterface interface {
	FindAllAgricultureCulture() ([]entities.AgricultureCultureEntity, error)
	FindByIdAgricultureCulture(id uint) (entities.AgricultureCultureEntity, error)
	CreateAgricultureCulture(agricultureCulture *entities.AgricultureCultureEntity)
	UpdateAgricultureCulture(id uint, agricultureCulture entities.AgricultureCultureEntity) error
	DeleteAgricultureCulture(id uint)
}

type AgricultureCultureRepository struct {
	db *gorm.DB
}

func NewAgricultureCultureRepository(db *gorm.DB) *AgricultureCultureRepository {

	return &AgricultureCultureRepository{db: db}
}

func (a *AgricultureCultureRepository) FindAllAgricultureCulture() ([]entities.AgricultureCultureEntity, error) {

	var agricultureCultures []entities.AgricultureCultureEntity
	if err := a.db.Find(&agricultureCultures).Error; err != nil {
		return agricultureCultures, fmt.Errorf("erro ao buscar dados %v", err)

	}

	return agricultureCultures, nil

}

func (a *AgricultureCultureRepository) FindByIdAgricultureCulture(id uint) (entities.AgricultureCultureEntity, error) {

	var agricultureCulture entities.AgricultureCultureEntity

	if err := a.db.First(&agricultureCulture, id).Error; err != nil {
		return agricultureCulture, fmt.Errorf("erro ao buscar cultura agrícola. Id não existe no banco de dados")

	}

	return agricultureCulture, nil
}

func (a *AgricultureCultureRepository) CreateAgricultureCulture(agriculutreCulture *entities.AgricultureCultureEntity) error {

	if err := a.db.Create(&agriculutreCulture).Error; err != nil {
		return fmt.Errorf("não foi possivel salvar tipo de cultura no banco de dados %v", err)
	}

	return nil
}

func (r *AgricultureCultureRepository) UpdateAgricultureCulture(id uint, agricultureCultureEntity entities.AgricultureCultureEntity) error {

	agricultureCulture, err := r.FindByIdAgricultureCulture(id)
	if err != nil {
		return fmt.Errorf("erro ao buscar cultura agricola: %v", err)
	}

	if err := r.db.Model(&entities.AgricultureCultureEntity{}).Where("id = ?", agricultureCulture.Id).Updates(agricultureCultureEntity).Error; err != nil {
		return fmt.Errorf("erro ao atualizar cultura agrícola: %v", err)
	}

	return nil
}

func (r *AgricultureCultureRepository) DeleteAgricultureCulture(id uint) error {

	agricultureCulture, err := r.FindByIdAgricultureCulture(id)
	if err != nil {
		return fmt.Errorf("erro ao buscar cultura agricola: %v", err)
	}

	if err := r.db.Where("id = ?", agricultureCulture.Id).Delete(&entities.AgricultureCultureEntity{}).Error; err != nil {
		return fmt.Errorf("erro ao deletar cultura agricola: %v", err)
	}

	return nil
}
