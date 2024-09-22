package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error)
}

type UserService struct {
	userRepo repositories.UserRepository
}

func ProvideUserService(userRepo repositories.UserRepository) UserUseCase {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error) {

	newUser, err := u.userRepo.Save(ctx, user)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
