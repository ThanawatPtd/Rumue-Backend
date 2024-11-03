package repositories

import (
	"context"
	"time"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type TransactionRepository interface {
	ListAll(ctx context.Context) ([]entities.Transaction, error)
	Save(ctx context.Context, transaction *entities.Transaction, userID string, vehicleID string) (*entities.Transaction, error)
	ListByID(ctx context.Context, id string) ([]entities.UserVehicleTransaction, error) // Get list of transaction
	Update(ctx context.Context, transaction *entities.Transaction, id string) error
	UpdateReceiptDate(ctx context.Context, id string) (*time.Time, error)
	ListTrasactionToday(c context.Context) ([]entities.UserVehicleTransaction, error)
	GetTransactionByID(ctx context.Context, transactionID string) (*entities.Transaction, error)
	GetExpiredTransactionThisWeek(ctx context.Context) ([]entities.UserVehicleTransaction, error)
	GetUserVehicleTransactionByID(ctx context.Context, transactionID string) (*entities.UserVehicleTransaction, error)
}
