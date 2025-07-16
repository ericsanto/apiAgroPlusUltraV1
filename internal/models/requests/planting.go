package requests

import "time"

type PlantingRequest struct {
	AgricultureCultureID uint      `json:"agriculture_culture_id" validate:"required"`
	IsPlanting           *bool     `json:"is_planting" validate:"required"`
	StartDatePlanting    time.Time `json:"start_date_planting"`
	ExpectedProduction   float64   `json:"expected_production"`
	SpaceBetweenPlants   float64   `json:"space_between_plants"`
	SpaceBetweenRows     float64   `json:"space_between_rows"`
	IrrigationTypeID     uint      `json:"irrigation_type_id"`
}
