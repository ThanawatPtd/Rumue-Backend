package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/exceptions"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
)

type TransactionUseCase interface {
	CreateTransaction(ctx context.Context, userID string, vehicleID string, transaction *entities.Transaction) (*entities.Transaction, error)
	GetAllTransactions(ctx context.Context) ([]entities.Transaction, error)
	CheckHistory(ctx context.Context, userId string) ([]entities.UserVehicleTransaction, error)
	FindTodayInsurances(ctx context.Context) ([]entities.UserVehicleTransaction, error) //Today
	UpdateTransaction(ctx context.Context, transaction *entities.Transaction, id string) error
	FindTransactionByID(ctx context.Context, transactionID string) (*entities.UserVehicleTransaction, error)
}

type TransactionService struct {
	userRepo         repositories.UserRepository
	transactionRepo  repositories.TransactionRepository
	vehicleOwnerRepo repositories.VehicleOwnerRepository
}

func ProvideTransactionService(userRepo repositories.UserRepository, transactionRepo repositories.TransactionRepository, vehicleOwnerRepo repositories.VehicleOwnerRepository) TransactionUseCase {
	return &TransactionService{
		userRepo:         userRepo,
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

func (t *TransactionService) CheckHistory(ctx context.Context, id string) ([]entities.UserVehicleTransaction, error) {
	selectUser, err := t.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if selectUser == nil {
		return nil, exceptions.ErrUserNotFound // checkUser and not found User
	}

	userVehicletransactions, err := t.transactionRepo.ListByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return userVehicletransactions, nil
}

// FindTodayInsurances implements EmployeeUseCase.
func (t *TransactionService) FindTodayInsurances(ctx context.Context) ([]entities.UserVehicleTransaction, error) {
	userVehicleTransactions, err := t.transactionRepo.ListTrasactionToday(ctx)

	if err != nil {
		return nil, err
	}

	return userVehicleTransactions, nil
}

// UpdateTransaction implements EmployeeUseCase.
func (t *TransactionService) UpdateTransaction(ctx context.Context, transaction *entities.Transaction, id string) error {
	getTransaction, err := t.transactionRepo.GetTransactionByID(ctx, transaction.ID)
	if err != nil {
		return err
	}

	if getTransaction == nil {
		return exceptions.ErrTransactionNotFound
	}

	if transaction.CipNumber == "" {
		transaction.CipNumber = getTransaction.CipNumber
	}
	if transaction.VipNumber == "" {
		transaction.VipNumber = getTransaction.VipNumber
	}

	return t.transactionRepo.Update(ctx, transaction, id)
}

// FindTransactionByID implements TransactionUseCase.
func (t *TransactionService) FindTransactionByID(ctx context.Context, transactionID string) (*entities.UserVehicleTransaction, error) {
	userVehicleTransaction, err := t.transactionRepo.GetUserVehicleTransactionByID(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	return userVehicleTransaction, nil
}