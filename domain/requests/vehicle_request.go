package requests

import (
	"time"
)

type CreateVehicleRequest struct {
	RegistrationDate      time.Time `validate:"required"`
	RegistrationNumber    string    `validate:"required"`
	Province              string    `validate:"required"`
	VehicleType           string	`validate:"required"`
	VehicleCategory       string `validate:"required"`
	Characteristics       string 	`validate:"required"`
	Brand                 string	`validate:"required"`
	Model                 string `validate:"required"`
	ModelYear             string `validate:"required"`
	VehicleColor          string `validate:"required"`
	VehicleNumber         string `validate:"required"`
	VehicleNumberLocation string	`validate:"required"`
	EngineBrand           string `validate:"required"`
	EngineNumber          string  `validate:"required"`
	EngineNumberLocation  string `validate:"required"`
	ChasisNumber          string `validate:"required"`
	FuelType              string	`validate:"required"`
	WheelType             string 	`validate:"required"`
	TotalPiston           int32 	`validate:"required"`
	Cc                    int32 `validate:"required"`
	HorsePower            int32	`validate:"required"`
	SeatingCapacity       int32	`validate:"required"`
	WeightUnlanden        float64 `validate:"required"`
	WeightLaden           float64 `validate:"required"`
	Miles                 float64	`validate:"required"`
} 
