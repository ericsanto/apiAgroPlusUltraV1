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

type PlantingRepositoryInterface interface {
	FindByParamPlanting(userID, farmID, batchID uint) (entities.PlantingEntity, error)
	CreatePlanting(entityPlanting entities.PlantingEntity) error
	FindByParamBatchNameOrIsActivePlanting(batchName string, active bool, userID, farmID uint) ([]responses.BatchPlantiesResponse, error)
	FindPlantingByID(batchID, farmID, userID, plantingID uint) (entities.PlantingEntity, error)
	UpdatePlanting(batchID, farmID, userID, plantingID uint, entityPlanting entities.PlantingEntity) error
	DeletePlanting(batchID, farmID, userID, plantingID uint) error
	FindAllPlanting(userID, farmID, batchID uint) ([]entities.PlantingEntity, error)
}

type PlantingRepository struct {
	db             interfaces.GORMRepositoryInterface
	farmRepository FarmRepositoryInterface
}

func NewPlantingRepository(db interfaces.GORMRepositoryInterface, farmRepository FarmRepositoryInterface) PlantingRepositoryInterface {
	return &PlantingRepository{db: db, farmRepository: farmRepository}
}

func (p *PlantingRepository) FindByParamPlanting(userID, farmID, batchID uint) (entities.PlantingEntity, error) {

	var responsePlanting entities.PlantingEntity

	query := `SELECT 
		batch_entities.name AS batch_name, 
		planting_entities.is_planting, 
		agriculture_culture_entities.name AS agriculture_culture_name,
		soil_type_entities.name AS soil_type_name,  
		planting_entities.start_date_planting,
		planting_entities.space_between_plants AS space_between_plants,
		planting_entities.space_between_rows AS space_between_rows,
		planting_entities.expected_production AS expected_production,
		irrigation_type_entities.name AS irrigation_type
	FROM planting_entities 
	INNER JOIN batch_entities ON batch_entities.id = planting_entities.batch_id 
	INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = planting_entities.agriculture_culture_id
	INNER JOIN soil_type_entities ON  soil_type_entities.id = agriculture_culture_entities.soil_type_id
	INNER JOIN irrigation_type_entities ON irrigation_type_entities.id = planting_entities.irrigation_type_id
	INNER JOIN farm_entities ON farm_entities.id = batch_entities.farm_id
	INNER JOIN user_models ON user_models.id = farm_entities.user_id  

	WHERE batch_entities.id = ? AND planting_entities.is_planting = true AND farm_entities.id = ? AND user_models.id = ?`

	if err := p.db.Raw(query, batchID, farmID, userID).Scan(&responsePlanting).Error; err != nil {
		return responsePlanting, myerror.NotFound(err)
	}

	return responsePlanting, nil
}

func (p *PlantingRepository) CreatePlanting(entityPlanting entities.PlantingEntity) error {

	if err := p.db.Create(&entityPlanting).Error; err != nil {
		return fmt.Errorf("erro ao tentar criar objeto")
	}

	return nil

}

func (p *PlantingRepository) FindByParamBatchNameOrIsActivePlanting(batchName string, active bool, userID, farmID uint) ([]responses.BatchPlantiesResponse, error) {

	_, err := p.farmRepository.FindByID(userID, farmID)

	if err != nil {
		return nil, err
	}

	//Sempre que for fazer uma busca com JOINS utilizando sql puro, é necessário criar alias das colunas com o nome igual ao do DTO de response
	query := `SELECT 
		batch_entities.name AS batch_name, 
		planting_entities.is_planting, 
		agriculture_culture_entities.name AS agriculture_culture_name,
		soil_type_entities.name AS soil_type_name,  
		planting_entities.start_date_planting,
		planting_entities.space_between_plants AS space_between_plants,
		planting_entities.space_between_rows AS space_between_rows,
		planting_entities.expected_production AS expected_production,
		irrigation_type_entities.name AS irrigation_type
	FROM planting_entities 
	INNER JOIN batch_entities ON batch_entities.id = planting_entities.batch_id 
	INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = planting_entities.agriculture_culture_id
	INNER JOIN soil_type_entities ON  soil_type_entities.id = agriculture_culture_entities.soil_type_id
	INNER JOIN irrigation_type_entities ON irrigation_type_entities.id = planting_entities.irrigation_type_id
	INNER JOIN farm_entities ON farm_entities.id = batch_entities.farm_id
	INNER JOIN user_models ON user_models.id = farm_entities.user_id  

	WHERE 
		( ? = '' OR REPLACE(batch_entities.name, ' ', '') ILIKE ?)
		 AND (planting_entities.is_planting = ?)
		 AND (farm_entities.id = ?)
		 AND (user_models.id = ?)`

	var entityListPlanting []responses.BatchPlantiesResponse

	batchNameFormated := fmt.Sprintf("%%%s%%", batchName)

	if err := p.db.Raw(query, batchName, batchNameFormated, active, farmID, userID).Scan(&entityListPlanting).Error; err != nil {
		return entityListPlanting, fmt.Errorf("erro ao buscar dados: %w", err)
	}

	return entityListPlanting, nil
}

