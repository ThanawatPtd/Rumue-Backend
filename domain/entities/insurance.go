package entities

type Insurance struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  string  `json:"year"`
	Price float64 `json:"price"`
}