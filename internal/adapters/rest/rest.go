package rest

type Handler struct {
	User     *UserRestHandler
	Employee *EmployeeHandler
	Vehicle  *VehicleRestHandler
}

func ProvideHandler(user *UserRestHandler, employee *EmployeeHandler, vehicle *VehicleRestHandler) *Handler {
	return &Handler{
		User:     user,
		Employee: employee,
		Vehicle:  vehicle,
	}
}
