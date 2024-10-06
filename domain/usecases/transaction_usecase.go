package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
)

type TransactionUseCase interface {
	CreateTransaction(ctx context.Context, userID string, vehicleID string, transaction *entities.Transaction) (*entities.Transaction, error)
	GetAllTransactions(ctx context.Context) (*[]entities.Transaction, error)
}

type TransactionService struct {
	transactionRepo  repositories.TransactionRepository
	vehicleOwnerRepo repositories.VehicleOwnerRepository
}

func ProvideTransactionService(transactionRepo repositories.TransactionRepository, vehicleOwnerRepo repositories.VehicleOwnerRepository) TransactionUseCase {
	return &TransactionService{
		transactionRepo: transactionRepo,
		vehicleOwnerRepo: vehicleOwnerRepo,
	}
}

func (t *TransactionService) CreateTransaction(ctx context.Context, userID string, vehicleID string, transaction *entities.Transaction) (*entities.Transaction, error) {
	vehicle, err := t.vehicleOwnerRepo.GetByID(&ctx, userID, vehicleID)
	if err != nil {
		return nil, err
	}

	transaction.VehicleOwnerID = vehicle.ID

	transaction, err = t.transactionRepo.Save(&ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionService) GetAllTransactions(ctx context.Context) (*[]entities.Transaction, error) {
	transactions, err := t.transactionRepo.ListAll(&ctx)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
 

