package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MileRepositoryImpl struct {
	Queries *dbmodel.Queries
}


func ProvidePostgresMileRepository(db *pgxpool.Pool) repositories.MileRepository {
	return &MileRepositoryImpl{
		Queries: dbmodel.New(db),
	}
}

func (m *MileRepositoryImpl) Save(c context.Context, mile *entities.Mile) (*entities.Mile, error) {
	paramsMile := new(dbmodel.CreateMileParams)
	if err := utils.MappingParser(mile, paramsMile); err != nil {
		return nil, err
	}

	created, err := m.Queries.CreateMile(c, *paramsMile)
	if err != nil {
		return nil, err
	}

	if err := utils.MappingParser(created, mile); err != nil {
		return nil, err
	}

	return mile, nil
}

func (m *MileRepositoryImpl) GetByID(c context.Context, id string) (*entities.Mile, error) {
	selected, err := m.Queries.GetMile(c, id)
	if err != nil{
		return nil ,err
	}

	mile := &entities.Mile{
		Rate: selected,
	}

	return mile, nil
}
