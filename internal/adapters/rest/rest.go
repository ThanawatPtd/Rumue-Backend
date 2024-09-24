package rest

type Handler struct {
	User *UserRestHandler
	Employee *EmployeeHandler
}

func ProvideHandler(user *UserRestHandler, employee *EmployeeHandler) *Handler {
	return &Handler{
		User: user,
		Employee: employee,
	}
}
