package responses

import "time"

type PlantingResponse struct {
	ID                   uint      `json:"id"`
	BatchID              uint      `json:"batch_id"`
	AgricultureCultureID uint      `json:"agriculture_culture_id"`
	IsPlanting           bool      `json:"is_planting"`
	StartDatePlanting    time.Time `json:"start_date_planting"`
}
