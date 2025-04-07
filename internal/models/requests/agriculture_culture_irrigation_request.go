package requests

type AgricultureCultureIrrigationRequest struct {
	AgricultureCultureID   uint `json:"agriculture_culture_id" validate:"required"`
	IrrigationRecomendedID uint `json:"irrigation_recomended_id" validate:"required"`
}
