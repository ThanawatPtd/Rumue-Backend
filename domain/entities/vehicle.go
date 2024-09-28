package entities

import "time"

type Vehicle struct {
	RegistrationDate                time.Time `json:"registrationDate"`
	RegistrationNumber              time.Time `json:"registrationNumber"`
	Province                        string    `json:"province"`
	VehicleType                     string    `json:"vehicleType"`
	VehicleCategory                 string    `json:"vehicleCategory"`
	Characteristics                 string    `json:"characteristics"`
	Brand                           string    `json:"brand"`
	Model                           string    `json:"model"`
	ModelYear                       string    `json:"modelYear"`
	VehicleColor                    string    `json:"vehicleColor"`
	EngineNumber                    string    `json:"engineNumber"`
	ChasisNumber                    string    `json:"chasisNumber"`
	FuelType                        string    `json:"fuelType"`
	HorsePower                      int32     `json:"horsePower"`
	SeatingCapacity                 int32     `json:"seatingCapacity"`
	WeightUnlanden                  float64   `json:"weightUnlanden"`
	WeightLaden                     float64   `json:"weightLaden"`
	TireCount                       int32     `json:"tireCount"`
	CompulsoryInsurancePolicyNumber string    `json:"compulsoryInsurancePolicyNumber"`
	VoluntaryInsurancePolicyNumber  string    `json:"voluntaryInsurancePolicyNumber"`
	InsuranceType                   string    `json:"insuranceType"`
}
