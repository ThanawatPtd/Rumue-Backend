package rest

type Handler struct {
	User        *UserRestHandler
	Employee    *EmployeeHandler
	Vehicle     *VehicleRestHandler
	Invoice     *InvoiceRestHandler
	Transection *TransactionRestHandler
	Auth        *AuthHandler
}

func ProvideHandler(user *UserRestHandler, employee *EmployeeHandler, vehicle *VehicleRestHandler, invoice *InvoiceRestHandler,
	transaction *TransactionRestHandler, auth *AuthHandler) *Handler {
	return &Handler{
		User:        user,
		Employee:    employee,
		Vehicle:     vehicle,
		Invoice:     invoice,
		Transection: transaction,
		Auth:        auth,
	}
}

