package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type TransactionRepository interface {
	ListAll(ctx context.Context) ([]entities.Transaction, error)
	Save(ctx context.Context, transaction *entities.Transaction, userID string, vehicleID string) (*entities.Transaction, error)
	ListByID(ctx context.Context, id string) ([]entities.Transaction, error)
}
