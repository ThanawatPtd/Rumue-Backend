package responses

import "time"

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserRegisterResponse struct {
	ID          string    `json:"id"`
	Fname       string    `json:"fname"`
	Lname       string    `json:"lname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Nationality string    `json:"nationality"`
	CitizenID   string    `json:"citizenID"`
	BirthDate   time.Time `json:"birthDate"`
}

type UserDefaultResponse struct {
	ID          string `json:"id"`
	Fname       string `json:"fname"`
	Lname       string `json:"lname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
