package requests

import "time"

type CreateTransactionRequest struct {
	TransactionType   string             `json:"transactionType"`
	TransactionStatus string             `json:"transactionStatus"`
	RequestDate       time.Time			 `json:"requestDate"`
	ResponseDate      time.Time		     `json:"responseDate"`
	ESlipImageUrl     string             `json:"eSlipImageUrl"`
}