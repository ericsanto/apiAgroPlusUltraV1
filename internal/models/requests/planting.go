package requests

type PlantingRequest struct {
	BatchID              uint  `json:"batch_id" validate:"required"`
	AgricultureCultureID uint  `json:"agriculture_culture_id" validate:"required"`
	IsPlanting           *bool `json:"is_planting" validate:"required"`
}
