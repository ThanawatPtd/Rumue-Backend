package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
)

type VehicleUseCase interface {
	CreateVehicle(ctx context.Context, userId string, vehicle *entities.Vehicle) (*entities.Vehicle, error)
}

type VehicleService struct {
	vehicleRepo  repositories.VehicleRepository
	vehicleOwner repositories.VehicleOwnerRepository
}

func ProvideVehicleService(vehicleRepo repositories.VehicleRepository, vehicleOwnerRepo repositories.VehicleOwnerRepository) VehicleUseCase {
	return &VehicleService{
		vehicleRepo:  vehicleRepo,
		vehicleOwner: vehicleOwnerRepo,
	}
}

// CreateVehicle implements VehicleUseCase.
func (v *VehicleService) CreateVehicle(ctx context.Context, userId string, vehicle *entities.Vehicle) (*entities.Vehicle, error) {
	//find first later

	createVehicle, err := v.vehicleRepo.CreateVehicle(ctx, vehicle)
	if err != nil {
		return nil, err
	}
	// Map User and Vehicle
	err = v.vehicleOwner.MapUserAndVehicle(ctx, userId, createVehicle.VehicleId)
	// If map fail return error
	if err != nil {
		return nil, err
	}

	return createVehicle, nil
}
