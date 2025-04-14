package requests

type AgricultureCulturePestMethodRequest struct {
	AgricultureCultureId     uint   `json:"agriculture_culture_id" validate:"required"`
	PestId                   uint   `json:"pest_id" validate:"required"`
	SustainablePestControlId uint   `json:"sustainable_pest_control_id" validate:"required"`
	Description              string `json:"description" validate:"required"`
}
