package responses

import "time"

type BatchPlantiesResponse struct {
	BatchName              string    `json:"batch_name"`
	IsPlanting             bool      `json:"is_planting"`
	AgricultureCultureName string    `json:"agriculture_culture_name"`
	StartDatePlanting      time.Time `json:"start_date_planting"`
}
