package requests


type PestRequest struct {

  Name     string `json:"name" validate:"required,min=10"`
  TypePestId uint   `json:"type_pest_id" validate:"required"`
}
