package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
)

type VehicleUseCase interface {
	CreateVehicle(ctx context.Context, userId string, vehicle *entities.Vehicle) (*entities.Vehicle, error)
	FindTemplate(ctx context.Context, userID string) ([]entities.Vehicle, error)
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

	createVehicle, err := v.vehicleRepo.CreateVehicle(ctx, vehicle)
	if err != nil {
		return nil, err
	}
	// Map User and Vehicle
	err = v.vehicleOwner.MapUserAndVehicle(ctx, userId, createVehicle.ID)
	// If map fail return error
	if err != nil {
		return nil, err
	}

	return createVehicle, nil
}

// FindTemplate implements VehicleUseCase.
func (v *VehicleService) FindTemplate(ctx context.Context, userID string) ([]entities.Vehicle, error) {
	vehicles, err := v.vehicleRepo.FindTemplate(ctx, userID)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
