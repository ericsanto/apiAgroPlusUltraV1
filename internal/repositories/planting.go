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
	FindByParamPlanting(batchID uint) (entities.PlantingEntity, error)
	CreatePlanting(entityPlanting entities.PlantingEntity) error
	FindByParamBatchNameOrIsActivePlanting(batchName string, active bool) ([]responses.BatchPlantiesResponse, error)
	FindPlantingByID(id uint) (entities.PlantingEntity, error)
	UpdatePlanting(id uint, entityPlanting entities.PlantingEntity) error
	DeletePlanting(id uint) error
	FindAllPlanting() ([]entities.PlantingEntity, error)
}

type PlantingRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewPlantingRepository(db interfaces.GORMRepositoryInterface) PlantingRepositoryInterface {
	return &PlantingRepository{db: db}
}

func (p *PlantingRepository) FindByParamPlanting(batchID uint) (entities.PlantingEntity, error) {

	var responsePlanting entities.PlantingEntity

	query := `SELECT * FROM planting_entities WHERE planting_entities.batch_id = ? AND  planting_entities.is_planting = true`

	if err := p.db.Raw(query, batchID).Scan(&responsePlanting).Error; err != nil {
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

func (p *PlantingRepository) FindByParamBatchNameOrIsActivePlanting(batchName string, active bool) ([]responses.BatchPlantiesResponse, error) {

	//Sempre que for fazer uma busca com JOINS utilizando sql puro, é necessário criar alias das colunas com o nome igual ao do DTO de response
	query := `SELECT 
		batch_entities.name AS batch_name, 
		planting_entities.is_planting, 
		agriculture_culture_entities.name AS agriculture_culture_name,
		soil_type_entities.name AS soil_type_name,  
		planting_entities.start_date_planting,
		planting_entities.space_between_plants AS space_between_plants,
		planting_entities.space_between_rows AS space_between_rows,
		irrigation_type_entities.name AS irrigation_type
	FROM planting_entities 
	INNER JOIN batch_entities ON batch_entities.id = planting_entities.batch_id 
	INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = planting_entities.agriculture_culture_id
	INNER JOIN soil_type_entities ON  soil_type_entities.id = agriculture_culture_entities.soil_type_id
	INNER JOIN irrigation_type_entities ON irrigation_type_entities.id = planting_entities.irrigation_type_id
	WHERE 
		( ? = '' OR REPLACE(batch_entities.name, ' ', '') ILIKE ?)
		 AND (planting_entities.is_planting = ?)`

	var entityListPlanting []responses.BatchPlantiesResponse

	batchNameFormated := fmt.Sprintf("%%%s%%", batchName)

	if err := p.db.Raw(query, batchName, batchNameFormated, active).Scan(&entityListPlanting).Error; err != nil {
		return entityListPlanting, fmt.Errorf("erro ao buscar dados: %w", err)
	}

	return entityListPlanting, nil
}

func (p *PlantingRepository) FindAllPlanting() ([]entities.PlantingEntity, error) {

	var entitiesPlanting []entities.PlantingEntity

	if err := p.db.Find(&entitiesPlanting).Error; err != nil {
		return entitiesPlanting, fmt.Errorf("erro ao buscar todas as plantações: %w", err)
	}

	return entitiesPlanting, nil
}

func (p *PlantingRepository) FindPlantingByID(id uint) (entities.PlantingEntity, error) {

	var entityPlanting entities.PlantingEntity

	if err := p.db.First(&entityPlanting, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entityPlanting, fmt.Errorf("não existe plantação com o id %d. %w", id, err)
		}

		return entityPlanting, fmt.Errorf("erro ao buscar plantações")
	}

	return entityPlanting, nil

}

func (p *PlantingRepository) UpdatePlanting(id uint, entityPlanting entities.PlantingEntity) error {

	if _, err := p.FindPlantingByID(id); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	//Pq criei um novo map e nao passei o entityPlanting que está chegando como parâmetro?
	//O gorm não atualiza struct com valores que são considerados zero value
	//No caso de campos booleanos, o valor false é um zero value, então o GORM ignora essa atualização
	//Nessa caso, foi preciso criar um map para forçar a ser atualizado como false
	updateEntity := map[string]interface{}{
		"is_planting":            entityPlanting.IsPlanting,
		"batch_id":               entityPlanting.BatchID,
		"agriculture_culture_id": entityPlanting.AgricultureCultureID,
	}

	if err := p.db.Model(&entities.PlantingEntity{}).Where("id = ?", id).Updates(&updateEntity).Error; err != nil {
		return fmt.Errorf("erro ao atualilzar plantação")
	}

	return nil
}

func (p *PlantingRepository) DeletePlanting(id uint) error {

	planting, err := p.FindPlantingByID(id)
	if err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	if err := p.db.Where("id = ?", id).Delete(&planting).Error; err != nil {
		return fmt.Errorf("erro ao tentar deletar plantação: %w", err)
	}

	return nil
}
