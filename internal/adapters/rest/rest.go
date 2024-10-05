package rest

type Handler struct {
	User     *UserRestHandler
	Employee *EmployeeHandler
	Vehicle  *VehicleRestHandler
	Invoice  *InvoiceRestHandler
	Transection *TransactionRestHandler
}

func ProvideHandler(user *UserRestHandler, employee *EmployeeHandler, vehicle *VehicleRestHandler, invoice *InvoiceRestHandler, 
	transaction *TransactionRestHandler) *Handler {
	return &Handler{
		User:     user,
		Employee: employee,
		Vehicle:  vehicle,
		Invoice:  invoice,
		Transection: transaction,
	}
}