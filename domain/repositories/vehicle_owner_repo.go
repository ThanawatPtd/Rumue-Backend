package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type VehicleOwnerRepository interface {
	MapUserAndVehicle(ctx context.Context, userId string, vehicleId string) error
	GetByID(ctx context.Context, userID string, vehicleID string) (*entities.VehicleOwner, error)
}
