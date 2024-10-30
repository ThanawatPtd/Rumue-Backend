package entities

import "time"

type Vehicle struct {
	ID                    string
	RegistrationDate      time.Time
	RegistrationNumber    string
	Province              string
	VehicleType           string
	VehicleCategory       string
	Characteristics       string
	Brand                 string
	Model                 string
	ModelYear             string
	VehicleColor          string
	VehicleNumber         string
	VehicleNumberLocation string
	EngineBrand           string
	EngineNumber          string
	EngineNumberLocation  string
	ChasisNumber          string
	FuelType              string
	WheelType             string
	TotalPiston           int32
	Cc                    int32
	HorsePower            int32
	SeatingCapacity       int32
	WeightUnlanden        float64
	WeightLaden           float64
	Miles                 float64
}
