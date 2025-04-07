package requests

type SustainablePestControlRequest struct {
	Name string `json:"name" validate:"required"`
}
