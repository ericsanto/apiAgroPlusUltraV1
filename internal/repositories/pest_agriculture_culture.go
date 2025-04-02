package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"gorm.io/gorm"
)

type PestAgricultureCultureRepositoryInterface interface {

	FindAllPestAgricultureCulture() ([]responses.PestAgricultureCultureResponse, error)
	FindByIdPestAgricultureCulture(id uint) (responses.PestAgricultureCultureResponse, error)
}


type PestAgricultureCultureRepository struct {

  db *gorm.DB
}

func NewPestAgricultureCultureRepository(db *gorm.DB) *PestAgricultureCultureRepository {
  return &PestAgricultureCultureRepository{db:db}
}

func(p *PestAgricultureCultureRepository) FindAllPestAgricultureCulture() ([]responses.PestAgricultureCultureResponse, error){

  var responsePestAgricultureCulutre []responses.PestAgricultureCultureResponse
  
  query := `SELECT pest_entities.name AS pest_name, type_pest_entities.name AS type_pest_name,
  agriculture_culture_entities.name AS agriculture_culture_name, description, image_url

  FROM pest_agriculture_cultures

  INNER JOIN pest_entities ON pest_entities.id = pest_agriculture_cultures.pest_id

  INNER JOIN type_pest_entities ON type_pest_entities.id = pest_entities.type_pest_id

  INNER JOIN agriculture_culture_entities ON agriculture_culture_entities.id = pest_agriculture_cultures.agriculture_culture_id`

  result := p.db.Raw(query).Scan(&responsePestAgricultureCulutre)
  if result.Error != nil {
    return responsePestAgricultureCulutre, fmt.Errorf("Erro no repositório ao fazer consulta %w", result.Error)
  }

	fmt.Println(result)

  return responsePestAgricultureCulutre, nil
}

func (p *PestAgricultureCultureRepository) FindByIdPestAgricultureCulture(pestId, cultureId uint) (responses.PestAgricultureCultureResponse, error) {
    var response responses.PestAgricultureCultureResponse
    

		query :=`SELECT 
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

    err := p.db.Raw(query,pestId, cultureId).Scan(&response).Error

    if err != nil {
        // if errors.Is(err, gorm.ErrRecordNotFound) {
        //     return response, fmt.Errorf("relação não encontrada")
        // }
         return response, fmt.Errorf("erro ao buscar relação: %w", err)
    }
    
    if response == (responses.PestAgricultureCultureResponse{}) {
        return response, fmt.Errorf("relação não encontrada")
    }
    
    return response, nil
}
