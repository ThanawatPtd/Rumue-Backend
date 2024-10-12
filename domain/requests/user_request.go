package requests


type CreateUserRequest struct {
	Email       string `json:"email"`
	Fname       string `json:"fname"`
	Lname       string `json:"lname"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type UpdateUserRequest struct {
	Fname       string `json:"fname"`
	Lname       string `json:"lname"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}


type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
