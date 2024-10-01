package repositories

import "context"

type VehicleOwnerRepository interface {
	MapUserAndVehicle(ctx context.Context, userId string, vehicleId string) error
}
