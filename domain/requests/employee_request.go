package requests

type CreateEmployeeRequest struct {
	ID     string `json:"id"`
	Salary float64 `json:"salary"`
}