package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"gorm.io/gorm"
)

type AgricultureCulturePestMethodRepository struct {
	db *gorm.DB
}

func NewAgricultureCulturePestMethodRepository(db *gorm.DB) *AgricultureCulturePestMethodRepository {
	return &AgricultureCulturePestMethodRepository{db: db}
}

func (a *AgricultureCulturePestMethodRepository) CreateAgricultureCulturePestMethod(entityAgricultureCulturePestMethod entities.AgricultureCulturePestMethodEntity) error {

	if err := a.db.Create(&entityAgricultureCulturePestMethod).Error; err != nil {
		return fmt.Errorf("erro ao criar objeto")
	}

	return nil
}

func (a *AgricultureCulturePestMethodRepository) FindAllAgricultureCulturePestMethod() ([]responses.AgricultureCulturePestMethodResponse, error) {

	var responseAgricultureCulturePestMethod []responses.AgricultureCulturePestMethodResponse

	query := `SELECT 
    agriculture_culture_entities.name AS agriculture_culture_name,
    pest_entities.name AS pest_name,
    sustainable_pest_control_entities.name AS sustainable_pest_control_method,
    agriculture_culture_pest_method_entities.description
FROM agriculture_culture_pest_method_entities
INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = agriculture_culture_pest_method_entities.agriculture_culture_id
INNER JOIN pest_entities ON pest_entities.id = agriculture_culture_pest_method_entities.pest_id
INNER JOIN sustainable_pest_control_entities ON sustainable_pest_control_entities.id = agriculture_culture_pest_method_entities.sustainable_pest_control_id`

	if err := a.db.Raw(query).Scan(&responseAgricultureCulturePestMethod).Error; err != nil {
		return responseAgricultureCulturePestMethod, fmt.Errorf("erro ao buscar todos os metodos de controle de praga")
	}

	return responseAgricultureCulturePestMethod, nil
}

func (a *AgricultureCulturePestMethodRepository) FindByQueryParamAgricultureCulturePestMethod(cultureName, pestName, methodName interface{}) ([]responses.AgricultureCulturePestMethodResponse, error) {

	var responseAgricultureCulturePestMethod []responses.AgricultureCulturePestMethodResponse

	query := `SELECT 
    agriculture_culture_entities.name AS agriculture_culture_name,
    pest_entities.name AS pest_name,
    sustainable_pest_control_entities.name AS sustainable_pest_control_method,
    agriculture_culture_pest_method_entities.description
FROM agriculture_culture_pest_method_entities
INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = agriculture_culture_pest_method_entities.agriculture_culture_id
INNER JOIN pest_entities ON pest_entities.id = agriculture_culture_pest_method_entities.pest_id
INNER JOIN sustainable_pest_control_entities ON sustainable_pest_control_entities.id = agriculture_culture_pest_method_entities.sustainable_pest_control_id 
WHERE 
	(? = '' OR agriculture_culture_entities.name = ?) AND
	(? = '' OR pest_entities.name = ?) AND
	(? = '' OR sustainable_pest_control_entities.name = ?)`

	//Não posso utilizar alias no where

	if err := a.db.Raw(query, cultureName, cultureName, pestName, pestName, methodName, methodName).Scan(&responseAgricultureCulturePestMethod).Error; err != nil {
		return responseAgricultureCulturePestMethod, fmt.Errorf("erro ao buscar no banco de dados")
	}

	return responseAgricultureCulturePestMethod, nil

}
