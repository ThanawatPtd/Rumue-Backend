package rest

type Handler struct {
	User *UserRestHandler
}

func ProvideHandler(user *UserRestHandler) *Handler {
	return &Handler{
		User: user,
	}
}
