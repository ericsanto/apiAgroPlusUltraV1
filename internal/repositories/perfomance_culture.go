package repositories

import (
	"errors"
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"gorm.io/gorm"
)

type PerformancePlantingRepository struct {
	db *gorm.DB
}

func NewPerformanceCultureRepository(db *gorm.DB) *PerformancePlantingRepository {
	return &PerformancePlantingRepository{db: db}
}

func (p *PerformancePlantingRepository) CreatePerformancePlanting(entityPerformanceCulutre entities.PerformancePlantingEntity) error {

	if err := p.db.Create(&entityPerformanceCulutre).Error; err != nil {
		switch {
		case myerror.IsViolatedForeingKeyConstraint(err):
			return fmt.Errorf("%w %s %d", myerror.ErrViolatedForeingKey, "planting com id", entityPerformanceCulutre.PlantingID)

		case myerror.IsUniqueConstraintViolated(err):
			return fmt.Errorf("%w %s %d", myerror.ErrDuplicateKey, "performance da cultura com planting_id ", entityPerformanceCulutre.PlantingID)
		default:
			return fmt.Errorf("erro ao cadastrar performance da cultura")
		}

	}

	return nil
}

func (p *PerformancePlantingRepository) FindAll() ([]responses.DbResultPerformancePlanting, error) {

	var dbResult []responses.DbResultPerformancePlanting

	query := `SELECT
	performance_planting_entities.id,
	batch_entities.name AS batch_name, 
	agriculture_culture_entities.name AS agriculture_culture_name,
	planting_entities.start_date_planting AS start_date_planting,
	performance_planting_entities.production_obtained AS production_obtained, 
	concat(performance_planting_entities.production_obtained, performance_planting_entities.unit_production_obtained) AS production_obtained_formated,
	performance_planting_entities.harvested_area AS harvested_area, 
	concat(performance_planting_entities.harvested_area, performance_planting_entities.unit_harvested_area) AS harvested_area_formated,
	performance_planting_entities.harvested_date AS harvested_date

	FROM performance_planting_entities

	INNER JOIN planting_entities ON planting_entities.id = performance_planting_entities.planting_id
	INNER JOIN batch_entities ON batch_entities.id = planting_entities.batch_id
	INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id =  planting_entities.agriculture_culture_id`

	err := p.db.Raw(query).Scan(&dbResult)
	if err.Error != nil {
		return nil, fmt.Errorf("erro ao buscar performance de plantação")
	}

	return dbResult, nil

}

func (p *PerformancePlantingRepository) FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID(id uint) (*responses.DbResultPerformancePlanting, error) {

	var dBResultPerformancePlanting responses.DbResultPerformancePlanting

	query := `SELECT
		performance_planting_entities.id,
		batch_entities.name AS batch_name, 
		agriculture_culture_entities.name AS agriculture_culture_name,
		planting_entities.start_date_planting AS start_date_planting,
		performance_planting_entities.production_obtained AS production_obtained, 
		concat(performance_planting_entities.production_obtained, performance_planting_entities.unit_production_obtained) AS production_obtained_formated,
		performance_planting_entities.harvested_area AS harvested_area, 
		concat(performance_planting_entities.harvested_area, performance_planting_entities.unit_harvested_area) AS harvested_area_formated,
		performance_planting_entities.harvested_date AS harvested_date

		FROM performance_planting_entities

		INNER JOIN planting_entities ON planting_entities.id = performance_planting_entities.planting_id
		INNER JOIN batch_entities ON batch_entities.id = planting_entities.batch_id
		INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id =  planting_entities.agriculture_culture_id
		WHERE performance_planting_entities.id = ?`

	if _, err := p.FindPerformancePlantingByID(id); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	err := p.db.Raw(query, id).Scan(&dBResultPerformancePlanting)
	if err.Error != nil {
		return nil, nil
	}

	return &dBResultPerformancePlanting, nil
}

func (p *PerformancePlantingRepository) FindPerformancePlantingByID(id uint) (*entities.PerformancePlantingEntity, error) {

	var entityPerformancePlanting entities.PerformancePlantingEntity

	if err := p.db.First(&entityPerformancePlanting, id).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, fmt.Errorf("performance de plantação com id %d %w", id, myerror.ErrNotFound)

		default:
			return nil, fmt.Errorf("erro buscar performance de plantação")
		}
	}

	return &entityPerformancePlanting, nil
}

func (p *PerformancePlantingRepository) UpdatePerformancePlanting(id uint, entityPerformancePlanting entities.PerformancePlantingEntity) error {

	if _, err := p.FindPerformancePlantingByID(id); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := p.db.Model(&entities.PerformancePlantingEntity{}).Where("id = ?", id).Updates(&entityPerformancePlanting).Error; err != nil {
		switch {
		case myerror.IsUniqueConstraintViolated(err):
			return fmt.Errorf("performance culture com planting_id  %d %w", id, myerror.ErrDuplicateKey)

		case myerror.IsViolatedForeingKeyConstraint(err):
			return fmt.Errorf("plantação com id %d %w", id, myerror.ErrViolatedForeingKey)

		default:
			return fmt.Errorf("erro ao atualizar performance culture")
		}

	}

	return nil
}

func (p *PerformancePlantingRepository) DeletePerformancePlanting(id uint) error {

	entityPerformancePlanting, err := p.FindPerformancePlantingByID(id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := p.db.Where("id = ?", id).Delete(&entityPerformancePlanting).Error; err != nil {
		return fmt.Errorf("erro ao deletar performance da cultura")
	}

	return nil
}
