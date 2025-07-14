package requests

import "time"

type ProductionCostRequest struct {
	Item        string    `json:"item_name" validate:"required"`
	Unit        string    `json:"unit" validate:"required"`
	Quantity    float32   `json:"quantity" validate:"required"`
	CostPerUnit float32   `json:"cost_per_unit" validate:"required"`
	CostDate    time.Time `json:"cost_date" validate:"required"`
}
