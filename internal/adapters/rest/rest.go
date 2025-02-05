package rest

type Handler struct {
	User        *UserRestHandler
	Employee    *EmployeeHandler
	Vehicle     *VehicleRestHandler
	Transaction *TransactionRestHandler
	Auth        *AuthHandler
	Insurance   *InsuranceHandler
	Mile        *MileHandler
	Priority    *PriorityHandler
	Email       *EmailHandler
}

func ProvideHandler(user *UserRestHandler, employee *EmployeeHandler, vehicle *VehicleRestHandler,
	transaction *TransactionRestHandler, auth *AuthHandler, insurance *InsuranceHandler, mile *MileHandler, priority *PriorityHandler, email *EmailHandler) *Handler {
	return &Handler{
		User:        user,
		Employee:    employee,
		Vehicle:     vehicle,
		Transaction: transaction,
		Auth:        auth,
		Insurance:   insurance,
		Mile:        mile,
		Priority:    priority,
		Email:       email,
	}
}
