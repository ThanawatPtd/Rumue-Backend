package responses

import (
	"time"
)

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
	Fname       string    `json:"Fname"`
	Lname       string    `json:"Lname"`
	Email       string    `json:"Email"`
	PhoneNumber string    `json:"PhoneNumber"`
	Address     string    `json:"Address"`
	Nationality string    `json:"Nationality"`
	CitizenID   string    `json:"CitizenID"`
	BirthDate   time.Time `json:"BirthDate"`	
}

type UserProfileResponse struct {
	Fname       string    `json:"fname"`
	Lname       string    `json:"lname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Nationality string    `json:"nationality"`
	CitizenID   string    `json:"citizenID"`
	BirthDate   time.Time `json:"birthDate"`	
	Salary      float64		`json:"salary"` 
}
