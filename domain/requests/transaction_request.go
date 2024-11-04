package requests

type CreateTransactionRequest struct {
	Price         float64 `json:"price" validate:"required"`
	InsuranceType string  `json:"insuranceType" validate:"required"`
	Status        string  `json:"status" validate:"required"`
	ESlipImageUrl string  `json:"eSlipImageUrl" validate:"required"`
	CrImageUrl    string  `json:"crImageUrl" validate:"required"`
	CipNumber     string  `json:"cipNumber"`
	VipNumber     string  `json:"vipNumber"`
}

type UpdateTransactionRequest struct {
	ID        string `json:"id" validate:"required"`
	Status    string `json:"status" validate:"required"`
	CipNumber string `json:"cipNumber"`
	VipNumber string `json:"vipNumber"`
}
