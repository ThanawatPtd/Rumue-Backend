package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
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
	createVehicleParam = dbmodel.CreateVehicleParams{
		RegistrationDate:                convert.TimeToTimestamptz(vehicle.RegistrationDate),
		RegistrationNumber:              vehicle.RegistrationNumber,
		Province:                        vehicle.Province,
		VehicleType:                     vehicle.VehicleType,
		VehicleCategory:                 vehicle.VehicleCategory,
		Characteristics:                 vehicle.Characteristics,
		Brand:                           vehicle.Brand,
		Model:                           vehicle.Model,
		ModelYear:                       vehicle.ModelYear,
		VehicleColor:                    vehicle.VehicleColor,
		EngineNumber:                    vehicle.EngineNumber,
		ChasisNumber:                    vehicle.ChasisNumber,
		FuelType:                        vehicle.FuelType,
		HorsePower:                      vehicle.HorsePower,
		SeatingCapacity:                 vehicle.SeatingCapacity,
		WeightUnlanden:                  vehicle.WeightLaden,
		WeightLaden:                     vehicle.WeightLaden,
		TireCount:                       vehicle.TireCount,
		CompulsoryInsurancePolicyNumber: vehicle.CompulsoryInsurancePolicyNumber,
		VoluntaryInsurancePolicyNumber:  convert.StringToText(vehicle.VoluntaryInsurancePolicyNumber),
		InsuranceType:                   convert.StringToText(vehicle.InsuranceType),
	}

	createVehicle, err := p.Queries.CreateVehicle(ctx, createVehicleParam)
	if err != nil {
		return nil, err
	}
	var vehicleRes entities.Vehicle
	vehicleRes = entities.Vehicle{
		ID:                       convert.UUIDToString(createVehicle.ID),
		RegistrationDate:                createVehicle.RegistrationDate.Time,
		RegistrationNumber:              createVehicle.RegistrationNumber,
		Province:                        createVehicle.Province,
		VehicleType:                     createVehicle.VehicleType,
		VehicleCategory:                 createVehicle.VehicleCategory,
		Characteristics:                 createVehicle.Characteristics,
		Brand:                           createVehicle.Brand,
		Model:                           createVehicle.Model,
		ModelYear:                       createVehicle.ModelYear,
		VehicleColor:                    createVehicle.VehicleColor,
		EngineNumber:                    createVehicle.EngineNumber,
		ChasisNumber:                    createVehicle.ChasisNumber,
		FuelType:                        createVehicle.FuelType,
		HorsePower:                      createVehicle.HorsePower,
		SeatingCapacity:                 createVehicle.SeatingCapacity,
		WeightUnlanden:                  createVehicle.WeightUnlanden,
		WeightLaden:                     createVehicle.WeightUnlanden,
		TireCount:                       createVehicle.TireCount,
		CompulsoryInsurancePolicyNumber: createVehicle.CompulsoryInsurancePolicyNumber,
		VoluntaryInsurancePolicyNumber:  createVehicle.VoluntaryInsurancePolicyNumber.String,
		InsuranceType:                   createVehicle.InsuranceType.String,
	}

	return &vehicleRes, nil
}
