package requests

import (
	"github.com/jackc/pgx/v5/pgtype"
)
	
type CreateUserRequest struct {
	Email       string `json:"email"`
	Fname       string `json:"first_name"`
	Lname       string `json:"last_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type GetUserByIDRowRequest struct {
	ID          pgtype.UUID `json:"id"`
	Email       string      `json:"email"`
	Fname       string      `json:"fname"`
	Lname       string      `json:"lname"`
	PhoneNumber string      `json:"phoneNumber"`
	Address     string      `json:"address"`
}


type UpdateUserRequest struct {
	ID          pgtype.UUID `json:"id"`
	Email       string      `json:"email"`
	Fname       string      `json:"fname"`
	Lname       string      `json:"lname"`
	Password    string      `json:"password"`
	PhoneNumber string      `json:"phoneNumber"`
	Address     string      `json:"address"`
}