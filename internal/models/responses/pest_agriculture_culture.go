package responses


type PestAgricultureCultureResponse struct {

  AgricultureCultureName string `json:"agriculture_culture_name"`
  PestName               string  `json:"pest_name"`
  Description            string `json:"description"`
  ImageUrl               string `json:"image_url"`
}
