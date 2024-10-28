package rest

type Handler struct {
	User        *UserRestHandler
	Employee    *EmployeeHandler
	Vehicle     *VehicleRestHandler
	Transaction *TransactionRestHandler
	Auth        *AuthHandler
}

func ProvideHandler(user *UserRestHandler, employee *EmployeeHandler, vehicle *VehicleRestHandler,
	transaction *TransactionRestHandler, auth *AuthHandler) *Handler {
	return &Handler{
		User:        user,
		Employee:    employee,
		Vehicle:     vehicle,
		Transaction: transaction,
		Auth:        auth,
	}
}
