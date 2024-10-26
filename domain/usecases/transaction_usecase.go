package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
)

type TransactionUseCase interface {
	CreateTransaction(ctx context.Context, userID string, vehicleID string, transaction *entities.Transaction) (*entities.Transaction, error)
	GetAllTransactions(ctx context.Context) ([]entities.Transaction, error)
	GetAllTransactionsByID(ctx context.Context, userId string) ([]entities.Transaction, error)
}

type TransactionService struct {
	transactionRepo  repositories.TransactionRepository
	vehicleOwnerRepo repositories.VehicleOwnerRepository
}

func ProvideTransactionService(transactionRepo repositories.TransactionRepository, vehicleOwnerRepo repositories.VehicleOwnerRepository) TransactionUseCase {
	return &TransactionService{
		transactionRepo:  transactionRepo,
		vehicleOwnerRepo: vehicleOwnerRepo,
	}
}

func (t *TransactionService) CreateTransaction(ctx context.Context, userID string, vehicleID string, transaction *entities.Transaction) (*entities.Transaction, error) {
	vehicle, err := t.vehicleOwnerRepo.GetByID(ctx, userID, vehicleID)
	if err != nil {
		return nil, err
	}
	if vehicle == nil {
		return nil, err //Vehicle Not found cant create transaction
	}
	savedTransaction, err := t.transactionRepo.Save(ctx, transaction, userID, vehicleID)
	if err != nil {
		return nil, err
	}

	return savedTransaction, nil
}

func (t *TransactionService) GetAllTransactions(ctx context.Context) ([]entities.Transaction, error) {
	transactions, err := t.transactionRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *TransactionService) GetAllTransactionsByID(ctx context.Context, id string) ([]entities.Transaction, error) {
	// selectUser, err := u.userRepo.GetByID(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// if selectUser == nil {
	// 	return nil, exceptions.ErrUserNotFound // checkUser and not found User
	// }

	transactions, err := t.transactionRepo.ListByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
