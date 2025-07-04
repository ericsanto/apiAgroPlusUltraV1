package responses

import "time"

type PlantingResponse struct {
	ID                   uint      `json:"id"`
	BatchID              uint      `json:"batch_id"`
	AgricultureCultureID uint      `json:"agriculture_culture_id"`
	IsPlanting           bool      `json:"is_planting"`
	StartDatePlanting    time.Time `json:"start_date_planting"`
	SpaceBetweenPlants   float64   `json:"space_between_plants"`
	ExpectedProduction   float64   `json:"expected_production"`
	SpaceBetweenRows     float64   `json:"space_between_rows"`
	IrrigationTypeID     uint      `json:"irrigation_type_id"`
}
