package requests

type BatchRequest struct {
	Name string  `json:"name" validate:"required"`
	Area float32 `json:"area" validate:"required"`
	Unit string  `json:"unit" validate:"required"`
}
