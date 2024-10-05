package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTransactionRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvidePostgresTransactionRepository(db *pgxpool.Pool) repositories.TransactionRepository {
	return &PostgresTransactionRepository{
		Queries: dbmodel.New(db),
		DB:      db,
	}
}

func (tr *PostgresTransactionRepository) ListAll(ctx *context.Context) (*[]entities.Transaction, error) {
	return nil, nil
}

func (tr *PostgresTransactionRepository) Save(ctx *context.Context, transaction *entities.Transaction) (*entities.Transaction, error) {
	return nil, nil
}