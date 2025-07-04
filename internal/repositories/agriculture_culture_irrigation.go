package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

var queryVerifyIdExists string = `SELECT EXISTS(SELECT 1 FROM agriculture_culture_irrigations WHERE agriculture_culture_irrigations.agriculture_culture_id = ? AND agriculture_culture_irrigations.irrigation_recomended_id = ?)`

type AgricultureCultureIrrigationRepositoryInterface interface {
	FindByIdAgricultureCultureIrrigation(cultureId uint) ([]responses.AgricultureCultureIrrigationResponse, error)
	CreateAgricultureCultureIrrigation(entityAgricultureCultureIrrigation entities.AgricultureCultureIrrigation) error
	UpdateAgricultureCultureIrrigation(cultureId, irrigationId uint, entityAgricultureCultureIrrigation entities.AgricultureCultureIrrigation) error
	DeleteAgricultureCulturueIrrigation(cultureId, irrigationId uint) error
}

type AgricultureCultureIrrigationRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewAgricultureCultureIrrigationRepository(db interfaces.GORMRepositoryInterface) AgricultureCultureIrrigationRepositoryInterface {
	return &AgricultureCultureIrrigationRepository{db: db}
}

// Ele vai buscar através do id da cultura. A partir dai vem a irrigação recomendada
func (a *AgricultureCultureIrrigationRepository) FindByIdAgricultureCultureIrrigation(cultureId uint) ([]responses.AgricultureCultureIrrigationResponse, error) {

	agricultureCultureRepository := NewAgricultureCultureRepository(a.db)

	if _, err := agricultureCultureRepository.FindByIdAgricultureCulture(cultureId); err != nil {
		return nil, fmt.Errorf("erro: %v", err)
	}

	query := `SELECT agriculture_culture_entities.name, irrigation_recomended_entities.phenological_phase, irrigation_recomended_entities.phase_duration_days,
	irrigation_recomended_entities.irrigation_max, irrigation_recomended_entities.irrigation_min, irrigation_recomended_entities.unit
	FROM agriculture_culture_irrigations
	INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = agriculture_culture_irrigations.agriculture_culture_id
	INNER JOIN irrigation_recomended_entities ON irrigation_recomended_entities.id = agriculture_culture_irrigations.irrigation_recomended_id
	WHERE agriculture_culture_entities.id = ?`

	var reponseAgricultureCultureIrrigation []responses.AgricultureCultureIrrigationResponse

	if err := a.db.Raw(query, cultureId).Scan(&reponseAgricultureCultureIrrigation).Error; err != nil {
		return reponseAgricultureCultureIrrigation, fmt.Errorf("erro ao buscar dados: %v", err)
	}

	return reponseAgricultureCultureIrrigation, nil
}

func (a *AgricultureCultureIrrigationRepository) CreateAgricultureCultureIrrigation(entityAgricultureCultureIrrigation entities.AgricultureCultureIrrigation) error {

	var exists bool

	if err := a.db.Raw(queryVerifyIdExists, entityAgricultureCultureIrrigation.AgricultureCultureId, entityAgricultureCultureIrrigation.IrrigationRecomendedId).Scan(&exists).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("id não existe: %v", err)
		}

		return fmt.Errorf("erro ao verificar relação: %v", err)
	}

	if exists {
		return fmt.Errorf("erro ao tentar cadastrar. já existe objeto com essa relação de id")
	}

	if err := a.db.Create(&entityAgricultureCultureIrrigation).Error; err != nil {

		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return fmt.Errorf("erro ao relacionar id. id inexistente")
		}

		return fmt.Errorf("erro ao tentar criar objeto %v", err)
	}

	return nil
}

func (a *AgricultureCultureIrrigationRepository) UpdateAgricultureCultureIrrigation(cultureId, irrigationId uint, entityAgricultureCultureIrrigation entities.AgricultureCultureIrrigation) error {

	var firstAgricultureCultureIrrigation entities.AgricultureCultureIrrigation

	if err := a.db.First(&firstAgricultureCultureIrrigation, cultureId, irrigationId).Error; err != nil {
		return fmt.Errorf("não existe objeto com id fornecido: %v", err)
	}

	if err := a.db.Model(entities.AgricultureCultureIrrigation{}).Where("agriculture_culture_irrigations.agriculture_culture_id = ? AND agriculture_culture_irrigations.irrigation_recomended_id = ?",
		firstAgricultureCultureIrrigation.AgricultureCultureId, firstAgricultureCultureIrrigation.IrrigationRecomendedId).Updates(&entityAgricultureCultureIrrigation).Error; err != nil {
		return fmt.Errorf("erro ao atulizar objeto %v", err)
	}

	return nil
}

func (a *AgricultureCultureIrrigationRepository) DeleteAgricultureCulturueIrrigation(cultureId, irrigationId uint) error {

	var entityAgricultureCultureIrrigation entities.AgricultureCultureIrrigation

	if err := a.db.First(&entityAgricultureCultureIrrigation, cultureId, irrigationId).Error; err != nil {
		return fmt.Errorf("não existe objeto com id fornecido: %v", err)
	}

	if err := a.db.Where("agriculture_culture_irrigations.agriculture_culture_id = ? AND agriculture_culture_irrigations.irrigation_recomended_id = ?",
		entityAgricultureCultureIrrigation.AgricultureCultureId, entityAgricultureCultureIrrigation.IrrigationRecomendedId).Delete(&entityAgricultureCultureIrrigation).Error; err != nil {
		return fmt.Errorf("erro ao deletar objeto: %v", err)
	}

	return nil

}
