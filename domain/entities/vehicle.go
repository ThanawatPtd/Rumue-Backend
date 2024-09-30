package entities

import "time"

type Vehicle struct {
	VehicleId                       string
	RegistrationDate                time.Time
	RegistrationNumber              time.Time
	Province                        string
	VehicleType                     string
	VehicleCategory                 string
	Characteristics                 string
	Brand                           string
	Model                           string
	ModelYear                       string
	VehicleColor                    string
	EngineNumber                    string
	ChasisNumber                    string
	FuelType                        string
	HorsePower                      int32
	SeatingCapacity                 int32
	WeightUnlanden                  float64
	WeightLaden                     float64
	TireCount                       int32
	CompulsoryInsurancePolicyNumber string
	VoluntaryInsurancePolicyNumber  string
	InsuranceType                   string
}
