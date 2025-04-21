package responses

type ProfitResponse struct {
	ValueSalePlantation float64 `json:"value_sale_plantiation"`
	TotalCost           float64 `json:"total_cost"`
	Profit              float64 `json:"profit"`
	ProfitMargin        float64 `json:"profit_margen"`
}
