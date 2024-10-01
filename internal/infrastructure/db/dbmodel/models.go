// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbmodel

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Admin struct {
	ID        pgtype.UUID        `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

type Employee struct {
	ID        pgtype.UUID        `json:"id"`
	Salary    pgtype.Float4      `json:"salary"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

type EmployeeManagement struct {
	EmployeeID pgtype.UUID        `json:"employeeId"`
	AdminID    pgtype.UUID        `json:"adminId"`
	CreatedAt  pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt  pgtype.Timestamptz `json:"updatedAt"`
}

type User struct {
	ID          pgtype.UUID        `json:"id"`
	Email       string             `json:"email"`
	Fname       string             `json:"fname"`
	Lname       string             `json:"lname"`
	Password    string             `json:"password"`
	PhoneNumber string             `json:"phoneNumber"`
	Address     string             `json:"address"`
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt   pgtype.Timestamptz `json:"updatedAt"`
}

type Vehicle struct {
	VehicleID                       pgtype.UUID        `json:"vehicleId"`
	RegistrationDate                pgtype.Timestamptz `json:"registrationDate"`
	RegistrationNumber              string             `json:"registrationNumber"`
	Province                        string             `json:"province"`
	VehicleType                     string             `json:"vehicleType"`
	VehicleCategory                 string             `json:"vehicleCategory"`
	Characteristics                 string             `json:"characteristics"`
	Brand                           string             `json:"brand"`
	Model                           string             `json:"model"`
	ModelYear                       string             `json:"modelYear"`
	VehicleColor                    string             `json:"vehicleColor"`
	EngineNumber                    string             `json:"engineNumber"`
	ChasisNumber                    string             `json:"chasisNumber"`
	FuelType                        string             `json:"fuelType"`
	HorsePower                      int32              `json:"horsePower"`
	SeatingCapacity                 int32              `json:"seatingCapacity"`
	WeightUnlanden                  float64            `json:"weightUnlanden"`
	WeightLaden                     float64            `json:"weightLaden"`
	TireCount                       int32              `json:"tireCount"`
	CompulsoryInsurancePolicyNumber string             `json:"compulsoryInsurancePolicyNumber"`
	VoluntaryInsurancePolicyNumber  pgtype.Text        `json:"voluntaryInsurancePolicyNumber"`
	InsuranceType                   pgtype.Text        `json:"insuranceType"`
	CreatedAt                       pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt                       pgtype.Timestamptz `json:"updatedAt"`
}

type VehicleOwner struct {
	UserID    pgtype.UUID `json:"userId"`
	VehicleID pgtype.UUID `json:"vehicleId"`
}
