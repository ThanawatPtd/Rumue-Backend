package usecases

import (
	"context"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error)
	GetUserByID(ctx context.Context, id *pgtype.UUID) (*dbmodel.GetUserByIDRow, error)
	DeleteByID(ctx context.Context, id *pgtype.UUID) error
	GetByEmail(ctx context.Context, email *string) (*dbmodel.GetUserByEmailRow, error)
	GetUsers(ctx context.Context) (*[]dbmodel.GetAllUsersRow, error)
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

func (u *UserService) GetByEmail(ctx context.Context, email *string) (*dbmodel.GetUserByEmailRow, error) {
	response, err := u.userRepo.GetByEmail(&ctx, email)
	if err != nil{
		return nil, err
	}

	return response, nil
}

func (u *UserService) GetUserByID(ctx context.Context, id *pgtype.UUID) (*dbmodel.GetUserByIDRow, error) {
	selected, err := u.userRepo.GetByID(&ctx, id)

	if err != nil {
		return nil, err
	}

	return selected, nil
}

func (u *UserService) GetUsers(ctx context.Context) (*[]dbmodel.GetAllUsersRow, error){
	list, err:= u.userRepo.ListAll(&ctx)

	if err != nil{
		return nil, err
	}

	return list, nil
}

func (u *UserService) CreateUser(ctx context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error) {

	newUser, err := u.userRepo.Save(&ctx, user)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *UserService) DeleteByID(ctx context.Context, id *pgtype.UUID) error {
	if err := u.userRepo.Delete(&ctx, id); err != nil {
		return err
	}

	return nil
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
// 	selected.Password = user.Password

// 	u.userRepo.Update(&ctx, user)
// }

// {
// 	pass: hello
// 	addreas :; kuyy
// }
