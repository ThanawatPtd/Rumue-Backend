package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *dbmodel.CreateUserParams) (*dbmodel.User, error)
	ReadUser(ctx context.Context, email *string) (*dbmodel.User, error)
}

type UserService struct {
}
