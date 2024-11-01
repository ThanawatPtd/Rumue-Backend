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

type PostgresTransactionRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

// SumThreeMonthIncome implements repositories.TransactionRepository.
func (tr *PostgresTransactionRepository) SumThreeMonthIncome(ctx context.Context) (*entities.Income, error) {
	income, err := tr.Queries.SumThreeMonth(ctx)
	if err != nil {
		return nil, err
	}
	var newIncome entities.Income
	newIncome.TotalIncome = float64(income)
	return &newIncome, err
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
		return nil, err
	}
	return newTransaction, nil
}

// ListByID implements repositories.TransactionRepository.
func (tr *PostgresTransactionRepository) ListByID(ctx context.Context, id string) ([]entities.UserVehicleTransaction, error) {
	uuid := convert.StringToUUID(id)
	selectTransactions, err := tr.Queries.GetAllTransactionsByUserID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	var userVehicleTransactions []entities.UserVehicleTransaction
	for _, value := range selectTransactions {
		var userVehicleTransaction entities.UserVehicleTransaction
		if err := utils.MappingParser(&value, &userVehicleTransaction.User); err != nil {
			return nil, err
		}
		if err := utils.MappingParser(&value, &userVehicleTransaction.Vehicle); err != nil {
			return nil, err
		}
		if err := utils.MappingParser(&value, &userVehicleTransaction.Transaction); err != nil {
			return nil, err
		}
		userVehicleTransactions = append(userVehicleTransactions, userVehicleTransaction)
	}
	return userVehicleTransactions, nil
}

// Update implements repositories.TransactionRepository.
func (tr *PostgresTransactionRepository) Update(ctx context.Context, transaction *entities.Transaction) error {
	var dbUpdateTransaction dbmodel.UpdateTransactionParams
	if err := utils.MappingParser(transaction, &dbUpdateTransaction); err != nil {
		return err
	}
	return tr.Queries.UpdateTransaction(ctx, dbUpdateTransaction)
}

// ListTrasactionToday implements repositories.EmployeeRepository.
func (tr *PostgresTransactionRepository) ListTrasactionToday(c context.Context) ([]entities.UserVehicleTransaction, error) {
	queryUserVehicleTransactions, err := tr.Queries.FindInsuranceToday(c)
	if err != nil {
		return nil, err
	}

	if queryUserVehicleTransactions == nil {
		return []entities.UserVehicleTransaction{}, nil
	}

	var userVehicleTransactions []entities.UserVehicleTransaction
	for _, value := range queryUserVehicleTransactions {
		var userVehicleTransaction entities.UserVehicleTransaction
		if err := utils.MappingParser(&value, &userVehicleTransaction.User); err != nil {
			return nil, err
		}
		if err := utils.MappingParser(&value, &userVehicleTransaction.Vehicle); err != nil {
			return nil, err
		}
		if err := utils.MappingParser(&value, &userVehicleTransaction.Transaction); err != nil {
			return nil, err
		}
		userVehicleTransactions = append(userVehicleTransactions, userVehicleTransaction)
	}
	return userVehicleTransactions, nil
}

// GetTransactionByID implements repositories.TransactionRepository.
func (tr *PostgresTransactionRepository) GetTransactionByID(ctx context.Context, transactionID string) (*entities.Transaction, error) {
	uuid := convert.StringToUUID(transactionID)
	transaction, err := tr.Queries.GetTransactionByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	newTransaction := &entities.Transaction{}
	if err := utils.MappingParser(&transaction, newTransaction); err != nil {
		return nil, err
	}
	return newTransaction, nil
}

// GetUserVehicleTransactionByID implements repositories.TransactionRepository.
func (tr *PostgresTransactionRepository) GetUserVehicleTransactionByID(ctx context.Context, trasactionID string) (*entities.UserVehicleTransaction, error) {
	uuid := convert.StringToUUID(trasactionID)
	userVehicleTransaction, err := tr.Queries.GetUserVehicleTransactionByID(ctx, uuid)

	if err != nil {
		return nil, err
	}
	newUserVehicleTransaction := entities.UserVehicleTransaction{}
	if err := utils.MappingParser(&userVehicleTransaction, &newUserVehicleTransaction.User); err != nil {
		return nil, err
	}
	if err := utils.MappingParser(&userVehicleTransaction, &newUserVehicleTransaction.Vehicle); err != nil {
		return nil, err
	}
	if err := utils.MappingParser(&userVehicleTransaction, &newUserVehicleTransaction.Transaction); err != nil {
		return nil, err
	}
	return &newUserVehicleTransaction, nil
}


func (tr *PostgresTransactionRepository) GetExpiredTransactionThisWeek(ctx context.Context) ([]entities.UserVehicleTransaction, error) {
	queryUserVehicleTransactions, err := tr.Queries.GetExpiredInsuranceTransactions(ctx)
	if err != nil {
		return nil, err
	}

	if queryUserVehicleTransactions == nil {
		return []entities.UserVehicleTransaction{}, nil
	}

	var userVehicleTransactions []entities.UserVehicleTransaction
	for _, value := range queryUserVehicleTransactions {
		var userVehicleTransaction entities.UserVehicleTransaction
		if err := utils.MappingParser(&value, &userVehicleTransaction.User); err != nil {
			return nil, err
		}
		if err := utils.MappingParser(&value, &userVehicleTransaction.Vehicle); err != nil {
			return nil, err
		}
		if err := utils.MappingParser(&value, &userVehicleTransaction.Transaction); err != nil {
			return nil, err
		}
		userVehicleTransactions = append(userVehicleTransactions, userVehicleTransaction)
	}
	return userVehicleTransactions, nil
}