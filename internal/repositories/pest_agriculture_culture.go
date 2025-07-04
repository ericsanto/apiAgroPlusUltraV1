package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

var queryVerifyIdExist string = `SELECT EXISTS(SELECT 1 FROM pest_agriculture_cultures WHERE pest_agriculture_cultures.pest_id = ? AND pest_agriculture_cultures.agriculture_culture_id = ?)`

type PestAgricultureCultureRepositoryInterface interface {
	FindAllPestAgricultureCulture() ([]responses.PestAgricultureCultureResponse, error)
	FindByIdPestAgricultureCulture(pestId, cultureId uint) (*responses.PestAgricultureCultureResponse, error)
	CreatePestAgricultureCulture(entityPestAgricultureCulture entities.PestAgricultureCulture) error
	UpdatePestAgricultureCulture(pestId, cultureId uint, entityPestAgricultureCulture entities.PestAgricultureCulture) error
	DeletePestAgricultureCulture(pestId, cultureId uint) error
}

type PestAgricultureCultureRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewPestAgricultureCultureRepository(db interfaces.GORMRepositoryInterface) PestAgricultureCultureRepositoryInterface {
	return &PestAgricultureCultureRepository{db: db}
}

func (p *PestAgricultureCultureRepository) FindAllPestAgricultureCulture() ([]responses.PestAgricultureCultureResponse, error) {

	var responsePestAgricultureCulutre []responses.PestAgricultureCultureResponse

	query := `SELECT pest_entities.name AS pest_name, type_pest_entities.name AS type_pest_name,
  agriculture_culture_entities.name AS agriculture_culture_name, description, image_url

  FROM pest_agriculture_cultures

  INNER JOIN pest_entities ON pest_entities.id = pest_agriculture_cultures.pest_id

  INNER JOIN type_pest_entities ON type_pest_entities.id = pest_entities.type_pest_id

  INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = pest_agriculture_cultures.agriculture_culture_id`

	if err := p.db.Raw(query).Scan(&responsePestAgricultureCulutre).Error; err != nil {
		return responsePestAgricultureCulutre, fmt.Errorf("erro no repositório ao fazer consulta %v", err)
	}

	return responsePestAgricultureCulutre, nil
}

func (p *PestAgricultureCultureRepository) FindByIdPestAgricultureCulture(pestId, cultureId uint) (*responses.PestAgricultureCultureResponse, error) {
	var response responses.PestAgricultureCultureResponse

	query := `SELECT 
            pest_entities.name AS pest_name, 
            type_pest_entities.name AS type_pest_name,
            agriculture_culture_entities.name AS agriculture_culture_name, 
            pest_agriculture_cultures.description, 
            pest_agriculture_cultures.image_url
        FROM pest_agriculture_cultures
        INNER JOIN pest_entities ON pest_entities.id = pest_agriculture_cultures.pest_id
        INNER JOIN type_pest_entities ON type_pest_entities.id = pest_entities.type_pest_id
        INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = pest_agriculture_cultures.agriculture_culture_id
        WHERE pest_agriculture_cultures.pest_id = ? AND pest_agriculture_cultures.agriculture_culture_id = ?`

	if err := p.db.Raw(query, pestId, cultureId).Scan(&response).Error; err != nil {

		return &response, fmt.Errorf("erro ao buscar relação: %v", err)
	}

	if response == (responses.PestAgricultureCultureResponse{}) {
		return &response, fmt.Errorf("relação não encontrada")
	}

	return &response, nil
}

func (p *PestAgricultureCultureRepository) CreatePestAgricultureCulture(entityPestAgricultureCulture entities.PestAgricultureCulture) error {

	var exists bool
	if err := p.db.Raw(queryVerifyIdExist, entityPestAgricultureCulture.PestId, entityPestAgricultureCulture.AgricultureCultureId).Scan(&exists).Error; err != nil {
		return fmt.Errorf("erro ao verificar realcionamento entre as chaves estrangeiras")
	}

	if exists {
		return fmt.Errorf("não foi possível salvar no banco de dados. pest_id=%d e agriculture_culture_id=%d já estão relacionados", entityPestAgricultureCulture.PestId,
			entityPestAgricultureCulture.AgricultureCultureId)
	}

	if err := p.db.Create(&entityPestAgricultureCulture).Error; err != nil {
		return fmt.Errorf("erro ao criar objeto %v", err)
	}

	return nil
}

func (p *PestAgricultureCultureRepository) UpdatePestAgricultureCulture(pestId, cultureId uint, entityPestAgricultureCulture entities.PestAgricultureCulture) error {

	var exists bool

	if err := p.db.Raw(queryVerifyIdExist, pestId, cultureId).Scan(&exists).Error; err != nil {
		fmt.Println(err)
		return fmt.Errorf("erro ao verificar a existência do id %v", err)
	}

	if !exists {
		return fmt.Errorf("objeto com id fornecido não existe")
	}

	if err := p.db.Model(&entities.PestAgricultureCulture{}).
		Where("pest_agriculture_cultures.pest_id = ? AND pest_agriculture_cultures.agriculture_culture_id = ?", pestId, cultureId).
		Updates(&entityPestAgricultureCulture).Error; err != nil {

		return fmt.Errorf("erro ao atualizar objeto %v", err)
	}

	return nil

}

func (p *PestAgricultureCultureRepository) DeletePestAgricultureCulture(pestId, cultureId uint) error {

	var exists bool
	if err := p.db.Raw(queryVerifyIdExist, pestId, cultureId).Scan(&exists).Error; err != nil {

		return fmt.Errorf("erro ao verificar existência do id %v", err)
	}

	if !exists {
		return fmt.Errorf("não existe objeto com o id fornecido")
	}

	if err := p.db.Where("pest_id = ? AND agriculture_culture_id = ?", pestId, cultureId).Delete(entities.PestAgricultureCulture{}).Error; err != nil {
		return fmt.Errorf("erro ao deletar objeto %v", err)
	}

	return nil
}
