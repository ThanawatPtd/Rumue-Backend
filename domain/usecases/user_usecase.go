package usecases

import (
	"context"
	"fmt"

	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error)
	GetUserByID(ctx context.Context, id *pgtype.UUID) (*dbmodel.GetUserByIDRow, error)
	// UpdateUser(c context.Context, id *pgtype.UUID, user *dbmodel.UpdateUserParams) (*dbmodel.UpdateUserRow, error)
}

type UserService struct {
	userRepo repositories.UserRepository
}

func ProvideUserService(userRepo repositories.UserRepository) UserUseCase {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetUserByID(ctx context.Context, id *pgtype.UUID) (*dbmodel.GetUserByIDRow, error){
	selected, err := u.userRepo.GetByID(&ctx, id)

	if err != nil{
		return nil, err
	}

	return selected, nil
}

func (u *UserService) CreateUser(ctx context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error) {

	newUser, err := u.userRepo.Save(&ctx, user)

	fmt.Println(newUser)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}


// func (u *UserService)UpdateUser(ctx context.Context, id *pgtype.UUID,user *dbmodel.UpdateUserParams) (*dbmodel.UpdateUserRow, error)  {
// 	selected, err := u.GetUserByID(ctx, id)

// 	if err != nil{
// 		return nil, err
// 	}

// 	selected.Address = user.Address
// 	selected.Email = user.Email
// 	selected.Fname = user.Fname
// 	selected.Lname = user.Lname
// 	selected.PhoneNumber = user.PhoneNumber

// }