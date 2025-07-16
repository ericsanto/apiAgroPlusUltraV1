package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type PerformancePlantingRepositoryInterface interface {
	CreatePerformancePlanting(batchID, farmID, userID, plantingID uint, entityPerformanceCulutre entities.PerformancePlantingEntity) error
	FindAll(batchID, farmID, userID uint) ([]responses.DbResultPerformancePlanting, error)
	FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID(batchID, farmID, userID, plantingID, performanceID uint) (*responses.DbResultPerformancePlanting, error)
	FindPerformancePlantingByID(batchID, farmID, userID, plantingID, performanceID uint) (*entities.PerformancePlantingEntity, error)
	UpdatePerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint, entityPerformancePlanting entities.PerformancePlantingEntity) error
	DeletePerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint) error
}

type PerformancePlantingRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewPerformanceCultureRepository(db interfaces.GORMRepositoryInterface) PerformancePlantingRepositoryInterface {
	return &PerformancePlantingRepository{db: db}
}

func (p *PerformancePlantingRepository) CreatePerformancePlanting(batchID, farmID, userID, plantingID uint, entityPerformanceCulutre entities.PerformancePlantingEntity) error {

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

func (p *PerformancePlantingRepository) FindAll(batchID, farmID, userID uint) ([]responses.DbResultPerformancePlanting, error) {

	var dbResult []responses.DbResultPerformancePlanting

	query := `SELECT
	performance_planting_entities.id AS id,
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
	INNER JOIN farm_entities ON farm_entities.id = batch_entities.farm_id
	INNER JOIN user_models ON user_models.id = farm_entities.user_id
	INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id =  planting_entities.agriculture_culture_id
	
	WHERE  farm_entities.id = ? AND batch_entities.id = ? AND user_models.id = ?`

	err := p.db.Raw(query, farmID, batchID, userID).Scan(&dbResult)
	if err.Error != nil {
		return nil, fmt.Errorf("erro ao buscar performance de plantação")
	}

	return dbResult, nil

}

func (p *PerformancePlantingRepository) FindPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByID(batchID, farmID, userID, plantingID, performanceID uint) (*responses.DbResultPerformancePlanting, error) {

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
		INNER JOIN farm_entities ON farm_entities.id = batch_entities.farm_id
		INNER JOIN user_models ON user_models.id = farm_entities.user_id
		INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id =  planting_entities.agriculture_culture_id
	
		WHERE  farm_entities.id = ? AND batch_entities.id = ? AND user_models.id = ? AND planting_entities.id = ? AND
		performance_planting_entities.id = ?`

	if _, err := p.FindPerformancePlantingByID(batchID, farmID, userID, plantingID, performanceID); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	err := p.db.Raw(query, farmID, batchID, userID, plantingID, performanceID).Scan(&dBResultPerformancePlanting)
	if err.Error != nil {
		return nil, nil
	}

	return &dBResultPerformancePlanting, nil
}

func (p *PerformancePlantingRepository) FindPerformancePlantingByID(batchID, farmID, userID, plantingID, performanceID uint) (*entities.PerformancePlantingEntity, error) {

	var entityPerformancePlanting entities.PerformancePlantingEntity

	if err := p.db.Model(entities.PerformancePlantingEntity{}).Joins("JOIN planting_entities ON planting_entities.id = performance_planting_entities.planting_id").
		Joins("JOIN batch_entities ON batch_entities.id = planting_entities.batch_id").
		Joins("JOIN farm_entities ON farm_entities.id = batch_entities.farm_id").
		Joins("JOIN user_models ON user_models.id = farm_entities.user_id").
		Where("planting_entities.id = ? AND batch_entities.id = ? AND farm_entities.id = ? AND user_models.id = ? AND performance_planting_entities.id = ?", plantingID, batchID, farmID, userID, performanceID).
		First(&entityPerformancePlanting).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, fmt.Errorf("performance de plantação com id %d %w", performanceID, myerror.ErrNotFound)

		default:
			return nil, fmt.Errorf("erro buscar performance de plantação")
		}
	}

	return &entityPerformancePlanting, nil
}

func (p *PerformancePlantingRepository) UpdatePerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint, entityPerformancePlanting entities.PerformancePlantingEntity) error {

	if _, err := p.FindPerformancePlantingByID(batchID, farmID, userID, plantingID, performanceID); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := p.db.Model(&entities.PerformancePlantingEntity{}).Where("id = ?", performanceID).Updates(&entityPerformancePlanting).Error; err != nil {
		switch {
		case myerror.IsUniqueConstraintViolated(err):
			return fmt.Errorf("performance culture com planting_id  %d %w", performanceID, myerror.ErrDuplicateKey)

		case myerror.IsViolatedForeingKeyConstraint(err):
			return fmt.Errorf("plantação com id %d %w", performanceID, myerror.ErrViolatedForeingKey)

		default:
			return fmt.Errorf("erro ao atualizar performance culture")
		}

	}

	return nil
}

func (p *PerformancePlantingRepository) DeletePerformancePlanting(batchID, farmID, userID, plantingID, performanceID uint) error {

	entityPerformancePlanting, err := p.FindPerformancePlantingByID(batchID, farmID, userID, plantingID, performanceID)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := p.db.Where("id = ?", performanceID).Delete(&entityPerformancePlanting).Error; err != nil {
		return fmt.Errorf("erro ao deletar performance da cultura")
	}

	return nil
}
