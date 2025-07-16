package repositories

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

type ProfitRepositoryInterface interface {
	FindProfit(batchID, farmID, userID, plantingID uint) (*responses.ProfitResponse, error)
}

type ProfitRepository struct {
	db                 interfaces.GORMRepositoryInterface
	plantingRepository PlantingRepositoryInterface
}

func NewProfitRepository(db interfaces.GORMRepositoryInterface, plantingRepository PlantingRepositoryInterface) ProfitRepositoryInterface {
	return &ProfitRepository{db: db, plantingRepository: plantingRepository}
}

func (p *ProfitRepository) FindProfit(batchID, farmID, userID, plantingID uint) (*responses.ProfitResponse, error) {

	if _, err := p.plantingRepository.FindPlantingByID(batchID, farmID, userID, plantingID); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	query := `SELECT 
	batch_entities.id AS batch_id,
	farm_entities.id AS farm_id,
	user_models.id AS user_id,
	SUM(production_cost_entities.quantity * production_cost_entities.cost_per_unit) AS total_cost, 
	SUM(sale_planting_entities.value_sale) AS value_sale_plantation, 
	SUM(sale_planting_entities.value_sale - (production_cost_entities.quantity * production_cost_entities.cost_per_unit)) AS profit,
	SUM((sale_planting_entities.value_sale - (production_cost_entities.quantity * production_cost_entities.cost_per_unit)) / 100) AS profit_margin
	
	FROM planting_entities
	INNER JOIN production_cost_entities ON production_cost_entities.planting_id = planting_entities.id
	INNER JOIN sale_planting_entities ON sale_planting_entities.planting_id = planting_entities.id
	INNER JOIN batch_entities ON batch_entities.id = planting_entities.batch_id
	INNER  JOIN farm_entities ON farm_entities.id = batch_entities.farm_id
	INNER JOIN user_models ON user_models.id = farm_entities.user_id
	WHERE planting_entities.id = ? AND user_models.id = ? AND farm_entities.id = ? AND batch_entities.id = ?
	GROUP BY farm_entities.id, batch_entities.id, user_models.id, planting_entities.id`

	var responseProfit responses.ProfitResponse

	if err := p.db.Raw(query, plantingID, userID, farmID, batchID).Scan(&responseProfit).Error; err != nil {
		return nil, fmt.Errorf("erro ao calcular lucro")
	}

	return &responseProfit, nil

}
