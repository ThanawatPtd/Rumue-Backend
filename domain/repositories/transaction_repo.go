package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type TransactionRepository interface {
	ListAll(ctx context.Context) ([]entities.Transaction, error)
	Save(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error)	
}