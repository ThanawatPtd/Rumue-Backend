package entities

import "time"

type User struct {
	ID          string
	Email       string
	Fname       string
	Lname       string
	Password    string
	PhoneNumber string
	Address     string
	Nationality string
	CitizenID   string
	BirthDate   time.Time
}
