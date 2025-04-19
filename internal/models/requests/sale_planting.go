package requests

type SalePlantingRequest struct {
	PlantingID uint    `json:"planting_id" validate:"required"`
	ValueSale  float32 `json:"value_sale" validate:"required"`
}
