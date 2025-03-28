package requests


type TypePestRequest struct {

  Name  string `json:"name" validate:"required,min=5"`
}
