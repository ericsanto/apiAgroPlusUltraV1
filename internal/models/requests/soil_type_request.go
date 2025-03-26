package requests


type SoilTypeRequest struct { 
  Name        string `json:"name" binding:"required,min=5"`
  Description string `json:"description" binding:"required,min=20"`

}


