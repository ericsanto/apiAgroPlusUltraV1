package responses

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"
)

type IrrigationTypeResponse struct {
	ID   uint                 `json:"id"`
	Name enums.IrrigationType `json:"name"`
}
