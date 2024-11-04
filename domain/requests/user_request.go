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
	Fname       string    `json:"fname" validate:"required"`
	Lname       string    `json:"lname" validate:"required"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
	Address     string    `json:"address" validate:"required"`
	Nationality string    `json:"nationality" validate:"required"`
	BirthDate   time.Time `json:"birthDate" validate:"required"`
	CitizenID   string    `json:"citizenID" validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
