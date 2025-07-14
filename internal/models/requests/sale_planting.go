package requests

type SalePlantingRequest struct {
	ValueSale float32 `json:"value_sale" validate:"required"`
}
