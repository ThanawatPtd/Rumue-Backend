package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresVehicleOwnerRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvidePostgresVehicleOwnerRepository(db *pgxpool.Pool) repositories.VehicleOwnerRepository {
	return &PostgresVehicleOwnerRepository{
		Queries: dbmodel.New(db),
		DB:      db,
	}
}

// MapUserAndVehicle implements repositories.VehicleOwnerRepository.
func (p *PostgresVehicleOwnerRepository) MapUserAndVehicle(ctx context.Context, userId string, vehicleId string) error {
	uuidUserId := convert.StringToUUID(userId)
	uuidVechicleId := convert.StringToUUID(vehicleId)

	createVehicleOwnderParam := dbmodel.CreateVehicleOwnerParams{
		UserID:    uuidUserId,
		VehicleID: uuidVechicleId,
	}

	_, err := p.Queries.CreateVehicleOwner(ctx, createVehicleOwnderParam)

	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresVehicleOwnerRepository) GetByID(ctx *context.Context, userID string, vehicleID string) (*entities.VehicleOwner, error) {
	uuidUserId := convert.StringToUUID(userID)
	uuidVechicleId := convert.StringToUUID(vehicleID)

	paramsVehicleOwner := dbmodel.GetVehicleOwnerByBothIdParams{
		UserID: uuidUserId,
		VehicleID: uuidVechicleId,
	}
	
	selectVehicleOwnerByBothID, err := p.Queries.GetVehicleOwnerByBothId(*ctx, paramsVehicleOwner)
	if err != nil {
		return nil, err
	}

	vehicleOwner := entities.VehicleOwner{}
	if err := utils.MappingParser(&selectVehicleOwnerByBothID, &vehicleOwner); err != nil {
		return nil, err
	}
	
	return &vehicleOwner, nil
}
