package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InsuranceRepositoryImpl struct {
	Queries *dbmodel.Queries
}

func ProvideInsuranceRepository(db *pgxpool.Pool) repositories.InsuranceRepository {
	return &InsuranceRepositoryImpl{
		Queries: dbmodel.New(db),
	}
}

func (i *InsuranceRepositoryImpl) GetInsurance(c context.Context, req *entities.Insurance) (*entities.Insurance, error) {
	params := dbmodel.GetInsurancePriceParams{
		Brand: req.Brand,
		Model: req.Model,
		Year: req.Year,
	}

	price, err := i.Queries.GetInsurancePrice(c, params)
	if err != nil{
		return nil, err
	}

	response := &entities.Insurance{
		Price: price,
	}

	return response, nil
}

func (i *InsuranceRepositoryImpl) GetInsurances(c context.Context) ([]entities.Insurance, error) {
	row, err := i.Queries.GetInsurances(c)
    if err != nil {
        return nil, err
    }

    list := []entities.Insurance{}

    for _, val := range row {
        var insurance entities.Insurance
        if err := utils.MappingParser(&val, &insurance); err != nil {
            return nil, err
        }
        list = append(list, insurance)
    }

    return list, nil
}
