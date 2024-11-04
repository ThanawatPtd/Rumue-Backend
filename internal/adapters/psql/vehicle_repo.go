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

type PostgresVehicleRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvidePostgresVehicleRepository(db *pgxpool.Pool) repositories.VehicleRepository {
	return &PostgresVehicleRepository{
		Queries: dbmodel.New(db),
		DB:      db,
	}
}

// CreateVehicle implements repositories.VehicleRepository.
func (p *PostgresVehicleRepository) CreateVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error) {
	var createVehicleParam dbmodel.CreateVehicleParams
	if err := utils.MappingParser(vehicle, &createVehicleParam); err != nil {
		return nil, err
	}

	createVehicle, err := p.Queries.CreateVehicle(ctx, createVehicleParam)
	if err != nil {
		return nil, err
	}
	var vehicleRes entities.Vehicle
	vehicleID := convert.UUIDToString(createVehicle)
	vehicleRes.ID = vehicleID
	return &vehicleRes, nil
}

// FindTemplate implements repositories.VehicleRepository.
func (p *PostgresVehicleRepository) FindTemplate(ctx context.Context, userID string) ([]entities.Vehicle, error) {
	uuid := convert.StringToUUID(userID)

	getVehicles, err := p.Queries.FindAllTemplate(ctx, uuid)
	if err != nil {
		return nil, err
	}
	var vehicles []entities.Vehicle
	for _, value := range getVehicles {
		var vehicle entities.Vehicle
		if err := utils.MappingParser(&value, &vehicle); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}
