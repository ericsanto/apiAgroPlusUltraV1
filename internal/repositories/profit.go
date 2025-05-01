package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"gorm.io/gorm"
)

type ProfitRepository struct {
	db *gorm.DB
}

func NewProfitRepository(db *gorm.DB) *ProfitRepository {
	return &ProfitRepository{db: db}
}

func (p *ProfitRepository) FindProfit(plantingID uint) (*responses.ProfitResponse, error) {

	newPlantingRepository := NewPlantingRepository(p.db)

	if _, err := newPlantingRepository.FindPlantingByID(plantingID); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	query := `SELECT 
	SUM(production_cost_entities.quantity * production_cost_entities.cost_per_unit) AS total_cost, 
	SUM(sale_planting_entities.value_sale) AS value_sale_plantation, 
	SUM(sale_planting_entities.value_sale - (production_cost_entities.quantity * production_cost_entities.cost_per_unit)) AS profit,
	SUM((sale_planting_entities.value_sale - (production_cost_entities.quantity * production_cost_entities.cost_per_unit)) / 100) AS profit_margin
	
	FROM planting_entities
	INNER JOIN production_cost_entities ON production_cost_entities.planting_id = planting_entities.id
	INNER JOIN sale_planting_entities ON sale_planting_entities.planting_id = planting_entities.id
	
	WHERE planting_entities.id = ?`

	var responseProfit responses.ProfitResponse

	if err := p.db.Raw(query, plantingID).Scan(&responseProfit).Error; err != nil {
		return nil, fmt.Errorf("erro ao calcular lucro")
	}

	return &responseProfit, nil

}
