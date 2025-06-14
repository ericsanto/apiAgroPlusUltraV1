package responses

import "time"

type BatchPlantiesResponse struct {
	BatchName              string    `json:"batch_name"`
	IsPlanting             bool      `json:"is_planting"`
	AgricultureCultureName string    `json:"agriculture_culture_name"`
	SoilTypeName           string    `json:"soil_type_name"`
	StartDatePlanting      time.Time `json:"start_date_planting"`
	SpaceBetweenPlants     float64   `json:"space_between_plants"`
	SpaceBetweenRows       float64   `json:"space_between_rows"`
	IrrigationType         string    `json:"irrigation_type"`
}
