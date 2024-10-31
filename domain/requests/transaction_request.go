package requests

type CreateTransactionRequest struct {
	Price         float64 `json:"price"`
	InsuranceType string  `json:"insuranceType"`
	Status        string  `json:"status"`
	ESlipImageUrl string  `json:"eSlipImageUrl"`
	CrImageUrl    string  `json:"crImageUrl"`
	CipNumber     string  `json:"cipNumber"`
	VipNumber     string  `json:"vipNumber"`
}

type UpdateTransactionRequest struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	CipNumber string `json:"cipNumber"`
	VipNumber string `json:"vipNumber"`
}
