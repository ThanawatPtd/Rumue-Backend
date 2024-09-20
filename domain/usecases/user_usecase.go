package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
)

type UserUseCase interface{
    CreateUser(user *dbmodel.User) (*dbmodel.User, error)
}

type UserService struct{
    user_repo repositories.UserRepository
	ctx *context.Context
}


func NewUserService(user_repo repositories.UserRepository, ctx *context.Context) UserUseCase{
    return &UserService{user_repo: user_repo, ctx: ctx}
}


func (us *UserService) CreateUser(user *dbmodel.User) (*dbmodel.User, error){
    return us.user_repo.Save(us.ctx, user)
}