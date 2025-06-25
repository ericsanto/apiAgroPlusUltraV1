package requests

import "time"

type PlantingRequest struct {
	BatchID              uint      `json:"batch_id" validate:"required"`
	AgricultureCultureID uint      `json:"agriculture_culture_id" validate:"required"`
	IsPlanting           *bool     `json:"is_planting" validate:"required"`
	StartDatePlanting    time.Time `json:"start_date_planting"`
	SpaceBetweenPlants   float64   `json:"space_between_plants"`
	SpaceBetweenRows     float64   `json:"space_between_rows"`
	IrrigationTypeID     uint      `json:"irrigation_type_id"`
}
