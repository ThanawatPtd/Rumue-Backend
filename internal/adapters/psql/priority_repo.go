package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PriorityRepositoryImpl struct {
	Queries *dbmodel.Queries
}

func ProvidePostgresPriorityRepository(db *pgxpool.Pool) repositories.PriorityRepository {
	return &PriorityRepositoryImpl{
		Queries: dbmodel.New(db),
	}
}

func (p *PriorityRepositoryImpl) Save(c context.Context, priority *entities.Priority) (*entities.Priority, error) {
	paramsPriority := new(dbmodel.CreatePriorityParams)
	if err := utils.MappingParser(priority, paramsPriority); err != nil {
		return nil, err
	}

	created, err := p.Queries.CreatePriority(c, *paramsPriority)
	if err != nil {
		return nil, err
	}

	if err := utils.MappingParser(created, priority); err != nil {
		return nil, err
	}

	return priority, nil
}

func (p *PriorityRepositoryImpl) GetByID(c context.Context, id string) (*entities.Priority, error) {
	selected, err := p.Queries.GetPriority(c, id)
	if err != nil {
		return nil, err
	}

	priority := &entities.Priority{
		Rate: selected,
	}

	return priority, nil
}
