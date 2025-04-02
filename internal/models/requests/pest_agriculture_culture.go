package requests

type PestAgricultureCultureRequest struct {
	AgricultureCultureId uint   `json:"agriculture_culture_id" validate:"required"`
	PestId               uint   `json:"pest_id" validate:"required"`
	Description          string `json:"description" validate:"required,min=10"`
	ImageUrl             string `json:"image" validate:"required,min=17"`
}
