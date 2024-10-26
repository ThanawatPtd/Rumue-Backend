package psql

import (
	"context"
	"fmt"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/emicklei/pgtalk/convert"
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

func (tr *PostgresTransactionRepository) ListAll(ctx context.Context) ([]entities.Transaction, error) {
	return nil, nil
}

func (tr *PostgresTransactionRepository) Save(ctx context.Context, transaction *entities.Transaction, userID string, vehicleID string) (*entities.Transaction, error) {
	var dbTransaction dbmodel.CreateTransactionParams
	if err := utils.MappingParser(transaction, &dbTransaction); err != nil {
		return nil, err
	}
	convertedUserID := convert.StringToUUID(userID)
	convertedVehicleID := convert.StringToUUID(vehicleID)
	dbTransaction.UserID = convertedUserID
	dbTransaction.VehicleID = convertedVehicleID
	savedTransaction, err := tr.Queries.CreateTransaction(ctx, dbTransaction)
	if err != nil {
		return nil, err
	}
	newTransaction := &entities.Transaction{}
	if err := utils.MappingParser(&savedTransaction, newTransaction); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return newTransaction, nil
}

// ListByID implements repositories.TransactionRepository.
func (tr *PostgresTransactionRepository) ListByID(ctx context.Context, id string) ([]entities.Transaction, error) {
	uuid := convert.StringToUUID(id)
	selectTransaction, err := tr.Queries.GetAllTransactionsByUserID(ctx, uuid)
	if err != nil {
		return []entities.Transaction{}, err
	}
	var transactions []entities.Transaction
	for _, value := range selectTransaction {
		var transaction entities.Transaction
		if err := utils.MappingParser(&value, &transactions); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
