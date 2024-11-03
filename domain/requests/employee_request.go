package requests

type CreateEmployeeRequest struct {
	ID     string `json:"id" validate:"required"` 
	Salary float64 `json:"salary" validate:"required"`
}