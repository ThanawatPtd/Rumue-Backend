package responses

type UserLoginResponse struct {
	ID    string `json:"id"`
	Fname  string `json:"fname"`
	Lname	string `json:"lname"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserRegisterResponse struct {
	ID          string `json:"id"`
	Fname  string `json:"fname"`
	Lname	string `json:"lname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type UserDefaultResponse struct {
	ID          string `json:"id"`
	Fname  string `json:"fname"`
	Lname	string `json:"lname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
