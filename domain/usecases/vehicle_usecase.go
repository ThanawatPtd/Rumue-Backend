package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
)

type VehicleUseCase interface {
	CreateVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error)
}

type VehicleService struct {
	vehicleRepo repositories.VehicleRepository
}


func ProvideVehicleService(vehicleRepo repositories.VehicleRepository) VehicleUseCase {
	return &VehicleService{
		vehicleRepo: vehicleRepo,
	}
}
// CreateVehicle implements VehicleUseCase.
func (v *VehicleService) CreateVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error) {
	//find first later

	createVehicle, err := v.vehicleRepo.CreateVehicle(ctx, vehicle)
	if err != nil{
		return nil, err
	}

	return createVehicle, nil
}
