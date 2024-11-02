package entities

import "time"

type Transaction struct {
	ID            string
	Price         float64
	InsuranceType string
	Status        string
	ESlipImageUrl string
	CrImageUrl    string
	CipNumber     string
	VipNumber     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
