package requests

import "time"

type CreateUserRequest struct {
	Email       string    `json:"email" validate:"required"`
	Fname       string    `json:"fname" validate:"required"`
	Lname       string    `json:"lname" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
	Address     string    `json:"address" validate:"required"`
	Nationality string    `json:"nationality" validate:"required"`
	BirthDate   time.Time `json:"birthDate" validate:"required"`
	CitizenID   string    `json:"citizenID" validate:"required"`
}

type UpdateUserRequest struct {
	Fname       string    `json:"fname"`
	Lname       string    `json:"lname"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Nationality string    `json:"nationality"`
	BirthDate   time.Time `json:"birthDate"`
	CitizenID   string    `json:"citizenID"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
