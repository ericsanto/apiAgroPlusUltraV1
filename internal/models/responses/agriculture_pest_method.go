package responses

type AgricultureCulturePestMethodResponse struct {
	AgricultureCultureName       string `json:"agriculture_culture_name"`
	PestName                     string `json:"pest_name"`
	SustainablePestControlMethod string `json:"sustainable_pest_control_method"`
	Description                  string `json:"description"`
}
