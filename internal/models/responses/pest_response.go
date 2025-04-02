package responses

type PestResponse struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	TypePestId uint   `json:"type_pest_id"`
}
