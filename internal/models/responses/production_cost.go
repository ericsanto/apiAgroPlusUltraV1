package responses

import "time"

type ProductionCostResponse struct {
	ID          uint      `json:"id"`
	PlantingID  uint      `json:"planting_id"`
	Item        string    `json:"item_name"`
	Unit        string    `json:"unit"`
	Quantity    float32   `json:"quantity"`
	CostPerUnit float32   `json:"cost_per_unit"`
	CostDate    time.Time `json:"cost_date"`
}
