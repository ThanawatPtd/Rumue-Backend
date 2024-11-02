package responses

type GetInsurancesResponse struct {
	Tree map[string]map[string][]string
}

type GetInsuranceResponse struct {
	Price float64 `json:"price"`
}
