package responses

type SalePlantingResponse struct {
	ID         uint    `json:"id"`
	PlantingID uint    `json:"planting_id"`
	ValueSale  float32 `json:"value_sale"`
}
