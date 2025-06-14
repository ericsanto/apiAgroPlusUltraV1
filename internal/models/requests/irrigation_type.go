package requests

import "github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"

type IrrigationTypeRequest struct {
	Name enums.IrrigationType `json:"name" validate:"required"`
}
