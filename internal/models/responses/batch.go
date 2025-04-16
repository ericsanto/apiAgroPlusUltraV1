package responses

type BatchResponse struct {
	ID   uint    `json:"id"`
	Name string  `json:"name"`
	Area float32 `json:"area"`
	Unit string  `json:"unit"`
}