func (p *PlantingRepository) FindAllPlanting(userID, farmID, batchID uint) ([]entities.PlantingEntity, error) {

	var entitiesPlanting []entities.PlantingEntity

	if err := p.db.Joins("JOIN batch_entities ON batch_entities.id = planting_entities.batch_id").
		Joins("JOIN farm_entities ON farm_entities.id = batch_entities.farm_id").
		Joins("JOIN user_models ON user_models.id = farm_entities.user_id").
		Where("batch_entities.id = ? AND farm_entities.id = ? AND user_models.id = ?", batchID, farmID, userID).Find(&entitiesPlanting).Error; err != nil {
		return entitiesPlanting, fmt.Errorf("erro ao buscar todas as plantações: %w", err)
	}

	return entitiesPlanting, nil
}

func (p *PlantingRepository) FindPlantingByID(batchID, farmID, userID, plantingID uint) (entities.PlantingEntity, error) {

	var entityPlanting entities.PlantingEntity

	if err := p.db.Model(&entities.PlantingEntity{}).
		Joins("JOIN batch_entities ON batch_entities.id = planting_entities.batch_id").
		Joins("JOIN farm_entities ON farm_entities.id = batch_entities.farm_id").
		Joins("JOIN user_models ON user_models.id = farm_entities.user_id").
		Where("batch_entities.id = ? AND farm_entities.id = ? AND user_models.id = ? AND planting_entities.id = ?", batchID, farmID, userID, plantingID).First(&entityPlanting).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entityPlanting, fmt.Errorf("não existe plantação com o id %d. %w", plantingID, err)
		}

		return entityPlanting, fmt.Errorf("erro ao buscar plantações")
	}

	return entityPlanting, nil

}

func (p *PlantingRepository) UpdatePlanting(batchID, farmID, userID, plantingID uint, entityPlanting entities.PlantingEntity) error {

	if _, err := p.FindPlantingByID(batchID, farmID, userID, plantingID); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	//Pq criei um novo map e nao passei o entityPlanting que está chegando como parâmetro?
	//O gorm não atualiza struct com valores que são considerados zero value
	//No caso de campos booleanos, o valor false é um zero value, então o GORM ignora essa atualização
	//Nessa caso, foi preciso criar um map para forçar a ser atualizado como false
	updateEntity := map[string]interface{}{
		"is_planting":            entityPlanting.IsPlanting,
		"agriculture_culture_id": entityPlanting.AgricultureCultureID,
	}

	if err := p.db.Model(&entities.PlantingEntity{}).Where("id = ?", plantingID).Updates(&updateEntity).Error; err != nil {
		return fmt.Errorf("erro ao atualilzar plantação")
	}

	return nil
}

func (p *PlantingRepository) DeletePlanting(batchID, farmID, userID, plantingID uint) error {

	planting, err := p.FindPlantingByID(batchID, farmID, userID, plantingID)
	if err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	if err := p.db.Model(&entities.PlantingEntity{}).
		Where("id = ?", plantingID).Delete(&planting).Error; err != nil {
		return fmt.Errorf("erro ao tentar deletar plantação: %w", err)
	}

	return nil
}
