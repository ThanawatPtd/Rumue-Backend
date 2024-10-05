package responses

import "time"

type DefaultTransactionResponse struct {
	ID                string             `json:"id"`
	VehicleOwnerID    string 		     `json:"vehicleOwnerId"`
	TransactionType   string             `json:"transactionType"`
	TransactionStatus string             `json:"transactionStatus"`
	RequestDate       time.Time          `json:"requestDate"`
	ResponseDate      time.Time          `json:"responseDate"`
	ESlipImageUrl     string             `json:"eSlipImageUrl"`
}