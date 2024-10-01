package rest

type Handler struct {
	User     *UserRestHandler
	Employee *EmployeeHandler
	Vehicle  *VehicleRestHandler
	Invoice  *InvoiceRestHandler
}

func ProvideHandler(user *UserRestHandler, employee *EmployeeHandler, vehicle *VehicleRestHandler, invoice *InvoiceRestHandler) *Handler {
	return &Handler{
		User:     user,
		Employee: employee,
		Vehicle:  vehicle,
		Invoice:  invoice,
	}
}
