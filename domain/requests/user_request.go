package requests

type CreateUserRequest struct {
	Email       string `json:"email"`
	Fname       string `json:"first_name"`
	Lname       string `json:"last_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
